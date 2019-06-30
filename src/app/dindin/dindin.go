package dindin

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
)

func Cordon(nodename string, export string) {
	formt := `
	{
		"msgtype": "markdown",
		"markdown": {
			"title":"aiops",
			"text": "#### aiops \n #### Node %s  is down. \n #### I moved this node out of the cluster. \n #### Please check kubernetes cluster."

		}
	}`
	body := fmt.Sprintf(formt, nodename)
	jsonValue := []byte(body)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post("https://oapi.dingtalk.com/robot/send?access_token=b2637e4dd5eee826a7a29868cf0ff9e0818a40175dadc0840455bb0d0a03a4ab", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}

func Worring(x interface{}) {
	formt := `
	{
		"msgtype": "markdown",
		"markdown": {
			"title":"aiops",
			"text": "#### aiops \n #### 老子都挂了，赶紧看看咋回事了！！！ \n caught panic: %v" 

		}
	}`
	body := fmt.Sprintf(formt, x)
	jsonValue := []byte(body)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post("https://oapi.dingtalk.com/robot/send?access_token=b2637e4dd5eee826a7a29868cf0ff9e0818a40175dadc0840455bb0d0a03a4ab", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}