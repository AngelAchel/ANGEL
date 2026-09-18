// Package doctor implements the ANGEL system compatibility checker.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "doctor" {
		runDoctor()
	} else {
		fmt.Println("Usage: angel doctor")
		os.Exit(1)
	}
}

func runDoctor() {
	fmt.Println("==========================================")
	fmt.Println("  ANGEL Doctor - System Check")
	fmt.Println("==========================================")
	fmt.Println()

	fmt.Println("--- OS ---")
	fmt.Printf("  OS: %s\n", getOS())
	fmt.Printf("  Arch: %s\n", runtime.GOARCH)
	fmt.Printf("  Kernel: %s\n", getKernel())
	fmt.Println()

	fmt.Println("--- Mandatory Dependencies ---")
	pass, fail := 0, 0
	for _, dep := range []string{"go", "curl", "git", "make"} {
		if cmdExists(dep) {
			fmt.Printf("  \033[0;32m✓\033[0m %s\n", dep)
			pass++
		} else {
			fmt.Printf("  \033[0;31m✗\033[0m %s (REQUIRED)\n", dep)
			fail++
		}
	}
	fmt.Println()

	fmt.Println("--- Optional Dependencies ---")
	for _, dep := range []string{"nmap", "netcat", "python3", "iptables"} {
		if cmdExists(dep) {
			fmt.Printf("  \033[0;32m✓\033[0m %s\n", dep)
			pass++
		} else {
			fmt.Printf("  \033[1;33m⚠\033[0m %s (optional)\n", dep)
		}
	}
	fmt.Println()

	fmt.Println("--- Docker ---")
	for _, dep := range []string{"Docker", "Docker Compose"} {
		if cmdExists(dep) {
			fmt.Printf("  \033[0;32m✓\033[0m %s\n", dep)
			pass++
		} else {
			fmt.Printf("  \033[1;33m⚠\033[0m %s (optional)\n", dep)
		}
	}
	fmt.Println()

	fmt.Println("--- Environment Variables ---")
	for _, env := range []string{"TEAMSERVER_KEY", "CRYPTO_KEY", "JWT_SECRET"} {
		if v := os.Getenv(env); v != "" {
			fmt.Printf("  \033[0;32m✓\033[0m %s set\n", env)
			pass++
		} else {
			fmt.Printf("  \033[1;33m⚠\033[0m %s not set\n", env)
		}
	}
	fmt.Println()

	fmt.Println("==========================================")
	fmt.Printf("  Results: %d PASS, %d FAIL\n", pass, fail)
	fmt.Println("==========================================")

	if fail > 0 {
		fmt.Println("  \033[0;31mSystem NOT ready for ANGEL\033[0m")
		os.Exit(1)
	} else {
		fmt.Println("  \033[0;32mSystem ready for ANGEL\033[0m")
	}
}

func getOS() string {
	out, err := exec.Command("sh", "-c", "cat /etc/os-release 2>/dev/null | grep PRETTY_NAME | cut -d'\"' -f2").Output()
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSpace(string(out))
}

func getKernel() string {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSpace(string(out))
}

func cmdExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
