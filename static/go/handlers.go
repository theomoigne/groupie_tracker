package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

var (
	agent     = ""
	pIndex, _ = ioutil.ReadFile("index.html")
	pErr, _   = ioutil.ReadFile("404.html")
)

// Handler handle "/"
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		fmt.Fprintf(w, "%s", pErr)
		return
	}
	for name, headers := range r.Header {
		if name == "User-Agent" {
			agent = headers[0]
			break
		}
	}
	if strings.Contains(agent, "curl") {
		fmt.Fprintln(w, "Open Web Browser and browse: http://localhost:2030/")
	} else {
		templ, err := template.ParseFiles("index.html")
		if err != nil {
			fmt.Fprintln(w, "500")
			return
		}
		err = GetJSONArtists(w)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		PrepMainStruct()
		templ.Execute(w, Main)
	}
}

// ArtistHandle handle "/Artist/"
func ArtistHandle(w http.ResponseWriter, r *http.Request) {
	ind := 0
	path := strings.Split(r.URL.Path, "/")
	if len(path) != 4 {
		http.Redirect(w, r, "../../", 301)
		return
	}

	t, err := template.ParseFiles("static/templates/artist.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for i, v := range Artists {
		if v.Name == path[2] {
			ind = i
			break
		}
	}

	t.Execute(w, Artists[ind])
}
