package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/angel-platform/angel/c2/generate"
)

func main() {
	os := flag.String("os", "linux", "target OS")
	arch := flag.String("arch", "amd64", "target arch")
	server := flag.String("server", "http://localhost:8443", "C2 server URL")
	output := flag.String("out", "lab/implants", "output directory")
	flag.Parse()

	cfg := &generate.Config{
		OS:     *os,
		Arch:   *arch,
		Format: "elf",
		Sleep:  10,
		Jitter: 0.1,
	}

	g := generate.NewGenerator(cfg, *output)
	path, err := g.Generate(*server)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generated:", path)
}
