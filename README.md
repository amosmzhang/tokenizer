# Tokenizer

Simple go implementation of social media tokenizer.

# Example

    import "github.com/bottlenose-inc/tokenizer" 
    
    t := tokenizer.Tokenize("RT: @hello--( world)https://google.com  `_´ #test")
    // [{begin_retweet RT <nil> _rt} {user @hello <nil> _user} {part_of_speech world <nil> nn} {url https://google.com <nil> _link} {emoticon `_´ {[D:< >:( D-:< >:-( :-@[1] ;( >:O >=O D< `_´] [angry mad] -95} _emoticon} {hashtag #test <nil> _tag}]

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
