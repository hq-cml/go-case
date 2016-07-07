package main

import (
    "io"
    "log"
    "net/http"
    "fmt"
    "os"
    "html/template"
)

const (
    ListDir      = 0x0001
    UPLOAD_DIR   = "/tmp/uploads"
    TEMPLATE_DIR = "./views"
)

func uploadHandler(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET"{
        t, err := template.ParseFiles("upload.html")
        if err != nil{
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        t.Execute(w, nil) //Execute,根据模板语法渲染输出结果，并将结果作为返回值
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

func isExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    return os.IsExist(err)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    imageId := r.FormValue("id")
    imagePath := UPLOAD_DIR + "/" + imageId
    if ok := isExists(imagePath); !ok {       //检查文件是否存在
        http.NotFound(w, r)
        return
    }

    w.Header().Set("Content-Type", "image")
    http.ServeFile(w, r, imagePath)          //将文件读取并作为服务端的返回值
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)  //注册分发请求指针
    http.HandleFunc("/upload", uploadHandler)
    http.HandleFunc("/view", viewHandler)

    err := http.ListenAndServe(":9527", nil)
    fmt.Println("End!")
    if err != nil {
        log.Fatal("ListenAndServer:", err.Error())
    }
}
