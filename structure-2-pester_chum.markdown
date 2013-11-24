```
├── main()
├── PesterChum(*Splitter) (*Chum)
├── Chum{}
│   ├── Read( func(string) ) ()
│   ├── WriteString(string) ()
│   └── Join(net.Conn) ()
└── Splitter{}                              # A bunch of straws dipping into the chat room
    └── Split() (int, chan string)

├── NewHorse() (Horse)
├── Horse{}
│   └── Respond(string) (string)            # Say something silly
├── PrefixFromSeed(int io.Reader) (*Prefix)
├── Prefix{}
│   ├── Shift(string) ()
│   └── String() (string)
├── NewChain(int) (*Chain)
└── Chain{}
    ├── Build(io.Reader) ()                 # Build the chain from a corpus
    └── Generate(int io.Reader) (string)    # Generate text form a seed
```
