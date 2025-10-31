// Package api kuGou lite API
package api

import (
    "errors"
    "net/http"
    "strconv"
    "strings"
    "time"
)

func adDayVip(dfid, userid, token string) (err error) {
    msNow := time.Now().UnixMilli()
    postData := map[string]string{
        `ad_id`:      `12307537187`,
        `play_end`:   strconv.FormatInt(msNow, 10),
        `play_start`: strconv.FormatInt(msNow-60000, 10),
    }

    c := new(CallAPIConfig)
    c.SignType = signAndroid
    c.HTTPMethod = http.MethodPost
    c.URL = `https://gateway.kugou.com/youth/v1/ad/play_report`
    c.Param = apiDefaultParam(dfid, userid, token)
    c.Body = postData

    _, err = CallKuGouAPI(c)
    return
}
func freeDayVip(dfid, userid, token string) (err error) {
    param := apiDefaultParam(dfid, userid, token)
    param.Add(`source_id`, `90137`)

    c := new(CallAPIConfig)
    c.SignType = signAndroid
    c.HTTPMethod = http.MethodPost
    c.URL = `https://gateway.kugou.com/youth/v1/recharge/receive_vip_listen_song`
    c.Param = param

    _, err = CallKuGouAPI(c)
    return
}

func freeDayVipUpgrade(dfid, userid, token string) (err error) {
    param := apiDefaultParam(dfid, userid, token)
    param.Add(`kugouid`, userid)
    param.Add(`ad_type`, `1`)

    c := new(CallAPIConfig)
    c.SignType = signAndroid
    c.HTTPMethod = http.MethodPost
    c.URL = `https://gateway.kugou.com/youth/v1/listen_song/upgrade_vip_reward`
    c.Param = param

    _, err = CallKuGouAPI(c)
    return
}

// FreeVipAll 让我想起了我最喜欢的音乐app-虾米: https://www.zhihu.com/question/68347374/answer/262267626
func FreeVipAll(dfid, userid, token string) (err error) {
    if userid == `` || userid == `0` || token == `` {
        err = errors.New(`请先登录后再领取vip`)
        return
    }

    e1 := adDayVip(dfid, userid, token)
    e2 := freeDayVip(dfid, userid, token)
    e3 := freeDayVipUpgrade(dfid, userid, token)

    // 任一成功
    if e1 == nil || e2 == nil || e3 == nil {
        return
    }

    err = e3
    if e1 != nil {
        err = e1
        return
    }
    if e2 != nil {
        err = e2
        return
    }

    // 297002 重复领取
    if err != nil && strings.Contains(err.Error(), `297002`) {
        err = nil
    }

    return
}
