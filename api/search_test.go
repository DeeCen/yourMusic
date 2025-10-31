// Package api kuGou lite API
package api

import (
    "testing"
)

func TestSearchSong(t *testing.T) {
    resp, err := SearchSong(``, ``, ``, `张学友` /*我爱黎明*/, 1)
    if err != nil {
        t.Fatalf(`TestSearch error %s`, err.Error())
    }
    t.Log(`TestSearchSong result`, resp)
}

/*func TestSearch(t *testing.T) {
    resp, err := Search(``, ``, ``, `张学友` , 1)
    if err != nil {
        t.Fatalf(`TestSearch error %s`, err.Error())
    }

    t.Log(`TestSearch result`, resp)
}*/
