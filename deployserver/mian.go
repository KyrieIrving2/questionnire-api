package main

import (
	"io"
	"net/http"
	"os/exec"
	"questionnaire_api/middleware/log"
)

func reLaunch() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello,this is deploy server!")
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8088", nil)
}
