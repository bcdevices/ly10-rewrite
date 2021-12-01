package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/bcdevices/ly10-rewrite/internal/config"
)

func main() {
	argLength := len(os.Args)
	if argLength < 4 {
		log.Fatalf("Usage: %s <config json> <wget args> <url>\n", os.Args[0])
	}

	configFileName := os.Args[1]
	url := os.Args[argLength-1]
	command := os.Args[2]

	config, err := config.Load(configFileName)
	if err != nil {
		panic(err)
	}

	os.Args[argLength-1] = config.AlternativeURL(url)
	cmd := exec.Command(command, os.Args[3:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
