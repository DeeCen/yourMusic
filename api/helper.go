// Package api kuGou lite API
package api

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/md5"
    "crypto/rsa"
    "crypto/x509"
    "encoding/hex"
    "encoding/json"
    "encoding/pem"
    "errors"
    "math/big"
    "net/http"
    "net/url"
    "sort"
    "strconv"
    "strings"
    "time"
)

const signKey = `LnT6xpN3khm36zse0QzvmgTZ3waWdRSA` // lite 版本
const signAppId = `3116`                           // lite appid
const signKeyVer = `11040`                         // lite 版本
const publicLiteRasKey = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDECi0Np2UR87scwrvTr72L6oO01rBbbBPriSDFPxr3Z5syug0O24QyQO8bg27+0+4kBzTBTBOZ/WWU0WryL1JSXRTXLgFVxtzIY41Pe7lPOgsfTCn5kZcvKhYKJesKnnJDNr5/abvTGf+rHG3YRwsCHcQ08/q6ifSioBszvb3QiwIDAQAB\n-----END PUBLIC KEY-----"

// SignatureAndroidParams 生成android lite版本签名
func SignatureAndroidParams(paramsString, jsonBody string) (ret string) {
    signStr := signKey + paramsString + jsonBody + signKey
    ret = md5Str(signStr)
    return
}

// SignatureRegisterParams 生成Register版本签名
func SignatureRegisterParams(paramsString string, appId string) (ret string) {
    signStr := appId + paramsString + appId
    ret = md5Str(signStr)
    return
}

// SignKeyA 请求的key签名1
func SignKeyA(str string) (ret string) {
    ret = md5Str(signAppId + signKey + signKeyVer + str)
    return
}

// SignKeyB 请求的key签名2
func SignKeyB(hash, appid, mid, userid string) (ret string) {
    ret = md5Str(hash + `185672dd44712f60bb1736df5a377e82` + appid + mid + userid)
    return
}

// RSAPublicEncryptNoPadding RSA 公钥加密（无填充）
// 对应 Node.js: crypto.publicEncrypt({ key: publicKey, padding: crypto.constants.RSA_NO_PADDING })
func RSAPublicEncryptNoPadding(dataBytes []byte, publicKeyPEM string) (ret string, err error) {
    block, _ := pem.Decode([]byte(publicKeyPEM))
    if block == nil {
        err = errors.New(`failed to parse PEM block containing public key`)
        return
    }

    pub, e := x509.ParsePKIXPublicKey(block.Bytes)
    if e != nil {
        err = errors.New("failed to parse public key: " + e.Error())
        return
    }

    rsaPub, ok := pub.(*rsa.PublicKey)
    if !ok {
        err = errors.New(`not a valid RSA public key`)
        return
    }

    keySize := rsaPub.Size()
    dataBytesLen := len(dataBytes)
    if dataBytesLen > keySize {
        err = errors.New(`data too large for RSA key size: got ` + strconv.Itoa(dataBytesLen) + ` max ` + strconv.Itoa(keySize))
        return
    }

    if len(dataBytes) < keySize {
        paddedData := make([]byte, keySize)
        copy(paddedData, dataBytes)
        dataBytes = paddedData
    }

    encrypted, err := rsaEncryptNoPadding(rsaPub, dataBytes)
    if err != nil {
        return
    }

    ret = hex.EncodeToString(encrypted)
    return
}

func rsaEncryptNoPadding(pub *rsa.PublicKey, data []byte) (ret []byte, err error) {
    k := (pub.N.BitLen() + 7) / 8
    if len(data) > k {
        err = errors.New(`data too large`)
        return
    }

    m := new(big.Int).SetBytes(data)
    if m.Cmp(pub.N) >= 0 {
        err = errors.New(`data too large for modulus`)
        return
    }

    c := new(big.Int).Exp(m, big.NewInt(int64(pub.E)), pub.N)
    ret = c.Bytes()

    if n := len(ret); n < k {
        ret = append(make([]byte, k-n), ret...)
    }

    return
}

func encryptAES256CBC(data []byte, key []byte, iv []byte) (ret, keyStr string, err error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return
    }

    // 填充明文以满足块大小
    dataNew := pkcs7Pad(data, block.BlockSize())
    ciphertext := make([]byte, len(dataNew))
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext, dataNew)

    ret = hex.EncodeToString(ciphertext)
    keyStr = hex.EncodeToString(key)
    return
}

// PKCS7 填充
func pkcs7Pad(data []byte, blockSize int) []byte {
    padding := blockSize - len(data)%blockSize
    padText := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(data, padText...)
}

// RequestParam API请求参数封装
type RequestParam map[string]string

// Add 添加key,val参数
func (r *RequestParam) Add(k, v string) {
    (*r)[k] = v
}

// Get 获取指定key的val
func (r *RequestParam) Get(k string) (ret string) {
    ret = (*r)[k]
    return
}

// ToStr 转为加密顺序字符串
func (r *RequestParam) toStr(isURLVal bool) (ret string) {
    // 获取所有键并排序
    keys := make([]string, 0, len(*r))
    for key := range *r {
        keys = append(keys, key)
    }
    sort.Strings(keys)

    // 构建参数字符串
    buff := bytes.Buffer{}
    for _, key := range keys {
        val := (*r)[key]
        if isURLVal {
            val = url.QueryEscape(val) + `&`
        }
        buff.WriteString(key + `=` + val)
    }
    ret = buff.String()
    return
}

