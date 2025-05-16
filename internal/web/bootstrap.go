package web

import (
	"fmt"
	"net/http"
	"strconv"
)

func Bootstrap() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	serve(8080)
}

func serve(port int) {
	fmt.Printf("Listening on port %d...\n", port)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
