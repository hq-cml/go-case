package main

import (
    "fmt"
    "net"
    "time"
)

const (
    SERVER_PROTOCAL  = "tcp"
    SERVER_ADDRESS  = "127.0.0.1:9527"
    //SERVER_NETWORK  = "tcp"
)

var logSn = 1

func printLog(format string, args ...interface{}) {
    fmt.Printf("%d: %s", logSn, fmt.Sprintf(format, args...))
    logSn++
}

//创建TCP的server
func tcpServer() {
    var listener net.Listener
    listener, err := net.Listen(SERVER_PROTOCAL, SERVER_ADDRESS)
    if err != nil {
        printLog("Listen Error: %s\n", err)
        return
    }
    defer listener.Close()
    printLog("Got listener for the server. (local address: %s)\n", listener.Addr())

    for {
        conn, err := listener.Accept() // 阻塞直至新连接到来
        if err != nil {
            printLog("Accept Error: %s\n", err)
        }
        printLog("Established a connection with a client application. (remote address: %s)\n", conn.RemoteAddr())
        go handleConn(conn)
    }
}

func tcpClient(id int) {
    //defer wg.Done()

    //建立连接
    conn, err := net.DialTimeout(SERVER_PROTOCAL, SERVER_ADDRESS, 2*time.Second)
    if err != nil {
        printLog("Dial error: %s (client[%d])", err, id)
    }
}

func main() {

}
