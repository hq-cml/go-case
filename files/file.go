package main

/*
 * 文件的读取与写入
 */
import (
    "strconv"
    "os"
    "fmt"
    "bufio"
    "io"
)

/*
 * 按行读取文件，将每行转成数字，存于一个slice中
 */
func read_file(infile string) (values []int, err error) {
    file, err1 := os.Open(infile)
    if err1 != nil {
        err = err1
        fmt.Println("Failed to open the input file ", infile)
        return
    }

    defer file.Close()

    br := bufio.NewReader(file)

    values = make([]int, 0)

    for {
        line, isPrefix, err1 := br.ReadLine()

        if err1 != nil {
            if err1 != io.EOF {
                err = err1
            }
            break
        }

        if isPrefix {
            fmt.Println("A too long line, seems unexpected.")
            //TODO 给err赋值
            return
        }

        str := string(line) // Convert []byte to string

        value, err1 := strconv.Atoi(str)

        if err1 != nil {
            err = err1
            return
        }

        values = append(values, value)
    }
    return
}

func write_file(values []int, outfile string) error {
    file, err := os.Create(outfile)
    if err != nil {
        fmt.Println("Failed to create the output file ", outfile)
        return err
    }

    defer file.Close()

    for _, value := range values {
        str := strconv.Itoa(value)
        file.WriteString(str + "\n")
    }
    return nil
}

func main() {
    rows, err := read_file("/tmp/test.file1")
    if err == nil{
        fmt.Println("File:", rows)
    }else{
        fmt.Println("Error:", err)
    }
}
