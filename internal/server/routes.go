package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"portfolio/cmd/web"
	"portfolio/cmd/web/components"
	"portfolio/internal/middleware"
	"portfolio/internal/store"
	"portfolio/internal/utils"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes(st store.ProjectStore) http.Handler {
	r := mux.NewRouter()

	// Apply CORS middleware
	r.Use(s.corsMiddleware)
	r.Use(middleware.LoggerMiddleware)

	r.Handle("/", templ.Handler(web.Home()))

	fileServer := http.FileServer(http.FS(web.Files))
	r.PathPrefix("/assets/").Handler(fileServer)

	r.HandleFunc("/like", func(w http.ResponseWriter, r *http.Request) {
		toast := components.ToastSuccess("Project liked")
		toast.Render(r.Context(), w)
	})

	r.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(web.HelloForm()).ServeHTTP(w, r)
	})

	r.HandleFunc("/projects/more", func(w http.ResponseWriter, r *http.Request) {
		projectsHtml := components.ProjectListNoParent(st.GetProjects())
		projectsHtml.Render(r.Context(), w)
	})

	r.HandleFunc("/mail", handleFormSubmission)

	r.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		components.ContactUs().Render(r.Context(), w)
	})

	r.HandleFunc("/repositories", func(w http.ResponseWriter, r *http.Request) {
		components.RepoListNoParent(st.GetRepos(6)).Render(r.Context(), w)
	})

	r.HandleFunc("/repositories/more", func(w http.ResponseWriter, r *http.Request) {
		components.RepoListNoParent(st.GetRepos(6)).Render(r.Context(), w)
	})

	r.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		web.Projects(st.GetProjects()).Render(r.Context(), w)
	})

	r.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		components.HomeComponent().Render(r.Context(), w)
	})

	return r

}

func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		components.ToastDanger("failed to parse the form email").Render(r.Context(), w)
		return
	}

	// Extract form values
	formData := FormData{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	if len(strings.Split(formData.Message, " ")) > 500 {
		components.ToastDanger("failed to send email messge too long it has to be less than 500").Render(r.Context(), w)
		return
	}

	// Send the form data to FormSubmit
	err = utils.SendToTGBot(formData.FormatPaylod())
	if err != nil {
		components.ToastDanger("failed to send email").Render(r.Context(), w)
		return
	}

	// Respond to the client
	components.ToastSuccess("email sent successfully").Render(r.Context(), w)
}

// FormData represents the structure of the form data
type FormData struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (f *FormData) FormatPaylod() string {
	// Use Markdown formatting for better readability
	message := fmt.Sprintf(
		"*New Message Received* 🚀\n\n"+
			"*Name:* %s\n"+
			"*Email:* `%s`\n"+
			"*Message:*\n%s",
		f.Name, f.Email, f.Message,
	)

	return message
}

// CORS middleware
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS Headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Wildcard allows all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Credentials not allowed with wildcard origins

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
