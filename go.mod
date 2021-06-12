module github.com/maldan/gam-hello-world

go 1.16

replace github.com/maldan/go-restserver => ../../../go_lib/restserver

require (
	github.com/maldan/go-restserver v1.1.1
	github.com/zserge/lorca v0.1.10
)
