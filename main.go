package main

import (
	"fmt"
	"flag"
	"strings"
)

func main() {
	fmt.Printf("goemail\n")
	serverHost := "smtp.163.com"
	serverPort := 465
	fromEmail := "anyswapinfo@163.com"
	fromPasswd := "password"

	if len(sendto) == 0 {
		fmt.Printf("Err: needed --to\n")
		return
	}
	myToers := fmt.Sprintf("%v", sendto) // 逗号隔开
	fmt.Printf("sendto: %v\n", myToers)
	myCCers := ""

	subject := subject
	body := formatInfo(info)
	// 结构体赋值
	myEmail := &EmailParam{
		ServerHost: serverHost,
		ServerPort: serverPort,
		FromEmail:  fromEmail,
		FromPasswd: fromPasswd,
		Toers:      myToers,
		CCers:      myCCers,
	}
	InitEmail(myEmail)
	SendEmail(subject, body)
}

func formatInfo(src string) string {
	if len(src) == 0 {
		return src
	}
	des := *(&src)
	des = strings.Replace(des, "#", "<br>", -1)
	return des
}

var subject string
var info string
var sendto string

func init() {
	flag.StringVar(&info, "body", "", "body")
	flag.StringVar(&subject, "subject", "", "subject")
	flag.StringVar(&sendto, "to", "", "sendto")
	flag.Parse()
}

