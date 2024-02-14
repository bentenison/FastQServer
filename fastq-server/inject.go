package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/bentenison/fastq-server/api/handlers"
	"github.com/bentenison/fastq-server/api/reports"
	"github.com/bentenison/fastq-server/api/repository"
	"github.com/bentenison/fastq-server/api/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func inject(ds *dataSources) (*gin.Engine, error) {
	// tokenRepository := repository.NewRedisTokenRepository(ds.RedisClient)
	userRepository := repository.NewUserRepository(ds.DB)
	ticketrepo := repository.NewTicketRepo(ds.DB)
	godotenv.Load("./.env.dev")
	privateKeyFile := os.Getenv("PRIV_KEY_FILE")
	publicKeyFile := os.Getenv("PUB_KEY_FILE")
	priv, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Println("error in reading private key file:", err)
	}
	privkey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)
	if err != nil {
		log.Println("error in parsing private key :", err)
	}
	pub, err := ioutil.ReadFile(publicKeyFile)
	if err != nil {
		log.Println("error in reading public key file:", err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)
	if err != nil {
		log.Println("error in parsing public key :", err)
	}
	// load refresh token secret from env variable
	resetSecret := os.Getenv("RESET_SECRET")
	refreshSecret := os.Getenv("REFRESH_SECRET")

	resetTokenExp := os.Getenv("RESET_TOKEN_EXP")

	// load expiration lengths from env variables and parse as int
	idTokenExp := os.Getenv("ID_TOKEN_EXP")
	refreshTokenExp := os.Getenv("REFRESH_TOKEN_EXP")

	idExp, err := strconv.ParseInt(idTokenExp, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse ID_TOKEN_EXP as int: %w", err)
	}

	refreshExp, err := strconv.ParseInt(refreshTokenExp, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse REFRESH_TOKEN_EXP as int: %w", err)
	}
	resetExp, err := strconv.ParseInt(resetTokenExp, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse REFRESH_TOKEN_EXP as int: %w", err)
	}

	mailHost := os.Getenv("MAIL_HOST")
	mailPort := os.Getenv("MAIL_PORT")
	MailPort, err := strconv.ParseInt(mailPort, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse REFRESH_TOKEN_EXP as int: %w", err)
	}

	mailFrom := os.Getenv("MAIL_FROM")
	mailService := services.NewMailService(&services.MailConfig{
		Host: mailHost,
		Port: int(MailPort),
		From: mailFrom,
	})

	tokenService := services.NewTokenService(&services.TSConfig{
		// TokenRepository:       tokenRepository,
		PrivKey:               privkey,
		PubKey:                pubKey,
		RefreshSecret:         refreshSecret,
		IDExpirationSecs:      idExp,
		RefreshExpirationSecs: refreshExp,
	})
	/*
	 * service layer
	 */
	userService := services.NewUserService(&services.USConfig{
		ResetTokenExp:  resetExp,
		ResetSectret:   resetSecret,
		UserRepository: userRepository,
		MailService:    mailService,
		// ImageRepository: imageRepository,
	})
	licenseService := services.NewLicenseService(&services.LSConfig{
		DB: ds.DB,
	})
	adminRepository := repository.NewAdminRepo(ds.DB)
	adminService := services.NewAdminService(&services.ASConfig{
		AdminRepository: adminRepository,
	})
	ticketService := services.NewTicketService(&services.TConfig{
		TicketRepository: ticketrepo,
		DB:               ds.DB,
	})
	configService := services.NewConfigService(&services.CoConfig{
		DB: ds.DB,
	})

	reportrepository := reports.NewReportsDAO(ds.DB)

	// initialize gin.Engine
	router := gin.Default()

	// read in ACCOUNT_API_URL
	// baseURL := os.Getenv("ACCOUNT_API_URL")

	// read in HANDLER_TIMEOUT

	handlers.NewHandler(&handlers.Config{
		R:              router,
		UserService:    userService,
		TokenService:   tokenService,
		AdminService:   adminService,
		TicketService:  ticketService,
		LicenseService: licenseService,
		ConfigService:  configService,
		ReportService:  reportrepository,
		// BaseURL:      baseURL,
		// TimeoutDuration: time.Duration(time.Duration(ht) * time.Second),
		// MaxBodyBytes:    mbb,
	})
	return router, nil

}
