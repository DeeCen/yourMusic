// Package api kuGou lite API
package api

import (
    "testing"
)

func TestSearchLyric(t *testing.T) {
    // A9343D9251A9A34B9CA49AEFE484232B
    // 833D6FD907BA55B2BE16E444F4A70F2F
    // 42A384463CB8625B40E7396BB07C79DC
    _, err := searchLyric(``, `0`, ``, `A9343D9251A9A34B9CA49AEFE484232B`)
    if err != nil {
        t.Fatalf(`TestSearchLyric err = %s`, err.Error())
    }
}

func TestGetLyric(t *testing.T) {
    _, err := GetLyric(``, ``, ``, ``)
    if err == nil {
        t.Fatalf(`TestGetLyric GetLyric hash empty NOT return error`)
    }

    _, err = GetLyric(``, ``, ``, `errorHash`)
    if err == nil {
        t.Fatalf(`TestGetLyric GetLyric hash empty NOT return error`)
    }

    resp, err := GetLyric(``, ``, ``, `42A384463CB8625B40E7396BB07C79DC`)
    if err != nil {
        t.Fatalf(`TestGetLyric GetLyric return error %s`, err.Error())
    }

    if len(resp) == 0 {
        t.Fatalf(`TestGetLyric resp empty`)
    }
}
