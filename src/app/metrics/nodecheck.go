package metrics

import (
	"io/ioutil"
    "net"
	"net/http"
	"time"
)


func nodecheck(name string, test int) {
	defer panics()

	pingtest := shell("ping -c1 " + name + ">/dev/null;echo -n $?")

	httptest := func(name string) string {
		client := http.Client{
			Transport: &http.Transport{
				Dial: func(netw, addr string) (net.Conn, error) {
					client, err := net.DialTimeout(netw, addr, time.Second*3) //设置建立连接超时
					if err != nil {
						return nil, err
					}
					client.SetDeadline(time.Now().Add(3 * time.Second)) //设置发送接收数据超时
					return client, nil
				},
				DisableKeepAlives: true,
			},
		}
		resp, err := client.Get("http://" + name + ":10255/healthz")
		if err != nil {
			return "err"
		}
		defer resp.Body.Close()
		respBytes, err := ioutil.ReadAll(resp.Body)
		return string(respBytes)
	}(name)

	connection, err := net.DialTimeout("tcp", name+":10250", time.Duration(3)*time.Second)
	if err == nil {
		connection.Close()
	}
	if pingtest != "0" && httptest != "ok" && err != nil && test == 0 {
		go label(name, test+1)
	}
	if pingtest != "0" && httptest != "ok" && err != nil && test == 1 {
		go label(name, test+1)
	}
	if pingtest != "0" && httptest != "ok" && err != nil && test == 2 {
		go label(name, test+1)
	}
	if pingtest != "0" && httptest != "ok" && err != nil && test == 3 {
		go label(name, test+1)
		go cmd(name)
	}
	if pingtest == "0" || httptest == "ok" || err == nil {
		go unlabel(name)
	}
}