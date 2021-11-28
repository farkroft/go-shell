package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.CommandContext(context.Background(), "sh", "-c", "ls")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	fmt.Println(string(output))
}
