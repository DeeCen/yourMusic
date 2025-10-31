// Package api kuGou lite API
package api

import (
    "testing"
)

const testMobile = `13800138001`
const testUserId = `2089100744`
const testUserToken = `387accb2c7a60ae8573bcb8d4c23839593ae2074639ed28569ed8df2a45f7a37`

func TestSendMobileCode(t *testing.T) {
    err := SendMobileCode(testMobile)
    if err != nil {
        t.Fatalf(`SendMobileCode error %s`, err.Error())
    }
}

func TestLoginByVerifyCode(t *testing.T) {
    resp, err := LoginByVerifyCode(testMobile, `123456`)
    if err != nil {
        t.Fatalf(`TestLoginByVerifyCode error %s`, err.Error())
    }

    t.Log(`token=`, resp.Data.Token)
}

func TestLoginRefresh(t *testing.T) {
    ret, err := LoginRefresh(``, testUserId, testUserToken)
    if err != nil {
        t.Fatalf(`TestLoginByVerifyCode error %s`, err.Error())
    }

    t.Log(`token=`, ret.Data.Token)
}
