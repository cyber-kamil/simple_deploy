package ssh_exec

import (
	"fmt"
	"log"
	"os"

	"github.com/cyber-kamil/simple_deploy/pkg/config"
	"golang.org/x/crypto/ssh"
)

func ExecSSH(server config.Server, script string) {
	privateKey, err := os.ReadFile(server.IdentityFile)
	if err != nil {
		log.Fatal(err)
	}

	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	config := &ssh.ClientConfig{
		User: server.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", server.HostName+":22", config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// You can now use the client to send commands to the server, or to open a new session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Run the command and print the output
	output, err := session.CombinedOutput(script)
	if err != nil {
		fmt.Println("Error running command on", server.Host+":", err)
	} else {
		fmt.Println("Output from", server.Host+":")
		fmt.Println(string(output))
	}

	fmt.Print(string(output))
}
