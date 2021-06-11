package helloworld

import (
	"flag"
	"fmt"

	"github.com/maldan/go-restserver"
)

func Start() {
	var port = flag.Int("port", 16000, "port")
	// var noGui = flag.Bool("nogui", false, "gui")
	flag.Parse()

	/*if !*noGui {
		go (func() {
			debug := true
			w := webview.New(debug)
			defer w.Destroy()
			w.SetTitle("Minimal webview example")
			w.SetSize(800, 600, webview.HintNone)
			w.Navigate("https://en.m.wikipedia.org/wiki/Main_Page")
			w.Run()
			os.Exit(0)
		})()
	}*/

	restserver.Start(fmt.Sprintf("127.0.0.1:%d", *port), map[string]interface{}{
		"/": "/",
	})
}
