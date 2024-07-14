package main

import (
	"testing"
)

func TestResolveIP(t *testing.T) {
	hostname := "example.com"
	ips, err := resolveIP(hostname)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(ips) == 0 {
		t.Fatalf("Expected at least one IP address, got none")
	}
}

func TestResolveIPInvalidHost(t *testing.T) {
	hostname := "invalid.hostname"
	_, err := resolveIP(hostname)
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
}
