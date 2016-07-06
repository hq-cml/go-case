package main

import (
    "io"
    "log"
    "net/http"
    "fmt"
    "os"
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
    }else if r.Method == "POST" {
        f, h, err := r.FormFile("image") //读取表单上传的image
        if err != nil{
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        filename := h.Filename
        defer f.Close()  //注册关闭

        t, err := os.Create(UPLOAD_DIR + "/" + filename) //创建一个接受文件
        if err != nil{
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer t.Close() //注册关闭

        _, err = io.Copy(t, f) //拷贝文件到接受文件
        if err != nil{
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        //重定向到展示文件
        http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
    }else{
        io.WriteString(w,"Unknow method!")
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
