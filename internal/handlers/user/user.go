package user

import (
	"bytes"
	"encoding/json"
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

type RequestLog struct {
	XForwardedFor []string `json:"x_forwarded_for,omitempty"`
	XRealIP       string   `json:"x_real_ip,omitempty"`
	RemoteAddr    string   `json:"remote_addr"`
	Body          string   `json:"body,omitempty"`
}

// prints out the complete POST-Request to /fullname
func (h *UserHandler) HandleFullname(w http.ResponseWriter, r *http.Request) {
	reqLog := RequestLog{
		RemoteAddr: r.RemoteAddr,
	}
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		reqLog.XForwardedFor = strings.Split(xff, ",")
	}
	if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
		reqLog.XRealIP = xrip
	}
	if r.Body != nil {
		body, err := io.ReadAll(r.Body)
		if err == nil {
			reqLog.Body = string(body)
			// restore body for further use
			r.Body = io.NopCloser(bytes.NewBuffer(body))
		}
	}
	if data, err := json.Marshal(reqLog); err == nil {
		log.Println(string(data))
	} else {
		log.Printf("error marshaling log: %v", err)
	}

	w.Write([]byte(`<div class='text-emerald-600'>Vielen Dank für Deine Mitarbeit!<br>Du kannst nun die IT kontaktieren, sodass sie das Upgrade anstoßen können.</div>`))
}




func init() {
	handlers.RegisterHandler(&UserHandler{})
}
