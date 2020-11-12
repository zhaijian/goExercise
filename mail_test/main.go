package main

import (
	"encoding/json"
	//	"fmt"
	"github.com/quexer/utee"
	"log"
	"net/url"
)

var (
	API_EMAIL_URL           = "http://sendcloud.sohu.com/webapi/mail.send_template.json"
	API_EMAIL_USER          = "*"
	API_EMAIL_KEY           = "*"
	API_EMAIL_FROM          = "*"
	API_EMAIL_FROMNAME      = "*"
	API_EMAIL_SUBJECT       = "* sdk update"
	API_EMAIL_TEMPLATE_NAME = "sdkUpdate"
)

func main() {
	vendors := []string{"***@qq.com", "***@gmail.com", "****@126.com", "***@yeah.net"}
	type S struct {
		Link     []string `json:"{{.link}}"`
		Mail     []string `json:"{{.mail}}"`
		Platform []string `json:"{{.platform}}"`
		Version  []string `json:"{{.version}}"`
		Update   []string `json:"{{.update}}"`
	}
	type M struct {
		To  []string `json:"to"`
		Sub S        `json:"sub"`
	}

	param := url.Values{
		"api_user":             {API_EMAIL_USER},
		"api_key":              {API_EMAIL_KEY},
		"from":                 {API_EMAIL_FROM},
		"fromname":             {API_EMAIL_FROMNAME},
		"subject":              {API_EMAIL_SUBJECT},
		"template_invoke_name": {API_EMAIL_TEMPLATE_NAME},
	}

	to := make([]string, 0)
	link := make([]string, 0)
	mail := make([]string, 0)
	platform := make([]string, 0)
	version := make([]string, 0)
	update := make([]string, 0)
	for i := 0; i < len(vendors); i++ {
		to = append(to, vendors[i])
		link = append(link, "http://doc.mpush.cn/sdk/ios/")
		mail = append(mail, vendors[i])
		platform = append(platform, "ios")
		version = append(version, "1.9.1")
		update = append(update, "优化SDK整体性能修复发现的问题")
		if len(to) < 2 && i <= len(vendors) {
			continue
		}
		m := &M{
			To: to,
			Sub: S{
				Link:     link,
				Mail:     mail,
				Platform: platform,
				Version:  version,
				Update:   update,
			},
		}
		msg, e := json.Marshal(m)
		utee.Chk(e)
		param.Set("substitution_vars", string(msg))
		log.Println("param", param)
		a, err := utee.HttpPost(API_EMAIL_URL, param)
		log.Print("a...", string(a))
		if err != nil {
			log.Println("send mail err", err)
		}
		to = nil
		link = nil
		mail = nil
		platform = nil
		version = nil
		update = nil
	}
}
