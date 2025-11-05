// Package api kuGou lite API
package api

import (
    "encoding/json"
    "errors"
    "net/http"
    "strings"
)

// SongQuality 音质参数
type SongQuality string

const (
    // SongQuality128 128K
    SongQuality128 SongQuality = `128`

    // SongQuality320 320K
    //SongQuality320 SongQuality = `320`
    // SongQualityFlac flac
    //SongQualityFlac SongQuality = `flac`

    // SongQualityHiRes high
    SongQualityHiRes SongQuality = `high`
)

// RespAPISongURL 获取歌曲url响应
type RespAPISongURL struct {
    FileSize int32    `json:"fileSize,omitempty"`
    FileName string   `json:"fileName,omitempty"`
    URL      []string `json:"url,omitempty"`
}

// GetSongURL 获取歌曲URL
func GetSongURL(dfid, userid, token, albumAudioId, albumId, hash string, quality SongQuality) (fileName string, URLList []string, size int32, err error) {
    if userid == `` || userid == `0` || token == `` {
        err = errors.New(`请先登录再获取歌曲信息`)
        return
    }

    if hash == `` {
        err = errors.New(`hash参数不能为空`)
        return
    }

    if albumAudioId == `` {
        albumAudioId = `0`
    }
    if albumId == `` {
        albumId = `0`
    }

    hash = strings.ToLower(hash)

    param := apiDefaultParam(dfid, userid, token)
    param.Add(`album_audio_id`, albumAudioId)
    param.Add(`album_id`, albumId)
    param.Add(`area_code`, `1`)
    param.Add(`hash`, hash)
    param.Add(`vipType`, `0`) // 该参数不影响url获取
    param.Add(`vipToken`, ``) // 该参数不影响url获取
    param.Add(`behavior`, `play`)
    param.Add(`pid`, `411`)
    param.Add(`cmd`, `26`)
    param.Add(`version`, `11040`)
    param.Add(`pidversion`, `3001`)
    param.Add(`IsFreePart`, `1`) //是否返回试听部分（仅部分歌曲）
    param.Add(`ssa_flag`, `is_fromtrack`)
    param.Add(`page_id`, `967177915`)
    param.Add(`quality`, string(quality))
    param.Add(`ppage_id`, `356753938,823673182,967485191`)
    param.Add(`cdnBackup`, `1`)
    param.Add(`kcard`, `0`)
    param.Add(`module`, `collection`)
    param.Add(`key`, SignKeyB(hash, param.Get(`appid`), param.Get(`mid`), userid))

    c := new(CallAPIConfig)
    c.SignType = signAndroid
    c.HTTPMethod = http.MethodGet
    c.URL = `https://gateway.kugou.com/v5/url`
    c.Param = param
    c.Header = map[string]string{
        `x-router`: `trackercdn.kugou.com`,
    }

    body, err := CallKuGouAPI(c)
    if err != nil {
        return
    }

    resp := new(RespAPISongURL)
    err = json.Unmarshal(body, resp)
    if err != nil {
        return
    }

    fileName = resp.FileName
    URLList = resp.URL // 不是vip时可能会空
    size = resp.FileSize
    if len(URLList) <= 0 {
        err = errors.New(`获取失败,可能权限不足`)
        if strings.Contains(string(body), `"priv_status":0`) {
            err = errors.New(`此歌曲暂无版权`)
        }
        return
    }

    return
}
