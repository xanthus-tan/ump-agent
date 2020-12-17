package listener

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// ActonListener 控制监听
func ActonListener() {
	log.Println("listener start......")
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe("127.0.0.1:9999", nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s: %s %q", r.RemoteAddr, r.Method, r.URL)
	fmt.Fprintf(w, "go-daemon: %q", html.EscapeString(r.URL.Path))
}
