```
├── main()
├── AddPesterBot(*Splitter)                 # Put a bot in the chat room
├── PesterChum(*Splitter) (*Chum)
├── Occupant{}                              # A human or robot in the chat room
│   ├── Write([]byte) (int, error)
│   ├── WriteString(string) ()
│   └── Read( func(string) ) ()
├── Chum{}                                  # A human in the chat room
│   └── Join(net.Conn) ()
└── Splitter{}                              # A bunch of straws dipping into the chat room
    └── Split() (int, chan string)

├── NewHorse() (Horse)
├── Horse{}
│   └── Respond(string) (string)            # Say something silly
├── PrefixFromSeed(int io.Reader) (*Prefix)
├── Prefix{}                                # A key in the Markov chain hash
│   ├── Shift(string) ()
│   └── String() (string)
├── NewChain(int) (*Chain)
└── Chain{}
    ├── Build(io.Reader) ()                 # Build the chain from a corpus
    └── Generate(int io.Reader) (string)    # Generate text form a seed
```
