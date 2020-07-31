package log

import (
	"log"
	"net/http"
	"os"
	"testing"
)

// 系统日志库初步使用
func SetupLogger() {
	logFileLocation, _ := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s: %s", url, err.Error())
	} else {
		log.Printf("Status Code for url %s: %s", url, resp.Status)
		resp.Body.Close()
	}
}

func TestSimpleLog(t *testing.T) {
	SetupLogger()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}
