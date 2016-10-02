package main

import (
    "net"
    "fmt"
)


func serve(a *args) {
    l, _ := net.Listen("tcp", a.host+":"+a.port)
    defer l.Close()
    for {
        conn, _ := l.Accept()
        go handleRequest(conn)
    }
}


func handleRequest(conn net.Conn) {
    buf := make([]byte, 1024)
    conn.Read(buf)
    data := string(buf)
    if len(data)>0 {
        fmt.Println(data)
    }
}
