package main

import (
  "net"
  "fmt"
  "./horse"
)

// Horse-eChat
// Did you turn devmotron into horse e-books?!
func main() {
  port := "3333"

  server, err := net.Listen("tcp", ":"+port)
  if err != nil { panic(err) }

  splitter := Splitter{}
  AddPesterBot(&splitter)

  for {
    conn, err := server.Accept()
    if err != nil { panic(err) }
    go PesterChum(&splitter).Join(conn)
  }
}

//------------------------------
// Have a robot friend
func AddPesterBot(splitter *Splitter) {
  robot                              := &Occupant{}
  robot.splitter                      = splitter
  robot.split_id, robot.from_splitter = splitter.Split()
  this_horse                         := horse.NewHorse()

  go robot.Read(func(msg string){
    neigh := this_horse.Respond(msg)
    if len(neigh) > 0 {
      horse_says := "⊚ " + neigh
      go fmt.Fprintln(robot, horse_says) //confident Go
      fmt.Println(horse_says)
    }
  })
}

// Occupant is now an io.Writter
func (occupant *Occupant) Write (message_bytes []byte) (int, error){
  for split_id, out_pipe := range occupant.splitter.splits {
    if occupant.split_id == split_id { continue }
    out_pipe <- string(message_bytes[:])
  }
  return len(message_bytes), nil
}

//------------------------------
// Bring your friends to the party
func PesterChum(splitter *Splitter) (chum *Chum) {
  chum                              = &Chum{}
  chum.MaxMessage                   = 4096
  chum.splitter                     = splitter
  chum.split_id, chum.from_splitter = splitter.Split()
  return // go knows to return `chum`
}

func (chum *Chum) Join(conn net.Conn) {
  go chum.Read(func(msg string){
    fmt.Fprint(conn, msg)
  })

  for {
    buf    := make([]byte, chum.MaxMessage)
    n, err := conn.Read(buf)
    if err != nil { return } // don't go tight loop if the client disconnects
    line   := string(buf[0:n])

    fmt.Print(line)
    chum.WriteString(line)
  }
}

func (occupant *Occupant) WriteString(message string) {
  for split_id, out_pipe := range occupant.splitter.splits {
    if occupant.split_id == split_id { continue }
    out_pipe <- message
  }
}

func (occupant *Occupant) Read(block func(msg string)) {
  for msg := range occupant.from_splitter {
    block(msg)
  }
}


type Chum struct {
  Occupant
  MaxMessage    int  // I sure would like to default this
}

type Occupant struct {
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
