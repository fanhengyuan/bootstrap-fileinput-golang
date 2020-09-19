package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
	// "strconv"
)

const destLocalPath = "/media/fhy/共享盘/惠民/upload/"

type ResponseJson struct {
	Code int `json:"code"`
	Message string `json:"message"`
	ErrorMessage string `json:"error"`
}

// UploadFile uploads a file to the server
func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	
	r.ParseMultipartForm(32)
	file, handle, err := r.FormFile("uploadfile[]")
	if err != nil {
		fmt.Fprintf(w, "%v", err.Error())
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	log.Println(mimeType)
	log.Println(handle.Filename)

	saveFile(w, file, handle)
	// switch mimeType {
	// case "image/jpeg", "image/jpg", "image/png", "application/pdf":
	// 	saveFile(w, file, handle)
	// default:
	// 	jsonResponse(w, http.StatusOK, "", "请检查上传附件类型！")
	// }
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {
	t := time.Now()
	formatted := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	outfile, err := os.Create(destLocalPath + formatted + "_" + handle.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	_, err = io.Copy(outfile, file)
	if err != nil {
		log.Fatal(err)
	}

	jsonResponse(w, http.StatusCreated, "上传成功", "")
}

func jsonResponse(w http.ResponseWriter, code int, message string, errorMessage string) {
	responseStruct := ResponseJson{
		Code: code,
		Message: message,
		ErrorMessage: errorMessage,
	}
	dataRes,_ := json.Marshal(responseStruct)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, string(dataRes))
}