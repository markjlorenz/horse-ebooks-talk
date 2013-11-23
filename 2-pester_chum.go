package main

import (
  "net"
  "bufio"
  "fmt"
)

// PesterChum, chat with real people
// What pumpkin?
func main() {
  port := "3333"

  server, err := net.Listen("tcp", ":"+port)
  if err != nil { }

  splitter := Splitter{}

  for {
    conn, err := server.Accept()
    if err != nil { panic(err) }
    go PesterChum(&splitter).Join(conn)
  }
}

//------------------------------
// Bring your friends to the party
func PesterChum(splitter *Splitter) (chum *Chum) {
  chum                              = &Chum{}
  chum.splitter                     = splitter
  chum.split_id, chum.from_splitter = splitter.Split()
  return // go knows to return `chum`
}

func (chum *Chum) Join(conn net.Conn) {
  go chum.Read(func(msg string){
    fmt.Fprint(conn, msg)
  })

  for {
    reader  := bufio.NewReader(conn)
    line, _ := reader.ReadString('\n')
    fmt.Print(line)
    chum.WriteString(line)
  }
}

func (chum *Chum) WriteString(message string) {
  for split_id, out_pipe := range chum.splitter.splits {
    if chum.split_id == split_id { continue }
    out_pipe <- message
  }
}

func (chum *Chum) Read(block func(msg string)) {
  for msg := range chum.from_splitter {
    block(msg)
  }
}

type Chum struct {
  split_id      int
  splitter      *Splitter
  from_splitter chan string
}

func (s *Splitter) Split() (split_id int, from_splitter chan string) {
  from_splitter = make(chan string)

  s.splits = append(s.splits, from_splitter)
  split_id = len(s.splits) - 1
  return // go knows to return `split_id` and `from_splitter`
}

type Splitter struct {
  splits []chan string
}
