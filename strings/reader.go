package main

import (
    "fmt"
    "strings"
)

func reader() {
   reader1 := strings.NewReader(
       "NewReader returns a new Reader reading from s. " +
           "It is similar to bytes.NewBufferString but more efficient and read-only.")
   fmt.Printf("The size of reader: %d\n", reader1.Size())
   fmt.Printf("The reading index in reader: %d\n",
       reader1.Size()-int64(reader1.Len()))

   buf1 := make([]byte, 47)
   n, _ := reader1.Read(buf1)
   fmt.Printf("%d bytes were read. (call Read)\n", n)
   fmt.Printf("The size of reader: %d\n", reader1.Size())
   fmt.Printf("The Len of reader: %d\n", reader1.Len())
   fmt.Printf("The reading index in reader: %d\n",
       reader1.Size()-int64(reader1.Len()))
   fmt.Println()
}