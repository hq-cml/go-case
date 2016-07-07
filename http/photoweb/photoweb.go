package main

import (
    "io"
    "log"
    "net/http"
    "fmt"
    "os"
    "html/template"
    "io/ioutil"
    "path"
)

const (
    ListDir      = 0x0001
    UPLOAD_DIR   = "/tmp/uploads"
    TEMPLATE_DIR = "/tmp/views"
)

var templates = make(map[string]*template.Template) //全局变量，预缓存模板

//init函数，在main之前执行 ：实现模板缓存的预加载等逻辑
func init() {
    fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
    if err != nil{
        //http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    var templateName, templatePath string
    for _, fileInfo := range fileInfoArr {
        templateName = fileInfo.Name()
        if ext := path.Ext(templateName); ext != ".html" { //仅加载
            continue
        }
        templatePath = TEMPLATE_DIR + "/" + templateName
        log.Println("Loading template:", templatePath)
        t := template.Must(template.ParseFiles(templatePath)) //Must表示ParseFiles必须要成功，否则直接触发错误，算是一种断言
        templates[templateName] = t
    }
}

//渲染模板
func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
    fmt.Println(tmpl+".html")
    tmpl += ".html"
    err = templates[tmpl].Execute(w, locals)//Execute,根据模板语法渲染输出结果，并将结果作为返回值, locals是传入模板参数

    if err != nil{
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
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
