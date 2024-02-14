package helpers

import (
	"context"
	"log"
	"net"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"errors"

	"github.com/grandcat/zeroconf"
)

func GetHDDId() (string, error) {
	// var diskId string
	out, err := exec.Command("wmic", "diskdrive", "get", "SerialNumber").Output()
	if err != nil {
		log.Println("Error occured while getting disk id ", err)
		return "", err
	}
	re := regexp.MustCompile(`(?m)^[^\r\n]+[\r\n]+\s*([^\r\n]+)`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) < 2 {
		return "", errors.New("unable to extract serial number")
	}

	diskId := strings.TrimSpace(matches[1])
	return diskId, nil
}
func GetCPUId() (string, error) {
	// var CPUId string
	out, err := exec.Command("wmic", "cpu", "get", "ProcessorId").Output()
	if err != nil {
		log.Println("Error occured while getting disk id ", err)
		return "", err
	}
	re := regexp.MustCompile(`(?m)^[^\r\n]+[\r\n]+\s*([^\r\n]+)`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) < 2 {
		return "", errors.New("unable to extract serial number")
	}

	CPUId := strings.TrimSpace(matches[1])
	return CPUId, nil
}

func GetIpAddr() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println("Error:", err)
		return "", err
	}

	// Iterate through each network interface
	for _, iface := range interfaces {
		// Get addresses for the current interface
		addrs, err := iface.Addrs()
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		// Iterate through each address for the current interface
		for _, addr := range addrs {
			// Extract the IP address from the address structure
			ip := net.ParseIP(addr.String())
			if ip == nil {
				log.Println("Error parsing IP address:", err)
				continue
			}
			log.Printf("Interface: %s, IP: %s\n", iface.Name, ip)
			return ip.String(), err
			// Print the IP address
		}
	}
	return "", err
}
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
func DiscoverService() string {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatal(err)
	}
	var ip string
	entries := make(chan *zeroconf.ServiceEntry)

	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			log.Printf("Discovered service: %v\n", entry)
			ip = entry.AddrIPv4[0].String()
			log.Println("No more entries.>>>", ip)
			if ip != "" {
				break
			}
			// Access entry.AddrIPv4 and entry.AddrIPv6 to get IP addresses
		}
		log.Println("No more entries.")
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = resolver.Browse(ctx, "_FASTQSERVER._tcp", "local.", entries)
	if err != nil {
		log.Fatal(err)
	}
	return ip
}
