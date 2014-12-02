package main

import (
	"fmt"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("query: date=%s user=%s\n", r.URL.Query()["date"][0], r.URL.Query()["user"][0])
	pattern := fmt.Sprintf("^%s.*%s", r.URL.Query()["date"][0], r.URL.Query()["user"][0])
	files, err := filepath.Glob("/var/log/logviewer/*")
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
	logger, err := syslog.New(syslog.LOG_NOTICE, "logviewer")
	if err != nil {
		panic(err)
	}
	log.SetOutput(logger)
	log.SetPrefix("logviewer:")
	log.Println("starting logviewer")
	portStr := os.Getenv("PORT_NUMBER")
	url := ":8888"
	if len(portStr) != 0 {
		url = fmt.Sprintf(":%s", portStr)
	}
	log.Printf("Url is %s\n", url)
	workDirStr := os.Getenv("WORK_DIR")
	if len(workDirStr) == 0 {
		os.Chdir("/usr/local/logviewer")
		log.Println("Work is /usr/local/logviewer")
	} else {
		os.Chdir(workDirStr)
		log.Printf("Work is %s\n", workDirStr)
	}
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/search", searchHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))
	err = http.ListenAndServe(url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(2)
	}
	log.Println("closing logviewer")
}
