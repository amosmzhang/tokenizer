# Tokenizer

Simple go implementation of social media tokenizer.

# Example

    import "github.com/bottlenose-inc/tokenizer" 
    
    t := tokenizer.Tokenize("RT: @hello--( world)https://google.com :D")
    // [{begin_retweet RT <nil>} {user @hello <nil>} {none world) <nil>} {url https://google.com <nil>} {emoticon :D {[:-D :D XD 8D =D =3] [laughing lol big smile big grin] 80}}]

# Token Types

    URL            TokenType = "url"
    Punctuation    TokenType = "punctuation_mark"
    Hashtag        TokenType = "hashtag"
    Cashtag        TokenType = "cashtag"
    User           TokenType = "user"
    Photo          TokenType = "begin_photo"
    Retweet        TokenType = "begin_retweet"
    Via            TokenType = "via"
    CC             TokenType = "cc"
    Emoticon       TokenType = "emoticon"
    None           TokenType = "none"
    
# Notes

Useful tool for figuring out regex in go: https://regex-golang.appspot.com/assets/html/index.html
