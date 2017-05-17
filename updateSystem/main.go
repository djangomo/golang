package main

import (
	"fmt"
	"log"
	"os/exec"
)
//TODO not working
func main() {
	fmt.Println("update System")
	sudo := "sudo"
	args := []string{"-Syy", "-Su"}
	cmd := exec.Command(sudo, args...)
	fmt.Println(cmd)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
