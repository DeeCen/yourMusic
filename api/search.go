// Package api kuGou lite API
package api

import (
    "encoding/json"
    "errors"
    "net/http"
    "strconv"
    "strings"
)

// RespAPISearchSong 歌曲搜索接口返回信息
type RespAPISearchSong struct {
    Data RespAPISearchSongData `json:"data,omitempty"`
}

// RespAPISearchSongData 歌曲搜索列表信息
type RespAPISearchSongData struct {
    Total int    `json:"total,omitempty"`
    Lists []Song `json:"lists,omitempty"`
}

// Song 歌曲信息
type Song struct {
    SingerName string   `json:"SingerName,omitempty"`
    Image      string   `json:"Image,omitempty"`
    FileHash   string   `json:"FileHash,omitempty"`
    AlbumID    string   `json:"AlbumID,omitempty"`
    FileName   string   `json:"FileName,omitempty"`
    SQ         FileHash `json:"SQ,omitempty"`
    HQ         FileHash `json:"HQ,omitempty"`
}

// FileHash 歌曲hash信息
type FileHash struct {
    Hash string `json:"Hash"`
}

// SearchSong 歌曲搜索
func SearchSong(dfid, userid, token, keyword string, page int) (ret RespAPISearchSongData, err error) {
    keyword = strings.TrimSpace(keyword)
    if keyword == `` {
        err = errors.New(`请输入搜索词`)
        return
    }

    param := apiDefaultParam(dfid, userid, token)
    param.Add(`albumhide`, `0`)
    param.Add(`iscorrection`, `1`)
    param.Add(`nocollect`, `0`)
    param.Add(`pagesize`, `10`)
    param.Add(`keyword`, keyword)
    param.Add(`page`, strconv.Itoa(page))
    param.Add(`platform`, `AndroidFilter`)

    c := new(CallAPIConfig)
    c.HTTPMethod = http.MethodGet
    c.URL = `https://complexsearch.kugou.com/v3/search/song` // gateway.kugou.com
    c.SignType = signAndroid
    c.Param = param
    c.Header = map[string]string{
        `x-router`: `complexsearch.kugou.com`,
    }

    body, err := CallKuGouAPI(c)
    if err != nil {
        return
    }

    resp := new(RespAPISearchSong)
    err = json.Unmarshal(body, resp)
    if err != nil {
        return
    }

    ret = resp.Data
    return
}

// RespAPISearch 综合搜索接口返回信息
/*type RespAPISearch struct {
    Data struct {
        Lists []struct {
            Lists []SongType `json:"lists,omitempty"`
            Type  string `json:"type,omitempty"`
        } `json:"lists,omitempty"`
    } `json:"data,omitempty"`
}

// SongType 歌曲信息
type SongType struct {
    SQFileSize int    `json:"SQFileSize,omitempty"`
    HQFileSize int    `json:"HQFileSize,omitempty"`
    SongName   string `json:"SongName,omitempty"`
    SQFileHash string `json:"SQFileHash,omitempty"`
    SingerName string `json:"SingerName,omitempty"`
    Image      string `json:"Image,omitempty"`
    FileHash   string `json:"FileHash,omitempty"`
    ID         string `json:"ID,omitempty"`
    Type       string `json:"Type,omitempty"`
    HQFileHash string `json:"HQFileHash,omitempty"`
    AlbumID    any    `json:"AlbumID,omitempty"`
}

// Search 综合搜索
func Search(dfid, userid, token, keyword string, page int) (ret []Song, err error) {
    keyword = strings.TrimSpace(keyword)
    if keyword == `` {
        err = errors.New(`请输入搜索词`)
        return
    }

    // 搜索不要求登录
    param := apiDefaultParam(dfid, userid, token)
    param.Add(`platform`, `AndroidFilter`)
    param.Add(`keyword`, keyword)
    param.Add(`page`, strconv.Itoa(page))
    param.Add(`pagesize`, `30`)
    param.Add(`cursor`, `0`)

    c := new(CallAPIConfig)
    c.HTTPMethod = http.MethodGet
    c.URL = `https://complexsearch.kugou.com/v6/search/complex`
    c.SignType = signAndroid
    c.Param = param

    body, err := CallKuGouAPI(c)
    if err != nil {
        return
    }

    resp := new(RespAPISearch)
    err = json.Unmarshal(body, resp)
    if err != nil {
        return
    }

    // 只返回 song 类型
    for _, v1 := range resp.Data.Lists {
        if v1.Type == `song` {
            for _, v2 := range v1.Lists {
                if v2.Type == `audio` {
                    ret = append(ret, v2)
                }
            }
            break
        }
    }

    return
}*/
