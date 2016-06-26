package main

/*
 * 命令行参数的读取：
 *
 * 命令行参数解析主要通过flag包实现
 *
 * // String defines a string flag with specified name, default value, and usage string.
 * // The return value is the address of a string variable that stores the value of the flag.
 * func String(name string, value string, usage string) *string {
 *	return CommandLine.String(name, value, usage)
 * }
 * 此外还有类似的Uint64、Float64等一系列方法，详见flag包结实
 */
import "flag"
import "fmt"

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
    flag.Parse()
    if infile != nil {
        fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
    }
}
