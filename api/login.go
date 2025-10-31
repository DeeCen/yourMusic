// Package api kuGou lite API
package api

import (
    "encoding/json"
    "errors"
    "net/http"
    "strings"
)

// SendMobileCode 发送验证码到手机号
func SendMobileCode(mobile string) (err error) {
    if len(mobile) != 11 {
        err = errors.New(`mobile length error`)
        return
    }

    c := new(CallAPIConfig)
    c.SignType = signAndroid
    c.HTTPMethod = http.MethodPost
    c.URL = `http://login.user.kugou.com/v7/send_mobile_code`
    c.Param = apiDefaultParam(``, ``, ``)
    c.Body = map[string]string{
        `businessid`: `5`,
        `mobile`:     mobile,
        `plat`:       `3`,
    }

    _, err = CallKuGouAPI(c)
    return
}

// LoginAPIResp 登录响应信息
type LoginAPIResp struct {
    Data struct {
        //IsVip    int    `json:"is_vip,omitempty"`
        //VipType  int    `json:"vip_type,omitempty"`
        //Mobile   int    `json:"mobile,omitempty"`
        Userid int    `json:"userid,omitempty"`
        Token  string `json:"token,omitempty"`
        Pic    string `json:"pic,omitempty"`
        //Username string `json:"username,omitempty"`
        //VipToken string `json:"vip_token,omitempty"`
        //Nickname string `json:"nickname,omitempty"`
    } `json:"data,omitempty"`
}

// LoginByVerifyCode 手机号登录
func LoginByVerifyCode(mobile, smsCode string) (resp LoginAPIResp, err error) {
    if len(mobile) != 11 {
        err = errors.New(`请输入11位手机号`)
        return
    }

    if len(smsCode) != 6 {
        err = errors.New(`请输入6位短信验证码`)
        return
    }

    ms := msStr()
    p2Data := map[string]string{
        `clienttime_ms`: ms,
        `code`:          smsCode,
        `mobile`:        mobile,
    }
    p2Str, err := json.Marshal(p2Data)
    if err != nil {
        return
    }

    p2, err := RSAPublicEncryptNoPadding(p2Str, publicLiteRasKey)
    if err != nil {
        return
    }

    postData := map[string]string{
        `plat`:          `1`,
        `support_multi`: `1`,
        `t1`:            `0`,
        `t2`:            `0`,
        `userid`:        `0`,
        `clienttime_ms`: ms,
        `mobile`:        mobile,
        `key`:           SignKeyA(ms),
        `p2`:            strings.ToUpper(p2),
    }

    c := new(CallAPIConfig)
    c.SignType = signAndroid
    c.HTTPMethod = http.MethodPost
    c.URL = `https://gateway.kugou.com/v6/login_by_verifycode`
    c.Param = apiDefaultParam(``, ``, ``)
    c.Body = postData
    c.Header = map[string]string{
        `x-router`: `login.user.kugou.com`,
    }

    body, err := CallKuGouAPI(c)
    if err != nil {
        return
    }

    err = json.Unmarshal(body, &resp)
    return
}

const liteKey = `c24f74ca2820225badc01946dba4fdf7`
const liteIv = `adc01946dba4fdf7`

// LoginRefresh 刷新登录状态(登录续期)
func LoginRefresh(dfid, userid, token string) (resp LoginAPIResp, err error) {
    if userid == `` || userid == `0` {
        err = errors.New(`请先登录后才能刷新状态`)
        return
    }
    if token == `` {
        err = errors.New(`请先登录后才能刷新状态`)
        return
    }

    param := apiDefaultParam(dfid, userid, token)
    ts := param.Get(`clienttime`)
    ms := ts + `123`
    p3Data := map[string]string{
        `clienttime`: ts,
        `token`:      token,
    }
    p3JSON, err := json.Marshal(p3Data)
    if err != nil {
        return
    }
    p3Str, _, err := encryptAES256CBC(p3JSON, []byte(liteKey), []byte(liteIv))
    if err != nil {
        return
    }

    paramStr, paramKey, err := encryptAES256CBC([]byte(`{}`), []byte(liteKey), []byte(liteIv))
    if err != nil {
        return
    }

    pkData := map[string]string{
        `clienttime_ms`: ms,
        `key`:           paramKey,
    }
    pkJSON, err := json.Marshal(pkData)
    if err != nil {
        return
    }

    pkStr, err := RSAPublicEncryptNoPadding(pkJSON, publicLiteRasKey)
    if err != nil {
        return
    }

    c := new(CallAPIConfig)
    c.SignType = signAndroid
    c.HTTPMethod = http.MethodPost
    c.URL = `http://login.user.kugou.com/v4/login_by_token`
    c.Param = param
    c.Header = map[string]string{
        `x-router`: `login.user.kugou.com`,
    }
    c.Body = map[string]string{
        `dfid`:          c.Param.Get(`dfid`),
        `p3`:            p3Str,
        `plat`:          `1`,
        `t1`:            `0`,
        `t2`:            `0`,
        `t3`:            `MCwwLDAsMCwwLDAsMCwwLDA=`,
        `pk`:            pkStr,
        `params`:        paramStr,
        `userid`:        userid,
        `clienttime_ms`: ms,
    }

    body, err := CallKuGouAPI(c)
    if err != nil {
        return
    }

    err = json.Unmarshal(body, &resp)
    return
}
