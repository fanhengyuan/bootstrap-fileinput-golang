package main

import (
	"log"
	"net/http"
	"upload/controllers"
	"strings"
	"encoding/base64"
	//_ "upload/statik"
)

func checkAuth(w http.ResponseWriter, r *http.Request) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 { return false }

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil { return false }

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 { return false }

	return pair[0] == "user" && pair[1] == "pass"
}


func main() {
	log.Println("静态服务已启动:8877")
	
	// 静态资源
/*	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(statikFS)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", http.FileServer(http.Dir("./")))*/

	fs := http.FileServer(http.Dir("./fileinput"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", http.FileServer(http.Dir("./")))

	// 上传服务
	http.HandleFunc("/upload", controllers.UploadFile)
	
	error := http.ListenAndServe(":8877", nil)
	if error != nil {
		log.Println(error.Error())
	}
}
