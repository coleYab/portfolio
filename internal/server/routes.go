package server

import (
	"encoding/json"
	"log"
	"net/http"

	"portfolio/cmd/web"
	"portfolio/cmd/web/components"
	"portfolio/types"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	// Apply CORS middleware
	r.Use(s.corsMiddleware)

	r.Handle("/", templ.Handler(web.Home()))

	fileServer := http.FileServer(http.FS(web.Files))
	r.PathPrefix("/assets/").Handler(fileServer)

	r.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(web.HelloForm()).ServeHTTP(w, r)
	})

	r.HandleFunc("/projects/more", func(w http.ResponseWriter, r *http.Request) {
		projects := []types.Project{
			{
				ImageURL:    "https://cdn.freelogovectors.net/wp-content/uploads/2023/02/react-logo-freelogovectors.net_.png",
				Title:       "React",
				Description: "A JavaScript library for building user interfaces.",
				TechStacks:  []string{"JavaScript", "React", "JSX", "Webpack"},
				Stars:       210000,
			},
			{
				ImageURL:    "https://cdn.freelogovectors.net/wp-content/uploads/2023/02/react-logo-freelogovectors.net_.png",
				Title:       "Vue.js",
				Description: "The Progressive JavaScript Framework.",
				TechStacks:  []string{"JavaScript", "Vue", "Vuex", "Vite"},
				Stars:       200000,
			},
		}

		projectsHtml := components.ProjectListNoParent(projects)
		projectsHtml.Render(r.Context(), w)
	})

	r.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
	

		web.Projects(projects).Render(r.Context(), w)
	})
	r.HandleFunc("/hello", web.HelloWebHandler)

	return r
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
