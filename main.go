package main

import (
	"log"
	"net/http"
	"upload/controllers"
	_ "upload/statik"
)

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
