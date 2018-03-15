package main

import (
	"flag"
	"fmt"
	"strings"
	"github.com/quexer/utee"
	"encoding/xml"
	"log"
	"bytes"
	"time"
	"io/ioutil"
	"net/http"
	"math/rand"
	"crypto/tls"
)

var (
	paySecret = "***"
	mchId = "***"
	tp int
)

func init() {
	flag.IntVar(&tp, "tp", 1, "tp")

	flag.Parse()
}

func main() {
	if tp == 1 {
		log.Println("验证方式1, mch_id", mchId)

		str, err := validWxByMode1()
		utee.Chk(err)
		log.Println(str)
		return
	}

	log.Println("验证方式2")
	validApiByMode2()
}

type Req struct {
	XMLName        xml.Name `xml:"xml"`
	MchId string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	Sign string `xml:"sign"`

}

type Resp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg string `xml:"return_msg"`
	MchId string `xml:"mch_id"`
	SandboxSignkey string `xml:"sandbox_signkey"`
}

func validWxByMode1() (string, error) {
	req := &Req{
		MchId: mchId,
		NonceStr:RandomStr(32),
	}
	str := fmt.Sprintf("mch_id=%s&nonce_str=%s&key=%s", mchId, req.NonceStr, paySecret)
	sign := strings.ToUpper(utee.PlainMd5(str))
	req.Sign = sign

	url := "https://apitest.mch.weixin.qq.com/sandboxnew/pay/getsignkey"

	b, err := xml.Marshal(req)
	utee.Chk(err)

	log.Println("req======", string(b))

	var resp *http.Response

	reqObj, err := http.NewRequest("POST", url, bytes.NewReader(b))
	utee.Chk(err)

	payCertFile := "***"
	payKeyFile := "***"

	cert, err := tls.LoadX509KeyPair(payCertFile, payKeyFile)
	utee.Chk(err)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}

	client := &http.Client{
		Transport: transport,
		Timeout: time.Minute * 2,
	}

	resp, err = client.Do(reqObj)
	utee.Chk(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	utee.Chk(err)

	log.Println("resp======", string(body))

	var respObj *Resp
	err = xml.Unmarshal(body, &respObj)
	utee.Chk(err)

	if respObj.ReturnCode == "SUCCESS" {
		return "success", nil
	}
	return "fail", nil
}

func validApiByMode2 (){
	url := "https://api.mch.weixin.qq.com"
	b, err := utee.HttpGet(url)
	utee.Chk(err)
	log.Println("return message======", string(b))
}

func RandomStr(size int) string {
	if size <= 0 {
		return ""
	}

	str := "123456789abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	for i := 0; i < size; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)

}