package main

import (
	"unicode"

	"github.com/gin-gonic/gin"
)

type Req struct {
	Text string `json:"text"`
}

type Res struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Lang    string `json:"lang"`
}

func main() {
	r := gin.Default()
	r.POST("/", detect)
	r.Run(":7891")
}

func detect(c *gin.Context) {
	var req Req
	c.BindJSON(&req)
	c.JSON(200, Res{true, 200, getLang(req.Text)})
}

func getLang(txt string) string {
	counts := make(map[string]int)

	for _, ch := range txt {
		if ch >= 0x3040 && ch <= 0x309F || ch >= 0x30A0 && ch <= 0x30FF {
			counts["ja"]++
		}
		if ch >= 0x4E00 && ch <= 0x9FFF {
			counts["zh"]++
		}
		if unicode.IsLetter(ch) && ch < 128 {
			counts["en"]++
		}
		if ch >= 0xAC00 && ch <= 0xD7AF || ch >= 0x1100 && ch <= 0x11FF {
			counts["ko"]++
		}
		if ch >= 0x0600 && ch <= 0x06FF {
			counts["ar"]++
		}
		if ch >= 0x0400 && ch <= 0x04FF {
			counts["ru"]++
		}
		if ch >= 0x0E00 && ch <= 0x0E7F {
			counts["th"]++
		}
		if ch >= 0x0900 && ch <= 0x097F {
			counts["hi"]++
		}
		if ch >= 0x0590 && ch <= 0x05FF {
			counts["he"]++
		}
		if ch >= 0x1780 && ch <= 0x17FF {
			counts["km"]++
		}
		if ch >= 0x0E80 && ch <= 0x0EFF {
			counts["lo"]++
		}
	}

	maxLang := "en"
	maxCnt := 0
	for lang, cnt := range counts {
		if cnt > maxCnt {
			maxCnt = cnt
			maxLang = lang
		}
	}

	return maxLang
}
