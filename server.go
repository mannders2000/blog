package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /blog", getBlogHandler)
	mux.HandleFunc("GET /blog/{blogID}", blogHandler)
	mux.HandleFunc("GET /project/{projectID}", projectHandler)
	mux.HandleFunc("GET /book", bookHandler)
	mux.Handle("GET /public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	fmt.Println("starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/templates/header.tmpl", "public/templates/footer.tmpl", "public/html/index.html"))
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getBlogHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/templates/header.tmpl", "public/templates/footer.tmpl", "public/html/blog.html"))
	err := tmpl.ExecuteTemplate(w, "blog.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/templates/header.tmpl", "public/templates/footer.tmpl", "public/html/book.html"))
	err := tmpl.ExecuteTemplate(w, "book.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	blogID := r.PathValue("blogID")
	tmpl := template.Must(template.ParseFiles("public/templates/header.tmpl", "public/templates/footer.tmpl", fmt.Sprintf("public/html/blog/%s.html", blogID)))
	err := tmpl.ExecuteTemplate(w, fmt.Sprintf("%s.html", blogID), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")
	tmpl := template.Must(template.ParseFiles("public/templates/header.tmpl", "public/templates/footer.tmpl", fmt.Sprintf("public/html/projects/%s.html", projectID)))
	err := tmpl.ExecuteTemplate(w, fmt.Sprintf("%s.html", projectID), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
