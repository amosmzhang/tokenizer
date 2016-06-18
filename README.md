# Tokenizer

Simple go implementation of social media tokenizer.

# Example

    import "github.com/bottlenose-inc/tokenizer" 
    
    t := tokenizer.Tokenize("RT: @hello--( world)https://google.com :D")
    // [{begin_retweet RT <nil>} {user @hello <nil>} {none world) <nil>} {url https://google.com <nil>} {emoticon :D {[:-D :D XD 8D =D =3] [laughing lol big smile big grin] 80}}]
