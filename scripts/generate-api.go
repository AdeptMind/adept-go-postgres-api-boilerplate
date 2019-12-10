package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/adeptmind/adept-go-postgres-api-boilerplate/internal/config"
)

func main() {
	var (
		_, b, _, _ = runtime.Caller(0)
		root       = filepath.Join(filepath.Dir(b), "..")
	)

	config.LoadConfig()
	c := config.GetConfig()

	stdoutStderr, err := exec.Command(
		"swagger",
		"generate",
		"server",
		"-A",
		c.DbName,
		"-f",
		filepath.Join(root,
			"swagger.yml"),
		"--exclude-main",
		"-t",
		filepath.Join(root, "gen"),
	).CombinedOutput()

	if err != nil {
		fmt.Println("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf(string(stdoutStderr))
}
