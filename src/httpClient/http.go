package httpClient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// ProcessPost 处理 POST 请求
func ProcessPost(url, data, sn, appCode, accessKey string) ([]byte, error) {

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return []byte(""), err
	}

	// 设置请求头
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("charset", "utf-8")

	if sn != "" {
		req.Header.Set("source-sn", sn)
	}
	if appCode != "" {
		req.Header.Set("X-App-Code", appCode)
	}
	if accessKey != "" {
		req.Header.Set("x-access-key", accessKey)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
		}
	}(resp.Body)

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	if resp.StatusCode == http.StatusOK {
		return body, nil
	} else {
		return []byte(""), fmt.Errorf("response status code: %d", resp.StatusCode)
	}
}

// ProcessGet 处理 GET 请求
func ProcessGet(urlStr, sn, appCode, accessKey string, params map[string]string) ([]byte, error) {
	// 构建查询字符串
	queryString := ""
	if params != nil {
		queryValues := url.Values{}
		for key, value := range params {
			queryValues.Add(key, value)
		}
		queryString = "?" + queryValues.Encode()
	}

	// 创建 HTTP 客户端
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	// 创建请求
	requestUrl := urlStr + queryString
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return []byte(""), err
	}

	// 设置请求头
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("charset", "utf-8")
	if sn != "" {
		req.Header.Set("source-sn", sn)
	}
	if appCode != "" {
		req.Header.Set("X-App-Code", appCode)
	}
	if accessKey != "" {
		req.Header.Set("x-access-key", accessKey)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
		}
	}(resp.Body) // 确保在函数退出时关闭响应体

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	// 检查状态码
	if resp.StatusCode == http.StatusOK {
		return body, nil
	} else {
		return []byte(""), fmt.Errorf("response status code: %d", resp.StatusCode)
	}
}
