package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	// third-party packages
	"github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	// process form submission
	if req.Method == http.MethodPost {
		mpf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer mpf.Close()
		// create sha for filename
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mpf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		// get path for new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}
		path := filepath.Join(wd, "public", "pics", fname)
		// create new file
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer nf.Close()
		// copy
		mpf.Seek(0, 0)
		io.Copy(nf, mpf)
		// add filename to this user's cookie
		c = appendValue(w, c, fname)
	}
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

// now also takes in a file name
func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
