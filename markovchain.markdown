# Markov Generator

## Corpus
```
And, as in uffish thought he stood, The Jabberwock, with eyes of flame, came whiffling through the tulgey wood, and burbled as it came!  One, two! One, two! and through and through the vorpal blade went snicker-snack!  He left it dead, and with its head he went galumphing back.  And, has thou slain the Jabberwock? Come to my arms, my beamish boy! O frabjous day! Callooh! Callay! He chortled in his joy.  Twas brillig, and the slithy toves did gyre and gimble in the wabe all mimsy were the borogoves and the mome raths outgrabe.
```

## Compute the chain
Given the prefix length, make a hash of arrays for the next possible word choice:
```
# prefix length = 2
{
  "And as"          => [ in ]
  "as in"           => [ uffish ]
  # ...
  "the Jabberwock"  => [ with, come ]
  "Jabberwock with" => [ eyes ]
  "through the"     => [ tulgey, vorpal ]
}
```

## Walk the chain:
 1. Find our starting prefix in the hash
 2. Randomly select one of the values from the array
 3. Shift the array word on to the prefix, bumping the left-most word out
    e.g  `the Jabberwock` -> `Jabberwock come`
 4. Find the hash key matching the new prefix and repeat. (stop if there's no match)

If someone says, "the Jabberwock", the Markov generator might say:
> the Jabberwock with eyes

With a sufficiently large corpus, some interesting sentences arise.

