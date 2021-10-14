package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
)

func main() {
	router := httprouter.New()

	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getCommandOutput(command string, arguments ...string) string {
	out, _ := exec.Command(command, arguments...).Output()
	return string(out)
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	output := getCommandOutput(`C:\Go\bin\go.exe`, "version")
	_, _ = io.WriteString(w, output)
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	content, _ := ioutil.ReadFile(filepath.Join("files", name))
	_, _ = io.WriteString(w, string(content))
}
