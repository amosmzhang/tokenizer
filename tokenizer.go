package tokenizer

import (
    "regexp"
    "strings"
)

// token types
type TokenType string
const (
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
)

type Token struct {
    Type           TokenType
    Value          string
    Classification interface{}
}

var (
    rePhoto = regexp.MustCompile(`(?i:photo):`)
    reRetweet = regexp.MustCompile(`(?i:(\bRT\b)(\s*:*\s*))`)
    reMention = regexp.MustCompile(`(?i:(@[a-zA-Z0-9_]+)([^a-zA-Z0-9_\s]+))`)
    reLink = regexp.MustCompile(`(?i:(\S)(https?|s?ftp|gopher|telnet)(:\/\/))`)
    reURL = regexp.MustCompile(`^https?\:\/\/`)

    reEllipsis = regexp.MustCompile(`(\.[\. ]+)`)
    reHyphen = regexp.MustCompile(`(-[- ]+)`)
    reComma = regexp.MustCompile(`(\w)(,+)(\w)`)
    reGenitive = regexp.MustCompile(`(\w)['‘’‛]s\b`)

    reSplits = regexp.MustCompile(`\s|,\s|\.\s|\:\s|\;\s`)

    reSkipwords = regexp.MustCompile(`\b[haHA][haHA]+\b|\b[aA][hHrRaAgG]+\b|\b[aA][wW]+\b`)

    rePunctuation = regexp.MustCompile(`^[-\"\?!.,:;\(\)\{\}\[\]\\“”‘’'‛]+$`)

    KeepPunctuation = false
)

// sanitize input text using set of precompiled regex aimed at social media
func Sanitize(text string) string {
    sanitized := rePhoto.ReplaceAllString(text, "/photo")
    sanitized = reRetweet.ReplaceAllString(sanitized, "$1 ")
    sanitized = reMention.ReplaceAllString(sanitized, "$1 $2")
    sanitized = reLink.ReplaceAllString(sanitized, "$1 $2$3")

    sanitized = reEllipsis.ReplaceAllString(sanitized, "$1 ")
    sanitized = reHyphen.ReplaceAllString(sanitized, "$1 ")
    sanitized = reComma.ReplaceAllString(sanitized, "$1$2 $3")

    sanitized = strings.Replace(sanitized, "&amp;", "&", -1)
    sanitized = strings.Replace(sanitized, "&gt;", ">", -1)
    sanitized = strings.Replace(sanitized, "&lt;", "<", -1)

    sanitized = strings.Replace(sanitized, "＃", "#", -1)
    sanitized = strings.Replace(sanitized, "#", " #", -1)

    return sanitized
}

// basic splitting of text into tokens
func Tokenize(text string) []Token {
    sanitized := Sanitize(text)
    raw := reSplits.Split(sanitized, -1)

    var result []Token
    for _, t := range raw {
        if len(t) == 0 {
            continue
        }
        if reSkipwords.MatchString(t) {
            continue
        }

        var token Token
        if rePunctuation.MatchString(t) {
            if KeepPunctuation {
                token = Token{Punctuation, t, nil}
            } else {
                continue
            }
        } else if reURL.MatchString(t) {
            token = Token{URL, t, nil}
        } else if strings.HasPrefix(t, "#") {
            s := reGenitive.ReplaceAllString(t, "$1")
            // TODO more checks on hashtag
            token = Token{Hashtag, s, nil}
        } else if strings.HasPrefix(t, "$") {
            token = Token{Cashtag, t, nil}
        } else if strings.HasPrefix(t, "@") {
            // TODO more checks on username
            token = Token{User, t, nil}
        } else if strings.HasPrefix(t, "/photo") {
            token = Token{Photo, t, nil}
        } else if strings.HasPrefix(t, "RT") {
            token = Token{Retweet, t, nil}
        } else if strings.HasPrefix(t, "via") {
            token = Token{Via, t, nil}
        } else if strings.HasPrefix(t, "cc") {
            token = Token{CC, t, nil}
        } else {
            emote, err := CheckEmoticon(t)
            if err == nil {
                token = Token{Emoticon, t, emote}
            } else {
                token = Token{None, t, nil}
            }
        }

        result = append(result, token)
    }

    return result
}
