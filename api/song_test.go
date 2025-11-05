// Package api kuGou lite API
package api

import (
    "testing"
)

func TestGetSongURL(t *testing.T) {
    name, list, size, err := GetSongURL(`2mSacH2E6Lek0j4B8h1GkeHm`, testUserId, testUserToken, `0`, `0`, `42A384463CB8625B40E7396BB07C79DC`, SongQuality128)
    if err != nil {
        t.Fatalf(`TestGetSongURL error %s`, err.Error())
    }

    t.Log(`TestGetSongURL result`, name, list, size)
}
