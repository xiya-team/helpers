package helpers

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func GetIpInfo(ip string) string {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(`http://ip.taobao.com/service/getIpInfo.php?ip=` + ip)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}
