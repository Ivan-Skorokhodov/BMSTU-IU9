package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gliderlabs/ssh"
)

func main() {
    ssh.Handle(func(s ssh.Session) {
        cmd := s.Command()
        if len(cmd) == 0 {
            fmt.Fprintln(s, "Welcome to GO SSH Server!")
            return
        }

		command := ""
		for i := 0; i < len(cmd); i++ {
			command += cmd[i] + " "
		}
		termCmd := exec.Command("bash", "-c", command)
		termCmd.Stdout = s
		termCmd.Stderr = s
		err := termCmd.Run()
		if err != nil {
			fmt.Fprintf(s, "Error in runtime: %v\n", err)
			return
		}
    })

    server := &ssh.Server{
        Addr: ":2222",
        PasswordHandler: func(ctx ssh.Context, pass string) bool {
            if ctx.User() == "test" && pass == "test" {
				return true
			}
			return false
        },
    }

    log.Printf("Starting SSH server on port 2222...")
    log.Fatal(server.ListenAndServe())
}