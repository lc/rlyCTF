package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)
func main() {

	// path to this directory
	path := "/root/rlyCTF"
	errormsg := "Something went terribly wrong - Contact @hacker_"
	fmt.Println("CTF Server Running...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		file := fmt.Sprintf("%s/files/ctf.html",path)
		index, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(w, errormsg)
		} else {
			fmt.Fprintf(w, string(index))
		}
	})
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
                file := fmt.Sprintf("%s/files/robots.txt",path)	
		robotstxt, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(w, errormsg)
		} else {
			fmt.Fprintf(w, string(robotstxt))
		}
	})
	http.HandleFunc("/solvers", func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "text/plain")
                file := fmt.Sprintf("%s/files/solvers.txt",path)
		solvers, err := ioutil.ReadFile(file)
                if err != nil {
                        fmt.Fprintf(w, errormsg)
                } else {
                        fmt.Fprintf(w, string(solvers))
                }
        })
	http.HandleFunc("/img", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		name := r.URL.Query().Get("url")
		file := strings.HasPrefix(strings.ToLower(name), "file://")
		if file != false {
			fmt.Fprintf(w, "File not found.")
		} else {
			if name != "" {
				cmd := exec.Command("curl", "-Ls", name)
				out, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Fprintf(w, "File not found.")
				}
				fmt.Fprintf(w, "%s", string(out))
			} else {
				fmt.Fprintf(w, "File not found.")

			}
		}
	})
	files := fmt.Sprintf("%s/static",path)
	fs := http.FileServer(http.Dir(files))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
