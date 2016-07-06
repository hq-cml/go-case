package main

import (
    "io"
    "log"
    "net/http"
    "fmt"
)

const (
    ListDir      = 0x0001
    UPLOAD_DIR   = "/tmp/uploads"
    TEMPLATE_DIR = "./views"
)

func uploadHandler(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET"{
        io.WriteString(w, `<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>Upload</title>
</head>
<body>
  <form method="POST" action="/upload" enctype="multipart/form-data">
    Choose an image to upload: <input name="image" type="file" />
    <input type="submit" value="Upload" />
  </form>
</body>
</html>`)
        return
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)  //注册分发请求指针
    http.HandleFunc("/upload", uploadHandler)

    err := http.ListenAndServe(":9527", nil)
    fmt.Println("End!")
    if err != nil {
        log.Fatal("ListenAndServer:", err.Error())
    }
}
