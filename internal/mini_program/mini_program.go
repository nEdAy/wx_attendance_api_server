package mini_program

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"github.com/nEdAy/wx_attendance_api_server/util"
)

const (
	// WeChat BaseURL 微信请求基础URL
	weChatBaseURL    = "https://api.weixin.qq.com"
	codeToSessionAPI = "/sns/jscode2session"
)

const (
	// weChatServerError 微信服务器错误时返回返回消息
	weChatServerError = "微信服务器发生错误"
)

// code 微信服务器返回状态码
type code int

// response 请求微信返回基础数据
type response struct {
	Errcode code   `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type loginForm struct {
	response
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

// phoneNumber 解密后的用户手机号码信息
type phoneNumber struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	Watermark       watermark `json:"watermark"`
}

// userInfo 解密后的用户信息
type userInfo struct {
	OpenID    string    `json:"openId"`
	Nickname  string    `json:"nickName"`
	Gender    int       `json:"gender"`
	Province  string    `json:"province"`
	Language  string    `json:"language"`
	Country   string    `json:"country"`
	City      string    `json:"city"`
	Avatar    string    `json:"avatarUrl"`
	UnionID   string    `json:"unionId"`
	Watermark watermark `json:"watermark"`
}

// Login 用户登录
// 返回 微信端 openid 和 session_key
func Login(appID, secret, code string) (string, string, error) {
	if code == "" {
		return "", "", errors.New("code can not be null")
	}

	api, err := code2url(appID, secret, code)
	if err != nil {
		return "", "", err
	}

	res, err := http.Get(api)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", "", errors.New(weChatServerError)
	}

	var data loginForm

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return "", "", err
	}

	if data.Errcode != 0 {
		return "", "", errors.New(data.Errmsg)
	}

	return data.Openid, data.SessionKey, nil
}

type watermark struct {
	AppID     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

// DecryptPhoneNumber 解密手机号码
//
// @ssk 通过 Login 向微信服务端请求得到的 session_key
// @data 小程序通过 api 得到的加密数据(encryptedData)
// @iv 小程序通过 api 得到的初始向量(iv)
func DecryptPhoneNumber(ssk, data, iv string) (phone phoneNumber, err error) {
	bts, err := util.CBCDecrypt(ssk, data, iv)
	if err != nil {
		return
	}

	err = json.Unmarshal(bts, &phone)
	return
}

// DecryptUserInfo 解密用户信息
//
// @rawData 不包括敏感信息的原始数据字符串，用于计算签名。
// @encryptedData 包括敏感数据在内的完整用户信息的加密数据
// @signature 使用 sha1( rawData + session_key ) 得到字符串，用于校验用户信息
// @iv 加密算法的初始向量
// @ssk 微信 session_key
func DecryptUserInfo(rawData, encryptedData, signature, iv, ssk string) (ui userInfo, err error) {
	if ok := util.Validate(rawData, ssk, signature); !ok {
		err = errors.New("数据校验失败")
		return
	}

	bts, err := util.CBCDecrypt(ssk, encryptedData, iv)
	if err != nil {
		return
	}

	err = json.Unmarshal(bts, &ui)
	return
}

// 拼接 获取 session_key 的 URL
func code2url(appID, secret, code string) (string, error) {
	codeToSessionUrl, err := url.Parse(weChatBaseURL + codeToSessionAPI)
	if err != nil {
		return "", err
	}

	query := codeToSessionUrl.Query()

	query.Set("appid", appID)
	query.Set("secret", secret)
	query.Set("js_code", code)
	query.Set("grant_type", "authorization_code")

	codeToSessionUrl.RawQuery = query.Encode()

	return codeToSessionUrl.String(), nil
}
