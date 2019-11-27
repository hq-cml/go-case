package main

import (
    "fmt"
    "strings"
)

func builder() {
   var builder1 strings.Builder
   builder1.WriteString("A Builder is used to efficiently build a string using Write methods.")
   fmt.Printf("The first output(%d):\n%q\n", builder1.Len(), builder1.String())
   fmt.Println()
   builder1.WriteByte(' ')
   builder1.WriteString("It minimizes memory copying. The zero value is ready to use.")
   builder1.Write([]byte{'\n', '\n'})
   builder1.WriteString("Do not copy a non-zero Builder.")
   fmt.Printf("The second output(%d):\n\"%s\"\n", builder1.Len(), builder1.String())
   fmt.Println()
}