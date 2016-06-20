package tokenizer

import (
    "encoding/json"
    "errors"

    "fmt"
)

var raw_emoticons = `
{
   "excitement": {
     "values": ["\\o/"],
     "tags": ["excitement", "praise", "jumping for joy"],
     "judgement": 95
   },
   "rock_on": {
     "values": ["\\,,/", "\\m/"],
     "tags": ["rock on", "yeah", "awesome"],
     "judgement": 85
   },
   "laughing": {
     "values": [":-D", ":D", "XD", "8D", "=D", "=3"],
     "tags": ["laughing", "lol", "big smile", "big grin"],
     "judgement": 80
   },
   "smile": {
     "values": [":-)", ":)", ":o)", ":]", ":3", ":c)", ":>", "=]", "=)", "C:"],
     "tags": ["smile", "happy", "grinning", "nice"],
     "judgement": 60
   },
   "love": {
     "values": ["<3", "<33", "<333"],
     "tags": ["love", "heart"],
     "judgement": 50
   },
   "surprise": {
     "values": [":-O", ":O", ":0", "8O"],
     "tags": ["surprise", "shock", "nondescript", "awe", "amazed"],
     "judgement": 50
   },
   "wink": {
     "values": [";-)", ";)", "*)", ";]", ";D", ";-D", ";-p"],
     "tags": ["wink", "friendly", "sarcastic"],
     "judgement": 40
   },
   "playful": {
     "values": [":-P", ":P", ":-p", ":p", "=p", "=P", ":-Þ", ":Þ", ":-b", ":b", ";p"],
     "tags": ["playful", "tongue", "sarcastic", "sticking", "just kidding", "not really", "cheeky", "joke"],
     "judgement": 20
   },
   "evil": {
     "values": [">:)", ">;)", ">:D", "}:->"],
     "tags": ["evil", "devil", "badass"],
     "judgement": 15
   },
   "shades": {
     "values": ["B)", "B-)", "8)", "8-)"],
     "tags": ["shades", "cool", "groovy"],
     "judgement": 10
   },
   "confused": {
     "values": ["@_@", "@.@"],
     "tags": ["confused", "shocked", "wtf"],
     "judgement": 0
   },
   "not_amused": {
     "values": ["~_~", "-_-"],
     "tags": ["not amused", "whatever", "boring"],
     "judgement": -10
   },
   "embarrassed": {
     "values": [":-X", ":x", ":X", "=X", "=x", ":-#", ":#"],
     "tags": ["embarrassed", "sealed lips"],
     "judgement": -15
   },
   "not_good": {
     "values": [":|", "=|"],
     "tags": ["not good", "straight face", "grim", "no expression"],
     "judgement": -20
   },
   "uneasy": {
     "values": [":-/", ":/", ":\\", "=/", "=\\", ":S"],
     "tags": ["uneasy", "skeptical", "annoyed", "undecided", "dissapointed"],
     "judgement": -40
   },
   "frown": {
     "values": [":-(", ":(", ":c", ":<", ":[", "=["],
     "tags": ["frown", "sad", "down"],
     "judgement": -60
   },
   "horror": {
     "values": ["D:", "D8", "D;", "D="],
     "tags": ["horror", "disgust", "sad"],
     "judgement": -80
   },
   "crying": {
     "values": [":'(", ";*(", "='["],
     "tags": ["crying", "tears"],
     "judgement": -85
   },
   "angry": {
     "values": ["D:<", ">:(", "D-:<", ">:-(", ":-@[1]", ";(", ">:O", ">=O", "D<", "` + "`" + `_´"],
     "tags": ["angry", "mad"],
     "judgement": -95
   }
}`

type EmoticonClass struct {
    Values     []string `json:"values"`
    Tags       []string `json:"tags"`
    Judgement  int      `json:"judgement"`
}

var EmoticonClasses = make(map[string]EmoticonClass)
var emoticonLoaded = false

func LoadEmoticonClasses() bool {
	err := json.Unmarshal([]byte(raw_emoticons), &EmoticonClasses)
	if err != nil {
        fmt.Println("unmarshal error")
		return false
	} else {
        return true
    }
}

func CheckEmoticon(text string) (EmoticonClass, error) {
    var e EmoticonClass
    if !emoticonLoaded {
        emoticonLoaded = LoadEmoticonClasses()
    }
    if !emoticonLoaded {
        return e, errors.New("Can't load emoticons")
    }

    switch text {
    case "\\o/":
        return EmoticonClasses["excitement"], nil
    case "\\,,/", "\\m/":
        return EmoticonClasses["rock_on"], nil
    case ":-D", ":D", "XD", "8D", "=D", "=3":
        return EmoticonClasses["laughing"], nil
    case ":-)", ":)", ":o)", ":]", ":3", ":c)", ":>", "=]", "=)", "C:":
        return EmoticonClasses["smile"], nil
    case "<3", "<33", "<333":
        return EmoticonClasses["love"], nil
    case ":-O", ":O", ":0", "8O":
        return EmoticonClasses["surprise"], nil
    case ";-)", ";)", "*)", ";]", ";D", ";-D", ";-p":
        return EmoticonClasses["wink"], nil
    case ":-P", ":P", ":-p", ":p", "=p", "=P", ":-Þ", ":Þ", ":-b", ":b", ";p":
        return EmoticonClasses["playful"], nil
    case ">:)", ">;)", ">:D", "}:->":
        return EmoticonClasses["evil"], nil
    case "B)", "B-)", "8)", "8-)":
        return EmoticonClasses["shades"], nil
    case "@_@", "@.@":
        return EmoticonClasses["confused"], nil
    case "~_~", "-_-":
        return EmoticonClasses["not_amused"], nil
    case ":-X", ":x", ":X", "=X", "=x", ":-#", ":#":
        return EmoticonClasses["embarrassed"], nil
    case ":|", "=|":
        return EmoticonClasses["not_good"], nil
    case ":-/", ":/", ":\\", "=/", "=\\", ":S":
        return EmoticonClasses["uneasy"], nil
    case ":-(", ":(", ":c", ":<", ":[", "=[":
        return EmoticonClasses["frown"], nil
    case "D:", "D8", "D;", "D=":
        return EmoticonClasses["horror"], nil
    case ":'(", ";*(", "='[":
        return EmoticonClasses["crying"], nil
    case "D:<", ">:(", "D-:<", ">:-(", ":-@[1]", ";(", ">:O", ">=O", "D<", "`_´":
        return EmoticonClasses["angry"], nil
    default:
        return e, errors.New("Not found")
    }
}
