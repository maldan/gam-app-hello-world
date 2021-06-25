package helloworld

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/maldan/go-restserver"
	"github.com/zserge/lorca"
)

func Start(frontFs embed.FS) {
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16000, "Server Port")
	var gui = flag.Bool("gui", false, "Use Gui")
	_ = flag.Int("width", 1100, "Window Width")
	_ = flag.Int("height", 900, "Window Height")
	_ = flag.String("data-dir", "db", "Data Directory")
	_ = flag.String("app-id", "id", "App id")
	flag.Parse()

	if *gui {
		go (func() {
			ui, _ := lorca.New("", "", 480, 320)
			defer ui.Close()
			ui.Load(fmt.Sprintf("http://%s:%d/", *host, *port))
			<-ui.Done()
			os.Exit(0)
		})()
	}

	restserver.Start(fmt.Sprintf("%s:%d", *host, *port), map[string]interface{}{
		"/": restserver.VirtualFs{Root: "frontend/build/", Fs: frontFs},
	})
}
