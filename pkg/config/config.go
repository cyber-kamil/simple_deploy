package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

func GetSigner() ssh.Signer {
	cl := filepath.Join(os.Getenv("HOME"), "ssh", "config")
	sshConfig, err := os.ReadFile(cl)
	if err != nil {
		log.Fatalf("Failed to read SSH config file: %s", err)
	}
	sshConfigLines := strings.Split(string(sshConfig), "\n")

	var host, identityFile string

	for _, line := range sshConfigLines {
		if strings.HasPrefix(line, "Host") {
			host = strings.TrimSpace(strings.TrimPrefix(line, "Host"))
		} else if strings.HasPrefix(line, "IdentityFile") {
			identityFile = strings.TrimSpace(strings.TrimPrefix(line, "IdentityFile"))
		}

		if host == "example.com" && identityFile != "" {
			break
		}
	}
	privateKey, err := os.ReadFile(identityFile)
	if err != nil {
		log.Fatalf("Failed to read private key: %s", err)
	}

	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %s", err)
	}
	// Use signer for ssh client config

	return signer
}
