package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// resolveIP takes a hostname and returns its IP address.
func resolveIP(hostname string) ([]string, error) {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, ip := range ips {
		result = append(result, ip.String())
	}

	return result, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: dns-resolver <hostname>")
	}

	hostname := os.Args[1]
	ips, err := resolveIP(hostname)
	if err != nil {
		log.Fatalf("Could not resolve IP for hostname %s: %v", hostname, err)
	}

	fmt.Printf("IP addresses for %s:\n", hostname)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
