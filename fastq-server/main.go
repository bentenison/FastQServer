package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"database/sql"
	"encoding/pem"
	"fmt"
	"io/fs"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/bentenison/fastq-server/api/helpers"
	"github.com/bentenison/fastq-server/api/services"
	"github.com/bentenison/fastq-server/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/grandcat/zeroconf"
)

func main() {
	log.Println("initializing the FASTQ server")
	log.Println("Starting server...")
	gin.SetMode(gin.ReleaseMode)
	ok, err := isMySQLInstalled()
	if err != nil {
		log.Println("error checking mysqql intallation", err)
	}
	log.Println("MySQL installed. Performing necessary actions...")
	if !ok {
		log.Fatal("error mysql not installed install it first!!!")
	}

	// initialize data sources
	ds, err := initDS()
	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	err = setupSQLAutomatic(context.Background(), ds.DB, "./Dump20240219_structures")
	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	router, err := inject(ds)

	if err != nil {
		log.Fatalf("Failure to inject data sources: %v\n", err)
		return
	}
	router.Use(cors.Default())
	// serve dist from the location
	// router.StaticFS("/", http.Dir("./client/dist"))
	router.StaticFS("/fastqclient", gin.Dir("./client/dist/", true))
	router.StaticFS("/js", http.Dir(filepath.Join(".", "client", "dist", "js")))
	router.StaticFS("/css", http.Dir(filepath.Join(".", "client", "dist", "css")))
	router.StaticFS("/img", http.Dir(filepath.Join(".", "client", "dist", "img")))
	router.StaticFS("/audio", http.Dir(filepath.Join(".", "client", "dist", "audio")))
	// router.StaticFile("/", "./client/dist/index.html")
	router.Use(gin.Recovery())
	RegisterService()

	if isprod, _ := strconv.ParseBool(os.Getenv("ISProd")); isprod {
		err = AddAudioEntries(ds.DB)
		if err != nil {
			log.Println("error adding uploaded audio", err)
		}
	}
	// ip := helpers.DiscoverService()
	// log.Println("Server IP found", ip)
	// if ip == "" {
	// }
	ip := helpers.GetOutboundIP()
	log.Println("Server IP found", ip)
	if ip == "" {
		log.Fatal("Server IP not fount")
	}
	services.UpdateServerDetailsService(ds.DB, ip)
	if err != nil {
		log.Fatalf("Failure to update server ip : %v\n", err)
		return
	}
	serverDetails, err := services.GetServerDetailsByIP(ip, ds.DB)
	if err != nil {
		log.Printf("Failure to get server details : %v\n", err)
		// return
	}
	if serverDetails.Id != "" {
		models.COMAPNY_CODE = serverDetails.Id
	}
	// services.RunCronJobs()
	// helpers.GetHDDId()
	// RegisterConsul()
	err = checkLicense()
	if err != nil {
		log.Fatalf("Failure to check license : %v\n", err)
		return
	}
	srv := &http.Server{
		Addr:    ":8090",
		Handler: router,
	}
	// if _, err := os.Stat(certFile); os.IsNotExist(err) {
	log.Println("Generating self-signed certificates...")
	if err := generateSelfSignedCert(); err != nil {
		log.Fatal("Failed to generate certificates:", err)
	}
	log.Println("Certificates generated successfully.")
	// }
	// certManager := autocert.Manager{
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("fastqsolutions.com"), // Your domain here
	// 	Cache:      autocert.DirCache("certs"),                   // Cache certificates in the "certs" directory
	// }
	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		fmt.Println("server is listening on port :8090")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
			return
		}
	}()
	go func() {
		fmt.Println("HTTPS server listening on :443")
		if err := http.ListenAndServeTLS(":443", certFile, keyFile, router); err != nil {
			log.Fatal(err)
			return
		}
	}()
	// httpsServer := &http.Server{
	// 	Addr:    ":443",
	// 	Handler: router,
	// }

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shutdown data sources
	if err := ds.close(); err != nil {
		log.Fatalf("A problem occurred gracefully shutting down data sources: %v\n", err)
	}

	// Shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
	log.Printf("Listening on port %v\n", srv.Addr)
	// Run the autocert background routine

	// select {}
}

