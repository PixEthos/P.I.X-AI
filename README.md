### TEST BRANTCH! ###

This branch is just the 'development branch', there's no generative for the AI as of now.
However, there's a lot I am working on, such as:
1. Markov chains, something I'm going to be working on with this, is implementing a markov chain system
2. Rewriting the token system, at the moment it doesn't really enumerate. So, I'll be working on that piece of the algorithm, basically encoding/decoding each variable. So, WIP on that piece.
3. Giving it more data to use; instead of just a basic piece, I'm going to be using the document directory for the training data. It's not doing anything, I did take that out, but it is there for the sake of later implementation.
4. More word matching, it currently understands around roughly 500 words, but it's not even 'understanding' the words; it basically can just match them. There's going to be work done there to optimize the data a bit better for matching, and improving how many words it can understand.

This is an early, early, early alpha release for those curious about what I had been working on.

The current 'release' here is something I felt I may as well create given all the crap I've been working on – this algorithm itself has been taking me a small while to build; but I have been persistent in making it.

Currently, all it does is do basic predictions, matching/filtering is a bit of a mess but that is still being worked on, and with all that; it also basically tries to figure out it's own output.

The algorithm _does_ use a GRU based system, more or less, it's crudly implemented at the moment.
