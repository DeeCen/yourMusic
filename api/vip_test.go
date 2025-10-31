// Package api kuGou lite API
package api

import "testing"

func TestFreeVipAll(t *testing.T) {
    err := FreeVipAll(`-`, testUserId, testUserToken)
    if err != nil {
        t.Fatalf(`TestGetSongURL error %s`, err.Error())
    }
}
