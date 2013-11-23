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
  AddPesterBot(&splitter)    // Have a robot friend

  for {
    conn, err := server.Accept()
    if err != nil { panic(err) }
    go PesterChum(&splitter).Join(conn)
  }
}

//------------------------------
// Have a robot friend
type Robot struct {
  Chum
  horse horse.Horse
}

func AddPesterBot(splitter *Splitter) {
  robot      := PesterChum(splitter)
  this_horse := horse.NewHorse()

  go robot.ReadSplitter(func(msg string){
    neigh := this_horse.Respond(msg)
    if len(neigh) > 0 {
      formatted_neigh := "⊚ " + neigh
      // go robot.WriteString(formatted_neigh)
      go fmt.Fprintln(robot, formatted_neigh) //confident Go
    }
  })
}

// Splitter is now an io.Writter
func (chum *Chum) Write (message_bytes []byte) (int, error){
  for split_id, out_pipe := range chum.splitter.splits {
    if chum.split_id == split_id { continue }
    out_pipe <- string(message_bytes[:])
  }
  return len(message_bytes), nil
}

//------------------------------
// Bring your friends to the party
type Chum struct {
  split_id      int
  splitter      *Splitter
  from_splitter chan string
}

func PesterChum(splitter *Splitter) (chum *Chum) {
  chum                              = &Chum{}
  chum.splitter                     = splitter
  chum.split_id, chum.from_splitter = splitter.Split()
  return // go knows to return `chum`
}

func (chum *Chum) Join(conn net.Conn) {
  go chum.ReadSplitter(func(msg string){
    fmt.Fprint(conn, msg)
  })

  for {
    reader  := bufio.NewReader(conn)
    line, _ := reader.ReadString('\n')
    chum.WriteString(line)
    fmt.Print(line)
  }
}

func (chum *Chum) WriteString(message string) {
  for split_id, out_pipe := range chum.splitter.splits {
    if chum.split_id == split_id { continue }
    out_pipe <- message
  }
}

func (chum *Chum) ReadSplitter(block func(msg string)) {
  for msg := range chum.from_splitter {
    block(msg)
  }
}

type Splitter struct {
  splits []chan string
}

func (s *Splitter) Split() (split_id int, from_splitter chan string) {
  from_splitter = make(chan string)

  s.splits = append(s.splits, from_splitter)
  split_id = len(s.splits) - 1
  return
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
