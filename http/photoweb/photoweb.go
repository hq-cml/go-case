package main

import (
    "io"
    "log"
    "net/http"
    "fmt"
    "os"
    "html/template"
    "io/ioutil"
)

const (
    ListDir      = 0x0001
    UPLOAD_DIR   = "/tmp/uploads"
    TEMPLATE_DIR = "./views"
)

//渲染模板
func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
    fmt.Println(tmpl+".html")
    t, err := template.ParseFiles(tmpl+".html")
    if err != nil{
        return
    }
    err = t.Execute(w, locals)//Execute,根据模板语法渲染输出结果，并将结果作为返回值, locals是传入模板参数
    return
}

func listHandler(w http.ResponseWriter, r *http.Request) {
    fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
    if err != nil{
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    locals := make(map[string]interface{})
    images := []string{}
    for _, fileInfo := range fileInfoArr {
        images = append(images, fileInfo.Name())
    }
    fmt.Println(images)
    locals["images"] = images
    if err := renderHtml(w, "list", locals); err != nil{
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func uploadHandler(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET"{
        if err := renderHtml(w, "upload", nil); err != nil{
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
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
    http.HandleFunc("/list", listHandler)

    err := http.ListenAndServe(":9527", nil)
    fmt.Println("End!")
    if err != nil {
        log.Fatal("ListenAndServer:", err.Error())
    }
}
