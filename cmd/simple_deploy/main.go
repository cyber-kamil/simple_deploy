package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cyber-kamil/simple_deploy/pkg/config"
	"golang.org/x/crypto/ssh"
)

func main() {
	servers := flag.String("servers", "", "Comma-separated list of servers")
	script := flag.String("script", "", "Location of script to run")
	flag.Parse()

	if *servers == "" || *script == "" {
		log.Fatal("Both 'servers' and 'script' flags are required")
	}

	serverList := strings.Split(*servers, ",")

	fmt.Println(config.GetSigner())
	os.Exit(1)

	for _, server := range serverList {
		client, err := ssh.Dial("tcp", server, &ssh.ClientConfig{
			User: "username",
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(config.GetSigner()),
			},
		})
		if err != nil {
			log.Fatalf("Failed to connect to %s: %s", server, err)
		}
		defer client.Close()

		session, err := client.NewSession()
		if err != nil {
			log.Fatalf("Failed to create session on %s: %s", server, err)
		}
		defer session.Close()

		output, err := session.CombinedOutput("bash " + *script)
		if err != nil {
			log.Fatalf("Failed to run script on %s: %s", server, err)
		}

		fmt.Printf("Output from %s: %s\n", server, output)
	}
}
