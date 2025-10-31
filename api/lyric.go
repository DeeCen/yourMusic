// Package api kuGou lite API
package api

import (
    "encoding/base64"
    "encoding/json"
    "errors"
    "net/http"
    "strings"
)

// GetLyric 获取歌词
func GetLyric(dfid, userid, token, hash string) (ret string, err error) {
    // 不需要登录
    hash = strings.ToLower(hash)
    if hash == `` {
        err = errors.New(`hash参数不能为空`)
        return
    }

    lyric, err := searchLyric(dfid, userid, token, hash)
    if err != nil {
        return
    }

    err = errors.New(`下载歌词失败`)
    if len(lyric.Candidates) > 0 {
        ret, err = getLyric(dfid, userid, token, lyric.Candidates[0].DownloadID, lyric.Candidates[0].AccessKey)
    }

    return
}

type RespAPISearchLyric struct {
    Status     int `json:"status,omitempty"`
    Candidates []struct {
        ID         string `json:"id,omitempty"`
        AccessKey  string `json:"accesskey,omitempty"`
        DownloadID string `json:"download_id,omitempty"`
    }   `json:"candidates,omitempty"`
}

func searchLyric(dfid, userid, token, hash string) (resp RespAPISearchLyric, err error) {
    hash = strings.ToLower(hash)
    param := apiDefaultParam(dfid, userid, token)
    param.Add(`album_audio_id`, `0`)
    param.Add(`duration`, `0`)
    param.Add(`hash`, hash)
    param.Add(`keyword`, ``)
    param.Add(`lrctxt`, `1`)
    param.Add(`man`, `no`)
    // clearDefaultParams 可能要清空默认参数

    c := new(CallAPIConfig)
    c.SignType = signAndroid
    c.HTTPMethod = http.MethodGet
    c.URL = `https://lyrics.kugou.com/v1/search`
    c.Param = param

    body, err := CallKuGouAPI(c)
    if err != nil {
        return
    }

    err = json.Unmarshal(body, &resp)
    if err != nil {
        return
    }

    if resp.Status != 200 || len(resp.Candidates) == 0 {
        err = errors.New(`暂无歌词`)
    }

    return
}

type RespAPIGetLyric struct {
    Status  int    `json:"status,omitempty"`
    Fmt     string `json:"fmt,omitempty"`
    Content string `json:"content,omitempty"`
}

func getLyric(dfid, userid, token, id, accessKey string) (ret string, err error) {
    param := apiDefaultParam(dfid, userid, token)
    param.Add(`ver`, `1`)
    param.Add(`client`, `android`)
    param.Add(`id`, id)
    param.Add(`accesskey`, accessKey)
    param.Add(`fmt`, `lrc`) // lrc普通歌词  krc逐字歌词
    param.Add(`charset`, `utf8`)

    c := new(CallAPIConfig)
    c.SignType = signAndroid
    c.HTTPMethod = http.MethodGet
    c.URL = `https://lyrics.kugou.com/download`
    c.Param = param

    body, err := CallKuGouAPI(c)
    if err != nil {
        return
    }

    resp := new(RespAPIGetLyric)
    err = json.Unmarshal(body, resp)
    if err != nil {
        return
    }

    if resp.Status != 200 || len(resp.Content) == 0 {
        err = errors.New(`下载歌词失败`)
        return
    }

    cont, err := base64.StdEncoding.DecodeString(resp.Content)
    if err != nil {
        return
    }

    if len(cont) == 0 {
        err = errors.New(`下载歌词为空`)
        return
    }

    ret = string(cont)
    return
}
