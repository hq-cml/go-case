package main

import (
    "bufio"
    "fmt"
    "strings"
)

func main() {
    comment := "Package bufio implements buffered I/O. " +
        "It wraps an io.Reader or io.Writer object, " +
        "creating another object (Reader or Writer) that " +
        "also implements the interface but provides buffering and " +
        "some help for textual I/O."
    basicReader := strings.NewReader(comment)
    fmt.Printf("The size of basic reader: %d\n", basicReader.Size())
    fmt.Println()

    // 示例1。
    fmt.Println("New a buffered reader ...")
    reader1 := bufio.NewReader(basicReader)
    fmt.Printf("The default size of buffered reader: %d\n", reader1.Size())
    // 此时reader1的缓冲区还没有被填充。
    fmt.Printf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
    fmt.Println()

    c, _ := reader1.ReadByte() //一旦读取，触发底层填充
    fmt.Println(string(c))
    fmt.Println(reader1.Buffered())
}

