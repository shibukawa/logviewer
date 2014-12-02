package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("query: date=%s user=%s\n", r.URL.Query()["date"][0], r.URL.Query()["user"][0])
	pattern := fmt.Sprintf("^%s.*%s", r.URL.Query()["date"][0], r.URL.Query()["user"][0])
	files, err := filepath.Glob("./logfiles/*")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "")
		return
	}
	if len(files) == 0 {
		log.Println("no file matches")
		fmt.Fprintf(w, "")
		return
	}
	args := make([]string, len(files)+1)
	args[0] = pattern
	for i := 0; i < len(files); i++ {
		args[i+1] = files[i]
	}
	out, err := exec.Command("grep", args...).CombinedOutput()
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "")
		return
	}
	fmt.Fprintf(w, string(out))
}

func main() {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/search", searchHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8888", nil)
}
