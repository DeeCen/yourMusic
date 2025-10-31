// Package api kuGou lite API
package api

import (
    "encoding/base64"
    "encoding/json"
    "net/http"
)

type respAPIDfid struct {
    Data struct {
        Dfid string `json:"dfid,omitempty"`
    } `json:"data,omitempty"`
}

// GetDfid 获取dfid
// 不知道有什么用,因为接口使用 - 作为dfid也是可以的, 这里为了保持统一就获取真的 dfid
func GetDfid() (ret string, err error) {
    param := apiDefaultParam(``, ``, ``)
    param.Add(`appid`, `1014`)
    param.Add(`p.token`, ``)
    param.Add(`platid`, `4`)

    dataB64 := map[string]string{
        `mid`:    param.Get(`mid`),
        `uuid`:   param.Get(`uuid`),
        `appid`:  param.Get(`appid`),
        `userid`: param.Get(`userid`),
    }

    dataB64JSON, err := json.Marshal(dataB64)
    if err != nil {
        return
    }

    c := new(CallAPIConfig)
    c.SignType = signRegister
    c.HTTPMethod = http.MethodPost
    c.URL = `https://userservice.kugou.com/risk/v1/r_register_dev`
    c.Param = param
    c.BodyRaw = toPostB64(dataB64JSON)

    body, err := CallKuGouAPI(c)
    if err != nil {
        return
    }

    resp := new(respAPIDfid)
    err = json.Unmarshal(body, resp)
    if err != nil {
        return
    }

    ret = resp.Data.Dfid
    return
}

func toPostB64(data []byte) (ret string) {
    ret = base64.StdEncoding.EncodeToString(data)
    return
}