// const MaxFileSize = 5 * 1024 * 1024 // 5 MB
func init() {
	logfile, err := os.OpenFile("app_"+time.Now().Format("2006-01-02")+".log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// mw := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	log.SetOutput(logfile)
}
func StartPreperation() {

}

func RegisterService() {
	hostname, _ := os.Hostname()
	log.Println("Hostname", hostname)
	_, err := zeroconf.Register(
		hostname,
		"_FASTQSERVER._tcp",
		"local.",
		9090, // Port on which your Golang server is running
		[]string{"txtv=0", "lo=1", "la=2"},
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	// defer server.Shutdown()

	// Your server logic goes here

	log.Println("Server is running on port 9090")

}
func checkLicense() error {
	filpathstr := "license_" + "FASTQ Solutions" + ".json"
	info, err := os.Stat(filpathstr)
	if err != nil {
		log.Println("license file not present, activate system", err)
		return err
	}
	licenseBytes, err := os.ReadFile(filpathstr)
	if err != nil {
		log.Println("unable to read activation file, activate system", err)
		return err
	}
	_, err = helpers.DecryptData(licenseBytes, info.Name())
	if err != nil {
		log.Println("unable to decrypt activation file, activate system", err)
		return err
	}
	// log.Println("decrypted Data", string(decryptedData))

	return nil
}
func AddAudioEntries(db *sql.DB) error {
	// entries, err := os.ReadDir("./uploads")
	// if err != nil {
	// 	log.Println("error reading uploads directory", err)
	// 	return err
	// }
	filepath.WalkDir("./uploads/", func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			log.Println("error in walk dir", e)
			return e
		}
		if d.IsDir() && d.Name() != "uploads" {
			return filepath.SkipDir
		}
		info, err := d.Info()
		if err != nil {
			log.Println("error occured reading file info", err)
			return err
		}
		if filepath.Ext(info.Name()) == ".wav" {
			v := models.Video{}
			v.FileName = info.Name()
			v.Size = int(info.Size())
			v.Type = "audio/wav"
			v.Id = uuid.NewString()
			t := time.Now()
			tm := t.Format("2006-01-02 15:04:05")
			v.ModifiedAt = tm
			v.CreatedAt = tm
			_, err = AddVideo(context.TODO(), db, v)
			if err != nil {
				log.Println("Error adding audio  file to selection", err)
				return err
			}
		}
		return nil
	})
	// for _, entry := range entries {

	// }
	return nil
}
func AddVideo(ctx context.Context, db *sql.DB, newVideo models.Video) (sql.Result, error) {
	query := "INSERT INTO videos (id, file_name, type, size, modified_at, created_at, " +
		"updated_by, branch_code, company_code, company_name, branch_name, isdeleted) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	// Execute the SQL query
	res, err := db.Exec(query,
		newVideo.Id,
		newVideo.FileName,
		newVideo.Type,
		newVideo.Size,
		newVideo.ModifiedAt,
		newVideo.CreatedAt,
		newVideo.UpdatedBy,
		newVideo.BranchCode,
		newVideo.CompanyCode,
		newVideo.CompanyName,
		newVideo.BranchName,
		newVideo.IsDeleted,
	)
	if err != nil {
		log.Println(err)
		return res, err
	}
	return res, nil
}

// func RegisterConsul() {
// 	config := api.DefaultConfig()
// 	client, err := api.NewClient(config)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	agent := client.Agent()

// 	serviceID := "golang-server"
// 	serviceName := "golang-service"
// 	servicePort := 9090

// 	registration := &api.AgentServiceRegistration{
// 		ID:      serviceID,
// 		Name:    serviceName,
// 		Port:    servicePort,
// 		Address: "localhost", // Use the appropriate IP or hostname here
// 	}

// 	err = agent.ServiceRegister(registration)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
const (
	certFile = "cert.pem"
	keyFile  = "key.pem"
)

func generateSelfSignedCert() error {
	priv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return err
	}

	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour) // Valid for 1 year

	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"FASTQ"},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return err
	}

	certFileData := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	privkeyBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		log.Println("error marshaling keyfile", err)
	}
	keyFileData := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privkeyBytes})

	err = os.WriteFile(certFile, certFileData, 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile(keyFile, keyFileData, 0644)
	if err != nil {
		return err
	}

	return nil
}
func isMySQLInstalled() (bool, error) {
	// Check if the MySQL executable is in the system PATH
	_, err := exec.LookPath("mysqld")
	if err != nil {
		log.Println("error in checking mysql executable", err)
		return false, nil
	}
	return err == nil, nil
}
