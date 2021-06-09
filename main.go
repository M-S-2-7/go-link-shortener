package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var links = make(map[string]string)
type Link struct {
	URL string
}


func redirecter(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/r/")
	link, ok := links[id]

	if r.URL.Path[1:] == "" {
		http.ServeFile(w, r, "static/index.html")
		return
	}

	if !ok {
		fmt.Fprintln(w, "Incorrect Link")
		return
	}

	http.Redirect(w, r, link, http.StatusSeeOther)
}

func shortener(w http.ResponseWriter, r *http.Request) {
	link, ok := r.URL.Query()["link"]

	if !ok || len(link) > 1 || len(link[0]) < 1{
		fmt.Fprintln(w, "give your link to 'link' parameter.")
		return
	}

	var alphabet []string
	for l := 'a'; l < 'z'; l++ {
		alphabet = append(alphabet, string(l))
	}

	randomText := ""
	for i := 0; i < 5; i++ {
		randomText += alphabet[rand.Intn(len(alphabet))]
	}

	links[randomText] = link[0]
	t, _ := template.ParseFiles("static/result.html")
	t.Execute(w, &Link{URL: "http://localhost:8080/r/" + randomText})

}

func getLinks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, links)
}

func main() {
	rand.Seed(time.Now().Unix())

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/shortener/", shortener)
	http.HandleFunc("/getlinks/", getLinks)
	http.HandleFunc("/r/", redirecter)


	log.Fatal(http.ListenAndServe(":8080", nil))
}
