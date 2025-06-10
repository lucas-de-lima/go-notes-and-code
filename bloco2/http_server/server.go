package httpserver

import (
	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func StartServer() {
	http.HandleFunc("/ping", PingHandler)
	fmt.Println("Servidor ouvindo em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}