// Package api kuGou lite API
package api

import (
    "net/http"
    "strings"
    "testing"
    "time"
)

func TestHTTPRequest(t *testing.T) {
    _, _, _, err := HTTPGet(`https://www.baidu.com/`, nil, 0)
    if err == nil || !strings.Contains(err.Error(), `deadline exceeded`) {
        t.Fatalf(`HTTPGet NOT timeout`)
    }

    code, body, cookie, err := HTTPGet(`https://www.baidu.com/`, map[string]string{
        `user-agent`: `curl/7.83.1`,
    }, time.Second*5)
    if err != nil {
        t.Fatalf(`HTTPGet error %s`, err.Error())
    }

    if code != http.StatusOK {
        t.Fatalf(`HTTPGet statusCode error %d`, code)
    }

    if strings.Contains(string(body), `<html>`) == false {
        t.Fatalf("HTTPGet body NOT Contains <html>\n\n%s", string(body))
    }

    isFindOrz := false
    cookieVal := ``
    for _, c := range cookie {
        cookieVal += c.Name + `=` + c.Value + `, `
        if c.Name == `BDORZ` {
            isFindOrz = true
        }
    }

    if isFindOrz == false {
        t.Fatalf(`HTTPGet send curl header NOT get cookie %d : %s`, len(cookie), cookieVal)
    }
}

func BenchmarkHTTPRequest(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _, _, _, _ = HTTPGet(`https://www.baidu.com/`, map[string]string{
            `user-agent`: `curl/7.83.1`,
        }, time.Second*5)
    }
}
