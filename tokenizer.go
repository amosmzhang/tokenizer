package tokenizer

import (
    "regexp"
    "strings"
)

var (
    rePhoto = regexp.MustCompile(`(?i:photo):`)
    reRetweet = regexp.MustCompile(`(?i:(\bRT\b)(\s*:*\s*))`)
    reMention = regexp.MustCompile(`(?i:(@[a-zA-Z0-9_]+)([^a-zA-Z0-9_\s]+))`)
    reLink = regexp.MustCompile(`(?i:(\S)(https?|s?ftp|gopher|telnet)(:\/\/))`)

    reEllipsis = regexp.MustCompile(`(\.[\. ]+)`)
    reHyphen = regexp.MustCompile(`(-[- ]+)`)
    reComma = regexp.MustCompile(`(\w)(,+)(\w)`)

    reSplits = regexp.MustCompile(` |,|\.|\:|\;|\'`)
)

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

func Tokenize(text string) []string {
    sanitized := Sanitize(text)
    tokens := reSplits.Split(sanitized, -1)

    var result []string
    for _, t := range tokens {
        if len(t) > 0 {
            result = append(result, t)
        }
    }

    return result
}
