// Package main wails music player
package main

import (
    "context"
    "strconv"
    "strings"
    "yourMusic/api"
)

// App struct
type App struct {
    ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
    return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
    a.ctx = ctx
}

// RespSearch 搜索响应格式
type RespSearch struct {
    ErrMsg string                    `json:"errMsg"`
    Data   api.RespAPISearchSongData `json:"data"`
}

// SearchMusic performs a music search
func (a *App) SearchMusic(dfid, userid, token, keyword string, page int) (ret RespSearch) {
    list, err := api.SearchSong(dfid, userid, token, keyword, page)
    if err != nil {
        ret.ErrMsg = err.Error()
    }
    ret.Data = list
    return
}

// RespLogin 登录响应格式
type RespLogin struct {
    ErrMsg string `json:"errMsg"`
    Data   struct {
        Userid int    `json:"userid,omitempty"`
        Token  string `json:"token,omitempty"`
        Pic    string `json:"pic,omitempty"`
        Dfid   string `json:"dfid,omitempty"`
    }   `json:"data"`
}

func autoGetVip(dfid string, userid int, token string) {
    if userid <= 0 || token == `` {
        return
    }

    go func() {
        _ = api.FreeVipAll(dfid, strconv.Itoa(userid), token)
    }()
}

// LoginByMobile 手机号登录
func (a *App) LoginByMobile(mobile string, code string) (ret RespLogin) {
    login, err := api.LoginByVerifyCode(mobile, code)
    if err != nil {
        ret.ErrMsg = err.Error()
        return
    }

    ret.Data.Token = login.Data.Token
    ret.Data.Pic = login.Data.Pic
    ret.Data.Userid = login.Data.Userid
    ret.Data.Dfid, _ = api.GetDfid()

    // 登录自动领取vip
    autoGetVip(ret.Data.Dfid, ret.Data.Userid, ret.Data.Token)

    return
}

// LoginByToken 登录续期
func (a *App) LoginByToken(dfid, userid, token string, isAutoGetVip int) (ret RespLogin) {
    login, err := api.LoginRefresh(dfid, userid, token)
    if err != nil {
        ret.ErrMsg = err.Error()
        return
    }

    ret.Data.Token = login.Data.Token
    ret.Data.Pic = login.Data.Pic
    ret.Data.Userid = login.Data.Userid
    ret.Data.Dfid, _ = api.GetDfid()

    // 登录自动领取vip
    if isAutoGetVip > 0 {
        autoGetVip(ret.Data.Dfid, ret.Data.Userid, ret.Data.Token)
    }

    return
}

// SendMobileCode 发送验证码
func (a *App) SendMobileCode(mobile string) (ret string) {
    err := api.SendMobileCode(mobile)
    if err != nil {
        ret = err.Error()
    }
    return
}

type RespGetSongURL struct {
    Size   int32    `json:"size"`
    ErrMsg string   `json:"errMsg"`
    Data   []string `json:"data"`
    Lyric  string   `json:"lyric"`
}

// GetSongURL 获取播放地址
func (a *App) GetSongURL(dfid, userid, token string, fileHashAll string) (ret RespGetSongURL) {
    hashArr := strings.SplitN(fileHashAll, `,`, 3)
    for _, hash := range hashArr {
        list, size, err := api.GetSongURL(dfid, userid, token, `0`, `0`, hash, api.SongQualityHiRes)
        if err == nil && len(list) > 0 {
            ret.ErrMsg = ``
            ret.Size = size
            ret.Data = list
            ret.Lyric, _ = api.GetLyric(dfid, userid, token, hash) // 播放器不支持动态更新显示
            return
        }
        ret.ErrMsg = err.Error()
    }
    return

    /* 这个参数没用 ?
       arr := []api.SongQuality{
             api.SongQuality128,
             api.SongQuality320,
             api.SongQualityFlac,
             api.SongQualityHiRes,
         }
         arrKey := 0
         switch quality {
         case `320`:
             arrKey = 1
         case `flac`:
             arrKey = 2
         case `HiRes`:
             arrKey = 3
         }

         // 可以降低音质播放
         for i := arrKey; i >= 0; i-- {
             list, size, err := api.GetSongURL(dfid, userid, token, `0`, `0`, fileHash, arr[i])
             if err == nil {
                 ret.ErrMsg = ``
                 ret.Quality = string(arr[i])
                 ret.Size = size
                 ret.Data = list
                 return
             }
             ret.ErrMsg = err.Error()
         }

       return
    */
}
