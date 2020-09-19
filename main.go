package main

import (
	"log"
	"net/http"
	"upload/controllers"
)

func main() {
	log.Println("静态服务已启动:8877")
	
	// 静态资源
	fs := http.FileServer(http.Dir("./fileinput"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", http.FileServer(http.Dir("./")))
	
	// 上传服务
	http.HandleFunc("/upload", controllers.UploadFile)
	
	http.ListenAndServe(":8877", nil)
}
