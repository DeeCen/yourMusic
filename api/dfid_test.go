// Package api kuGou lite API
package api

import "testing"

func TestGetDfid(t *testing.T) {
    ret, err := GetDfid()
    if err != nil {
        t.Fatalf(`TestGetDfid error: %s`, err.Error())
    }

    if ret == `` {
        t.Fatalf(`TestGetDfid return empty`)
    }

    t.Log(`TestGetDfid result:`, ret)
}
func TestToPostB64(t *testing.T) {
    str := `<>?12`
    want := `PD4/MTI=`

    rs := toPostB64([]byte(str))
    if rs != want {
        t.Fatalf(`toPostB64 err get %s want %s`, rs, want)
    }
}
