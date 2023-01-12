package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cyber-kamil/simple_deploy/pkg/config"
	"github.com/cyber-kamil/simple_deploy/pkg/ssh_exec"
)

func main() {
	servers := flag.String("servers", "", "Comma-separated list of servers")
	scriptPath := flag.String("script", "", "Location of script to run")
	flag.Parse()

	if *servers == "" || *scriptPath == "" {
		log.Fatal("Both 'servers' and 'script' flags are required")
	}

	// Get the contents of the bash script to run from a file
	scriptBytes, err := os.ReadFile(*scriptPath)
	if err != nil {
		fmt.Println("Error reading bash script file:", err)
		return
	}
	script := string(scriptBytes)

	serverList := strings.Split(*servers, ",")
	for _, server := range serverList {
		serverConfig := config.GetServerConfig(server)
		if serverConfig.User == "" {
			continue
		}

		ssh_exec.ExecSSH(serverConfig, script)
	}

}
