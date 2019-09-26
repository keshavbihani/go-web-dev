package main

import (
	"fmt"	
	"github.com/julienschmidt/httprouter"
	"net/http"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "mime/multipart"
    "bytes"
)	

func main() {
	r := httprouter.New()
	r.POST("/upload", uploadFile)
	http.ListenAndServe("localhost:8080", r)
}

func uploadFile(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
    s, err := session.NewSession(&aws.Config{
      Region: aws.String("us-east-1"), 
      })

    if err != nil {
    fmt.Println("Error creating session",err)
    }  

	fmt.Println("File Upload Endpoint Hit")
	r.ParseMultipartForm(20000)
    if err != nil {
        fmt.Fprintln(w, err)
        return
    }
	formdata := r.MultipartForm 
    files := formdata.File["multiplefiles"] 


    for i, _ := range files { // loop through the files one by one
        file, err := files[i].Open()
        defer file.Close()
        if err != nil {
            fmt.Fprintln(w, err)
            return
        }
        fileName, err := fileUpload(s, file, files[i])
        if err != nil {
            fmt.Println("Error uploading file",err)
        } 
        fmt.Println(fileName)
    } 
}   

func fileUpload(s *session.Session, file multipart.File, 
    fileHeader *multipart.FileHeader) (string,error){
    size := fileHeader.Size
    buffer := make([]byte, size)
    file.Read(buffer)
    fmt.Println(fileHeader.Filename)
    tempFileName := "pics/" + fileHeader.Filename
    _, err := s3.New(s).PutObject(&s3.PutObjectInput{
        Bucket:               aws.String("golang-test-1"),
        Key:                  aws.String(tempFileName),
        ACL:                  aws.String("public-read"),// could be private if you want it to be access by only authorized users
        Body:                 bytes.NewReader(buffer),
        ContentLength:        aws.Int64(int64(size)),
        ContentType:          aws.String(http.DetectContentType(buffer)),
        ContentDisposition:   aws.String("attachment"),
        ServerSideEncryption: aws.String("AES256"),
        StorageClass:         aws.String("INTELLIGENT_TIERING"),
    })
    if err != nil {
        return "", err
    }

    return tempFileName, err    
}