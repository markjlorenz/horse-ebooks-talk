package main

import (
  "net"
  "bufio"
  "strings"
  "fmt"
)

// Yellie, the yelling-est chat server
// Warning: not for sensitive people.
func main() {
  port := "3333"

  server, err := net.Listen("tcp", ":"+port)
  if err != nil { }

  for {
    conn, err := server.Accept()
    if err != nil { panic(err) }
    go Join(conn)
  }
}

func Join(conn net.Conn) {
  for {
    reader  := bufio.NewReader(conn)
    line, _ := reader.ReadString('\n')
    fmt.Print(line)
    fmt.Fprint(conn, strings.ToUpper(line))
  }
}
