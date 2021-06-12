package main

import (
	"embed"

	helloworld "github.com/maldan/gam-hello-world/internal/app/hello_world"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
