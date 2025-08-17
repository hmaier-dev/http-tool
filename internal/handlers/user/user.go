package user

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/hmaier-dev/http-tool/internal/handlers"
)

type UserHandler struct{}

var _ handlers.DisplayHandler = (*UserHandler)(nil)

// Sets /search and all sub-routes
func (h *UserHandler) Routes(router *mux.Router) {
	router.HandleFunc("/", h.DisplayForm).Methods("GET")
	router.HandleFunc("/fullname", h.HandleFullname).Methods("POST")
}

//-----------------------------------------------------
// Displays pages
//-----------------------------------------------------
func (h *UserHandler) DisplayForm(w http.ResponseWriter, r *http.Request) {
		var templates = []string{
			"user/templates/prompt.html",
		}
		tmpl := handlers.LoadTemplates(templates)
		err := tmpl.Execute(w, map[string]any{
			"Endpoint": "/fullname",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal("Loading / wen't wrong... \n", err)
		}
}

//-----------------------------------------------------
// Handles user request
//-----------------------------------------------------

// prints out the complete POST-Request to /fullname
func (h *UserHandler) HandleFullname(w http.ResponseWriter, r *http.Request) {
	log.Println("")
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		log.Printf("X-Forwarded-For: %s \n", ips)
	}
	if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
		log.Printf("X-Real-IP: %s \n", xrip)
	}
	log.Printf("RemoteAddr: %s \n", r.RemoteAddr)
	// Optionally log body (if present)
	if r.Body != nil {
		body, err := io.ReadAll(r.Body)
		if err == nil {
			log.Printf("Body: %s", string(body))
			r.Body = io.NopCloser(bytes.NewBuffer(body))
		}
	}
	
	w.Write([]byte(""))
}




func init() {
	handlers.RegisterHandler(&UserHandler{})
}
