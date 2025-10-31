// Package api kuGou lite API
package api

import (
    "context"
    "io"
    "net/http"
    "strings"
    "time"
)

// 全局 HTTP 客户端，支持连接复用
var httpClient = &http.Client{
    Transport: &http.Transport{
        MaxIdleConns:           5,                // 最大空闲连接数
        MaxIdleConnsPerHost:    5,                // 每个主机的最大空闲连接数
        IdleConnTimeout:        60 * time.Second, // 空闲连接超时时间
        MaxResponseHeaderBytes: 1024 * 1024 * 10, // 10 M
    },
    Timeout: 30 * time.Second, // 默认超时时间
}

// HTTPRequest 通用 HTTP 请求函数
func HTTPRequest(method, urlStr string, data io.Reader, headers map[string]string, timeout time.Duration) (statusCode int, body []byte, cookie []*http.Cookie, err error) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, method, urlStr, data)
    if err != nil {
        return
    }

    // 设置请求头
    for key, value := range headers {
        req.Header.Set(key, value)
    }

    resp, err := httpClient.Do(req)
    if err != nil {
        return
    }
    defer func() {
        _ = resp.Body.Close()
    }()

    statusCode = resp.StatusCode
    cookie = resp.Cookies()
    body, err = io.ReadAll(resp.Body)
    return
}

// HTTPGet 发送 GET 请求
func HTTPGet(urlStr string, headers map[string]string, timeout time.Duration) (statusCode int, body []byte, cookie []*http.Cookie, err error) {
    statusCode, body, cookie, err = HTTPRequest(http.MethodGet, urlStr, nil, headers, timeout)
    return
}

// HTTPPostJSON 发送 POST JSON 请求
func HTTPPostJSON(urlStr string, jsonBody string, headers map[string]string, timeout time.Duration) (statusCode int, body []byte, cookie []*http.Cookie, err error) {
    // 设置默认 Content-Type
    if headers == nil {
        headers = make(map[string]string)
    }
    if _, exists := headers[`Content-Type`]; !exists {
        headers[`Content-Type`] = `application/json`
    }

    statusCode, body, cookie, err = HTTPRequest(`POST`, urlStr, strings.NewReader(jsonBody), headers, timeout)
    return
}

// HTTPPostForm 发送 POST 表单请求
/*func HTTPPostForm(urlStr string, postData string, headers map[string]string, timeout time.Duration) (statusCode int, body []byte, cookie []*http.Cookie, err error) {
    // 设置默认 Content-Type
    if headers == nil {
        headers = make(map[string]string)
    }
    if _, exists := headers[`Content-Type`]; !exists {
        headers[`Content-Type`] = `application/x-www-form-urlencoded`
    }

    statusCode, body, cookie, err = HTTPRequest(`POST`, urlStr, strings.NewReader(postData), headers, timeout)
    return
}*/
