// Package api kuGou lite API
package api

import "testing"

func TestSignatureAndroidParams(t *testing.T) {
    param := `appid=3116clienttime=1761648408clientver=11040dfid=-mid=329346087999717826199985607920306107310token=387accb2c7a60ae8573bcb8d4c2383955b3e190c2605d219c40a8dfa683d8082userid=0uuid=7fa82cf6476e2150aced01a50fdfc52b`
    jsonBody := `{"plat":1,"support_multi":1,"t1":0,"t2":0,"clienttime_ms":1761648408024,"mobile":"13800138000","key":"41a1f6292006fb1496b9f5fa3d44ad94","p2":"B1B6B9957209734ACC26F41B8201AD51BB9B426AE080203E1E22DCBB0FEE420AB1CB36E9FFA815E231C2733C8AB17B526964AE13BEBA48BF964A44D0FFAE5E5B22EC1D27C5C62096905A8D4CEF097030961ED792B146B563C84B6E30605576DC76C328DD8975E596992F3304D7A36D66AF0352B834BA1AD32EF93AE185F8485A"}`
    rs := SignatureAndroidParams(param, jsonBody)
    want := `0037838cfff4f6c1d3ed0dfce577cb4e`
    if rs != want {
        t.Fatalf(`SignatureAndroidParams check fail %s want %s`, rs, want)
    }
}

func TestRequestParam_ToSignStr(t *testing.T) {
    p := make(RequestParam)
    p.Add(`appid`, `3116`)
    p.Add(`clienttime`, `1761660803`)
    p.Add(`clientver`, `11040`)
    p.Add(`dfid`, `-`)
    p.Add(`mid`, `336d5ebc5436534e61d16e63ddfca327`)
    p.Add(`userid`, `0`)
    p.Add(`uuid`, `15e772e1213bdd0718d0c1d10d64e06f`)

    rs := p.ToSignStr()
    if rs != `appid=3116clienttime=1761660803clientver=11040dfid=-mid=336d5ebc5436534e61d16e63ddfca327userid=0uuid=15e772e1213bdd0718d0c1d10d64e06f` {
        t.Fatal(`TestRequestParam_ToStr error: ` + rs)
    }
}

func TestSignKeyA(t *testing.T) {
    rs := SignKeyA(`1761647922205`)
    want := `b8c0e667eb5165eddb72cd808b7a830b`
    if rs != want {
        t.Fatalf(`SignKeyA check fail %s want %s`, rs, want)
    }
}

func TestSignKeyB(t *testing.T) {
    rs := SignKeyB(`a855e67454ac2e6344fc90382d8949ce`, signAppId, `329346087999717826199985607920306107310`, `0`)
    want := `4ee8ded4945c4829e345a5d40f8c8e48`
    if rs != want {
        t.Fatalf(`SignKeyB check fail %s want %s`, rs, want)
    }
}

func TestRSAPublicEncryptNoPadding(t *testing.T) {
    data := []byte(`{"clienttime_ms":1761650648071,"code":"1234","mobile":"13800138000"}`)
    rs, err := RSAPublicEncryptNoPadding(data, publicLiteRasKey)
    want := `3f385e30264241447e0fae040514957be7044023ded5990e3f2f984b1f063d38f2044a039b05a12999ca1783a281e189f401bd7e5ea9a418035cae39d360c54da629e8939fdf50641590f1ac39bdaf4469cbafe4a8ac98e26bb876f92c04e611988fd1d2bcfbf919f9e14189ccdcdeb0f46da83b2d7401767fbcf90faae03e43`
    if err != nil || rs != want {
        t.Fatalf(`RSAPublicEncryptNoPadding check fail %s want %s err=%v`, rs, want, err)
    }
}