// ToSignStr 转为加密字符串
func (r *RequestParam) ToSignStr() (ret string) {
    ret = r.toStr(false)
    return
}

// ToURLStr 转为拼接到url尾部的参数字符串
func (r *RequestParam) ToURLStr() (ret string) {
    ret = r.toStr(true)
    return
}

// ToRegisterSignStr 转为Register加密字符串
func (r *RequestParam) ToRegisterSignStr() (ret string) {
    // 获取所有键并排序
    val := make([]string, 0, len(*r))
    for _, v := range *r {
        val = append(val, v)
    }
    sort.Strings(val)
    ret = strings.Join(val, ``)
    return
}

// RespErrorCode 获取接口错误码字段结构体
type RespErrorCode struct {
    ErrorCode int `json:"error_code,omitempty"`
    Data      any `json:"data,omitempty"`
}

func getErrorCodeFromBody(body []byte) (errorCode int, errorMsg string) {
    errorCode = 404
    if len(body) <= 10 {
        return
    }

    resp := new(RespErrorCode)
    if e := json.Unmarshal(body, resp); e == nil {
        errorCode = resp.ErrorCode
        if msg, ok := resp.Data.(string); ok {
            errorMsg = msg
        }
    }
    return
}

type signatureType uint

const (
    signAndroid  signatureType = 1
    signRegister signatureType = 2
)

// CallAPIConfig 调用api配置
type CallAPIConfig struct {
    SignType   signatureType     // 添加签名signature的方式
    HTTPMethod string            // 请求方式 GET,POST
    URL        string            // 请求完整地址
    BodyRaw    string            // 请求body原生数据,用于不使用json传输的情况,优先级比Body高
    Body       map[string]string // 请求body数据,传输时会转为json
    Header     map[string]string // 请求header数据
    Param      RequestParam      // 请求url参数
}

// CallKuGouAPI 统一对外调用接口函数
func CallKuGouAPI(c *CallAPIConfig) (body []byte, err error) {
    if c == nil {
        err = errors.New(`CallAPIConfig empty`)
        return
    }

    if c.Param == nil {
        err = errors.New(`CallAPIConfig Param empty`)
        return
    }

    if c.URL == `` {
        err = errors.New(`CallAPIConfig URL empty`)
        return
    }

    postStr := ``
    if c.Body != nil {
        c.HTTPMethod = http.MethodPost
        postByte, _ := json.Marshal(c.Body)
        postStr = string(postByte)
    }

    // 高优先级的 BodyRaw
    if c.BodyRaw != `` {
        postStr = c.BodyRaw
    }

    switch c.SignType {
    case signAndroid:
        signature := SignatureAndroidParams(c.Param.ToSignStr(), postStr)
        c.Param.Add(`signature`, signature)

    case signRegister:
        signature := SignatureRegisterParams(c.Param.ToRegisterSignStr(), c.Param.Get(`appid`))
        c.Param.Add(`signature`, signature)
    }

    urlStr := c.URL + `?` + c.Param.ToURLStr()
    header := apiDefaultHeader(c.Param, c.Header)

    statusCode := 0
    switch c.HTTPMethod {
    case http.MethodGet:
        statusCode, body, _, err = HTTPGet(urlStr, header, time.Second*5)
    case http.MethodPost:
        statusCode, body, _, err = HTTPPostJSON(urlStr, postStr, header, time.Second*5)
    default:
        err = errors.New(`HTTPMethod error:` + c.HTTPMethod)
    }

    if err != nil {
        return
    }

    if statusCode < 200 || statusCode >= 400 {
        err = errors.New(`http error:` + strconv.Itoa(statusCode))
        return
    }

    // 所有接口不返回空数据 ?
    if len(body) == 0 {
        err = errors.New(`http body empty`)
        return
    }

    // 所有接口均返回 error_code ?
    errorCode, errorMsg := getErrorCodeFromBody(body)
    if errorCode != 0 {
        err = errors.New(strconv.Itoa(errorCode) + `:` + errorMsg)
        return
    }

    return
}

func md5Str(s string) (ret string) {
    hash := md5.Sum([]byte(s))
    ret = hex.EncodeToString(hash[:])
    return
}

func tsStr() (ret string) {
    ret = strconv.Itoa(int(time.Now().Unix()))
    return
}

func msStr() (ret string) {
    ret = strconv.Itoa(int(time.Now().UnixMilli()))
    return
}

func apiDefaultParam(dfid, userid, token string) (ret RequestParam) {
    if dfid == `` {
        dfid = `-`
    }
    if userid == `` {
        userid = `0`
    }

    ret = make(RequestParam, 10)
    ret.Add(`mid`, md5Str(dfid))
    ret.Add(`uuid`, `-`)
    ret.Add(`appid`, signAppId)
    ret.Add(`userid`, userid)
    ret.Add(`dfid`, dfid)
    ret.Add(`clientver`, signKeyVer)
    ret.Add(`clienttime`, tsStr())
    if token != `` {
        ret.Add(`token`, token)
    }

    return
}

func apiDefaultHeader(r RequestParam, headerExt map[string]string) (ret map[string]string) {
    if r == nil {
        return
    }

    ret = make(map[string]string, 4)
    ret[`dfid`] = r.Get(`dfid`)
    ret[`clienttime`] = r.Get(`clienttime`)
    ret[`mid`] = r.Get(`mid`)
    ret[`User-Agent`] = `Android15-1070-` + signKeyVer + `-201-0-LOGIN-wifi`
    for k := range headerExt {
        ret[k] = headerExt[k]
    }

    return
}
