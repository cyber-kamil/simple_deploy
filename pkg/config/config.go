package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/kevinburke/ssh_config"
)

type Server struct {
	Host         string
	HostName     string
	User         string
	IdentityFile string
}

func GetServerConfig(server string) Server {
	currentServer := Server{
		server,
		ssh_config.Get(server, "HostName"),
		ssh_config.Get(server, "User"),
		ssh_config.Get(server, "IdentityFile"),
	}

	a := strings.Split(currentServer.IdentityFile, "/")
	i := 0

	// Remove the ~/ at the begining of file so we get exact identity file path
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = ""     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	simplePath := strings.Join(a[:], "/")

	currentServer.IdentityFile = filepath.Join(os.Getenv("HOME"), simplePath)

	return currentServer
}
