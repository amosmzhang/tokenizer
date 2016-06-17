# Tokenizer

Simple go implementation of social media tokenizer.

# Example

    import "github.com/bottlenose-inc/tokenizer" 
    
    t := tokenizer.Tokenize("RT: @hello--(world)!")
    // t is [RT @hello -- (world)!]

# TODO

Currently working on returning more informative tokens.
