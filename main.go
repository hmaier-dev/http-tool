
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/hmaier-dev/http-tool/internal/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

  const port = "8080"
  host := fmt.Sprintf("0.0.0.0:%s", port)
  srv := server.NewServer()

	log.Printf("Starting tool on %s \n", host)
	err := http.ListenAndServe(host, srv.Router)
	if err != nil {
		log.Fatal("cannot listen and server", err)
	}
}
