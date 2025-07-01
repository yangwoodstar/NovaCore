package httpClient

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// 新增公共请求执行器
func executeRequestWithRetry(client *http.Client, req *http.Request, maxRetries int, initialDelay time.Duration) ([]byte, error) {
	retryDelay := initialDelay
	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			if attempt < maxRetries-1 {
				time.Sleep(retryDelay)
				retryDelay *= 2
				continue
			}
			break
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			lastErr = err
			if attempt < maxRetries-1 {
				time.Sleep(retryDelay)
				retryDelay *= 2
				continue
			}
			break
		}

		if resp.StatusCode == http.StatusOK {
			return body, nil
		}
		fmt.Println("HTTP request failed with status:", resp.StatusCode, "Response body:", string(body))
		lastErr = fmt.Errorf("status code: %d", resp.StatusCode)
		if attempt < maxRetries-1 {
			time.Sleep(retryDelay)
			retryDelay *= 2
		}
	}

	return nil, fmt.Errorf("after %d attempts: %v", maxRetries, lastErr)
}

// 公共头设置函数
func setCommonHeaders(req *http.Request, sn, appCode, accessKey string) {
	if sn != "" {
		req.Header.Set("source-sn", sn)
	}
	if appCode != "" {
		req.Header.Set("X-App-Code", appCode)
	}
	if accessKey != "" {
		req.Header.Set("x-access-key", accessKey)
	}
}

func ProcessPost(urlStr, data, sn, appCode, accessKey string, options ...func(*http.Request)) ([]byte, error) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	// 请求对象创建（移出循环）
	reqBody := bytes.NewBufferString(data)
	req, err := http.NewRequest("POST", urlStr, reqBody)
	if err != nil {
		return nil, err
	}

	// 请求头设置顺序调整
	for _, option := range options {
		option(req)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Charset", "utf-8")
	setCommonHeaders(req, sn, appCode, accessKey)

	return executeRequestWithRetry(client, req, 3, 100*time.Millisecond)
}

// ProcessGet 优化版
func ProcessGet(urlStr, sn, appCode, accessKey string, params map[string]string, options ...func(*http.Request)) ([]byte, error) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	// 正确解析原始URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	// 合并查询参数
	query := u.Query()
	for k, v := range params {
		query.Add(k, v)
	}
	u.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	// 移除不必要的Content-Type
	for _, option := range options {
		option(req)
	}

	req.Header.Set("Charset", "utf-8")
	setCommonHeaders(req, sn, appCode, accessKey)

	return executeRequestWithRetry(client, req, 3, 100*time.Millisecond)
}

// WithBasicAuth 设置Basic认证信息
func WithBasicAuth(username, password string) func(*http.Request) {
	return func(req *http.Request) {
		if username != "" || password != "" {
			auth := username + ":" + password
			encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
			req.Header.Set("Authorization", "Basic "+encodedAuth)
		}
	}
}
