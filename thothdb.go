package main

import (
  //  "bufio"
  "fmt"
  //"io"
	"log"
  "io/ioutil"
	"net/http"
  "os"
	"strings"
)

var dbFolder string = "db/"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func fileList() []os.FileInfo{
	files, err := ioutil.ReadDir(dbFolder)
  check(err)
	return files



}

func getMarkdownFile(filename string) string {
	log.Println(filename)
  dat, err := ioutil.ReadFile("db/" + filename)

	if (err != nil) {
		var baseMessage = "An Error occurred processing this request: "
		baseMessage += err.Error()
		log.Println(baseMessage)
		return baseMessage
	} else {
		return string(dat)
	}
}


func getBodyHandler(w http.ResponseWriter, r *http.Request) {
	var returnFiles = ""
	files := fileList()

	for _, f := range files {
		returnFiles += f.Name()
	}

	fmt.Fprintf(w, returnFiles)
}

func getItemHandler(w http.ResponseWriter, r *http.Request) {
	urlPaths := strings.Split(string(r.URL.Path[1:]), "/")
	log.Println(urlPaths)
	fileList := getMarkdownFile(urlPaths[1])
	fmt.Fprintf(w, fileList)
}

func main() {
	http.HandleFunc("/", getBodyHandler)
	http.HandleFunc("/item/", getItemHandler)
	http.ListenAndServe(":8080", nil)
}
