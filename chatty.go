package main

import (
  "net"
  "bufio"
  "fmt"
  "./horse"
)

func main() {
  port := "3333"

  server, err := net.Listen("tcp", ":"+port)
  if err != nil { }

  splitter := Splitter{} // Bring your friends
  NewRobot(&splitter)    // Have a robot friend

  for {
    conn, err := server.Accept()
    if err != nil { panic(err) }
    go handle(conn, &splitter)
  }
}

//------------------------------
// Have a robot friend
func NewRobot(splitter *Splitter) {
  from_splitter := make(chan string)
  splitter.Split(from_splitter)

  this_horse := horse.NewHorse()

  go ReadSplitter(from_splitter, func(msg string){
    neigh := this_horse.Respond(msg)
    if len(neigh) > 0 {
      go fmt.Fprintln(splitter, neigh) //confident Go
    }
  })
}

// Splitter is now an io.Writter
func (s *Splitter) Write (message_bytes []byte) (int, error){
  for _, out_pipe := range s.splits {
    out_pipe <- string(message_bytes[:])
  }
  return len(message_bytes), nil
}

//------------------------------
// Bring your friends to the party
func handle(conn net.Conn, splitter *Splitter) {
  from_splitter := make(chan string)
  splitter.Split(from_splitter)

  go ReadSplitter(from_splitter, func(msg string){
    fmt.Fprint(conn, msg)
  })

  for {
    reader  := bufio.NewReader(conn)
    line, _ := reader.ReadString('\n')
    splitter.WriteString(line)
    fmt.Print(line)
  }
}

func ReadSplitter (from_splitter chan string, block func(msg string)) {
  for msg := range from_splitter {
    block(msg)
  }
}

type Splitter struct {
  splits []chan string
}

func (s *Splitter) Split (c chan string) {
  s.splits = append(s.splits, c)
}

func (s *Splitter) WriteString (message string) {
  for _, out_pipe := range s.splits {
    out_pipe <- message
  }
}

// ------------------------------
// BoRiNg
// func handle(conn net.Conn) {
//   for {
//     reader  := bufio.NewReader(conn)
//     line, _ := reader.ReadString('\n')
//     fmt.Print(line)
//     fmt.Fprintln(conn, "HI")
//   }
// }
