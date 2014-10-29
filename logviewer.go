package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query()["user"][0])
	arg := fmt.Sprintf("^%s.*\\t%s", r.URL.Query()["date"][0], r.URL.Query()["user"][0])
	fmt.Println(arg)
	out, err := exec.Command("grep", arg, "/Users/shibukawa.yoshiki/develop/logserver/log/*").CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(out))
		fmt.Fprintf(w, "")
		return
	}
	fmt.Println(out)
	fmt.Fprintf(w, string(out))
}

func main() {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/search", searchHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8888", nil)
}
