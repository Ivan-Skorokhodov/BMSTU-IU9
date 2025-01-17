package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

func main() {
    config := &ssh.ClientConfig{
		
		User: "test",
        Auth: []ssh.AuthMethod{
            ssh.Password("test"),
        },
		/*
        User: "root",
        Auth: []ssh.AuthMethod{
            ssh.Password("gOsQ5p7FUJ9w"),
        },
*/
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

	
    client, err := ssh.Dial("tcp", "localhost:2222", config)
	//client, err := ssh.Dial("tcp", "185.102.139.168:22", config)
    if err != nil {
        log.Fatalf("Failed to dial: %v", err)
    }
    defer client.Close()

    fmt.Println("Connected to SSH server!")
    fmt.Println("Type 'exit' to quit")

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Skorokhodov's custom SSH client> ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if input == "exit" {
            break
        }

        if input == "" {
            continue
        }

        session, err := client.NewSession()
        if err != nil {
            log.Printf("Failed to create session: %v", err)
            continue
        }

        session.Stdout = os.Stdout
        session.Stderr = os.Stderr

        err = session.Run(input)
        if err != nil {
            log.Printf("Failed to run command: %v", err)
        }

        session.Close()
    }
}