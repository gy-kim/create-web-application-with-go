package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	/* // Custom File Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("../public" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		defer f.Close()
		var contentType string
		switch {
		case strings.HasSuffix(r.URL.Path, "css"):
			contentType = "text/css"
		case strings.HasSuffix(r.URL.Path, "html"):
			contentType = "text/html"
		case strings.HasSuffix(r.URL.Path, "png"):
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)
		io.Copy(w, f)
	})
	http.ListenAndServe(":8000", nil)
	*/
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "../public"+r.URL.Path)
		})
		http.ListenAndServe(":8000", nil)
	*/
	// http.ListenAndServe(":8000", http.FileServer(http.Dir("../public")))

	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir("../public")))
	http.Handle("/css/", http.FileServer(http.Dir("../public")))
	http.ListenAndServe(":8000", nil) // http://localhost:8000/home
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "../templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
