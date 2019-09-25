package main

import (
	"fmt"	
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)	

func main() {
	r := httprouter.New()
	r.POST("/upload", uploadFile)
	http.ListenAndServe("localhost:8080", r)
}

func uploadFile(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
	fmt.Println("File Upload Endpoint Hit")
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("myFile")
	if err != nil {
        fmt.Println("Error getting file")
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", header.Filename)
    fmt.Printf("File Size: %+v\n", header.Size)
    fmt.Printf("MIME Header: %+v\n", header.Header)

    tempFile, err := ioutil.TempFile("pics", header.Filename)
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
    // return that we have successfully uploaded our file!
    fmt.Fprintf(w, "Successfully Uploaded File\n")
}   