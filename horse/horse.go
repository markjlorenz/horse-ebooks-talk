package horse

import (
  "math/rand"
  "os"
  "time"
  "log"
  "strings"
  "io"
  "bufio"
  "bytes"
)

type Horse struct {
  chain   *Chain
}

func NewHorse() Horse{
  rand.Seed(time.Now().UnixNano())

  corpus_path := "./corpus.txt"
  corpus, err := os.Open(corpus_path)
  if err != nil { log.Fatal(err) }

  prefix_len := 2
  chain      := NewChain(prefix_len)
  chain.Build(corpus)

  return Horse{chain}
}

func (h Horse) Respond (str string) string {
  if len(str) < 3 { return "" } // no shorties
  how_chatty := 7 //words
  reader     := bytes.NewBufferString(str)
  return h.chain.Generate(how_chatty, reader)
}

//--------------------
type Prefix []string

func (p Prefix) String() string {
  return strings.Join(p, "")
}

func (p Prefix) Shift(word string) {
  copy(p, p[1:])
  p[len(p)-1] = word
}

func PrefixFromSeed(prefix_len int, seed io.Reader) *Prefix {
  p    := make(Prefix, prefix_len)
  scan := bufio.NewScanner(seed)
  scan.Split(bufio.ScanWords)

  for scan.Scan() {
    copy(p[:prefix_len-1], p[1:prefix_len])
    p[prefix_len-1] = scan.Text()
  }
  return &p
}

//--------------------
type Chain struct {
  chain     map[string][]string
  prefix_len int
}

func NewChain(prefix_len int) *Chain {
  return &Chain{ make(map[string][]string), prefix_len }
}

func (c *Chain) Build(r io.Reader) {
  scan := bufio.NewScanner(r)
  p    := make(Prefix, c.prefix_len)
  scan.Split(bufio.ScanWords)
  for scan.Scan() {
    key         := p.String()
    c.chain[key] = append(c.chain[key], scan.Text())
    p.Shift(scan.Text())
  }
}

func (c *Chain) Generate(n int, seed io.Reader) string {
  p := PrefixFromSeed(c.prefix_len, seed)

  var words []string
  for i := 0; i <n; i++ {
    choices := c.chain[p.String()]
    if len(choices) == 0   { break }
    next := choices[rand.Intn(len(choices))]
    words = append(words, next)
    p.Shift(next)
  }
  return strings.Join(words, " ")
}

