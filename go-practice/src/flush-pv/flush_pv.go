package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// random num from [begin, end)
func randNum(begin int, end int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(end-begin) + begin
}

func chooseUserAgent() string {
	user_agent_list := [...]string{
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko)",
		"Chrome/45.0.2454.85 Safari/537.36 115Browser/6.0.3",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SE 2.X MetaSr 1.0; SE 2.X MetaSr 1.0; .NET CLR 2.0.50727; SE 2.X MetaSr 1.0)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
	}
	return user_agent_list[randNum(0, len(user_agent_list))]
}

// Get获取响应
func doRequest(urls string, ip string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", "http://baidu.com", nil)
	req.Header.Set("User-Agent", chooseUserAgent())
	req.Header.Set("Connection", "keep-alive")

	proxy, _ := url.Parse(ip)

	client := &http.Client{}
	if ip != "local" {
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
			Timeout: 20 * time.Second,
		}
	}

	// 调用http请求
	return client.Do(req)
}

func flushPv(urls string, ip string) {
	resp, err := doRequest(urls, ip)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("req %s by ip=%s failed. err %s\n", urls, ip, err)
		return
	} else {
		fmt.Printf("req %s by ip=%s success.\n", urls, ip)
		printResp(resp)
		return
	}

}

// 打印resp 内容
func printResp(resp *http.Response) {
	if resp.StatusCode != 200 {
		fmt.Printf("resp fail statusCode=%d \n", resp.StatusCode)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
}

func main() {
	urls := "http://www.baidu.com"
	ip := "local"

	flushPv(urls, ip)
}
