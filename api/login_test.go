// Package api kuGou lite API
package api

import (
    "testing"
)

const testMobile = `13800138001`
const testUserId = `0`
const testUserToken = ``

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
