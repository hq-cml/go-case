package main

/*
 * TCP server & client例子：
 * server端计算client传来的数字，求立方根之后返回
 */
import (
    "fmt"
    "net"
    "time"
    "bytes"
    "io"
)

const (
    SERVER_PROTOCAL = "tcp"
    SERVER_ADDRESS  = "127.0.0.1:9527"
    DELIMITER       = '\t'            //定界符
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
    defer listener.Close() //注册析构listener的行为
    printLog("Got listener for the server. (local address: %s)\n", listener.Addr())

    for {
        conn, err := listener.Accept() // 阻塞直至新连接到来
        if err != nil {
            printLog("Accept Error: %s\n", err)
        }
        printLog("Established a connection with a client application. (remote address: %s)\n", conn.RemoteAddr())
        //启动处理子协程
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

//read方法
func read(conn net.Conn) (string, error) {
    readBytes := make([]byte, 1)
    var buffer bytes.Buffer  //bytes缓冲区，保存读取数据
    for {
        _, err := conn.Read(readBytes) //每次接收一个字节
        if err != nil {
            return "", err //交给上层去处理
        }
        readByte := readBytes[0]
        if readByte == DELIMITER {
            break
        }
        buffer.WriteByte(readByte)
    }
    return buffer.String(), nil //返回
}

// 法二：利用bufio来实现上述read方法
//func read(conn net.Conn) (string, error) {
//	reader := bufio.NewReader(conn)
//	readBytes, err := reader.ReadBytes(DELIMITER)
//	if err != nil {
//		return "", err
//	}
//	return string(readBytes[:len(readBytes)-1]), nil
//}

//write方法，将错误交个上层梳理
func write(conn net.Conn, content string) (int, error) {
    var buffer bytes.Buffer
    buffer.WriteString(content) //将写入内容放入缓冲区
    buffer.WriteByte(DELIMITER) //定界符
    return conn.Write(buffer.Bytes())
}

//实际处理的子协程
func handleConn(conn net.Conn) {
    defer func() {  //注册子协程的析构操作
        conn.Close()
        wg.Done()
    }()

    //无限循环，等待干活
    for {
        conn.SetReadDeadline(time.Now().Add(10 * time.Second)) 
        strReq, err := read(conn)
        if err != nil {
            if err == io.EOF {
                printLog("The connection is closed by another side. (Server)\n")
            } else {
                printLog("Read Error: %s (Server)\n", err)
            }
            break
        }
        printLog("Received request: %s (Server)\n", strReq)
        i32Req, err := convertToInt32(strReq)
        if err != nil {
            n, err := write(conn, err.Error())
            if err != nil {
                printLog("Write Error (written %d bytes): %s (Server)\n", err)
            }
            printLog("Sent response (written %d bytes): %s (Server)\n", n, err)
            continue
        }
        f64Resp := cbrt(i32Req)
        respMsg := fmt.Sprintf("The cube root of %d is %f.", i32Req, f64Resp)
        n, err := write(conn, respMsg)
        if err != nil {
            printLog("Write Error: %s (Server)\n", err)
        }
        printLog("Sent response (written %d bytes): %s (Server)\n", n, respMsg)
    }
}

func main() {

}
