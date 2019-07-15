package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"sort"
)

func getMapKeys(m map[string]string) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func sortKeys(keys []string) []string {
	sort.Strings(keys)
	return keys
}

type Sign struct {
	AccessKeyId      string
	AccessKeySecret  string
	SignatureMethod  string
	SignatureVersion string
}

// NewSign 创建签名参数
func NewSign(accessKeyId, accessKeySecret string) *Sign {
	return &Sign{
		AccessKeyId:      accessKeyId,
		AccessKeySecret:  accessKeySecret,
		SignatureMethod:  "HmacSHA256",
		SignatureVersion: "2",
	}
}

// Get 获取签名
func (s *Sign) Get(method, host, path, timestamp string, params map[string]string) string {
	var str = method + "\n" + host + "\n" + path + "\n"
	params["AccessKeyId"] = s.AccessKeyId
	params["SignatureMethod"] = s.SignatureMethod
	params["SignatureVersion"] = s.SignatureVersion
	params["Timestamp"] = timestamp
	mapCloned := make(map[string]string)
	for key, value := range params {
		mapCloned[key] = url.QueryEscape(value)
	}
	strParams := Map2UrlQueryBySort(mapCloned)
	strPayload := str + strParams
	return ComputeHmac256(strPayload, s.AccessKeySecret)
}

// MapSortByKey 对Map按着ASCII码进行排序
// mapValue: 需要进行排序的map
// return: 排序后的map
func MapSortByKey(mapValue map[string]string) map[string]string {
	var keys []string
	for key := range mapValue {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	mapReturn := make(map[string]string)
	for _, key := range keys {
		mapReturn[key] = mapValue[key]
	}

	return mapReturn
}

// MapValueEncodeURI 对Map的值进行URI编码
// mapParams: 需要进行URI编码的map
// return: 编码后的map
func MapValueEncodeURI(mapValue map[string]string) map[string]string {
	for key, value := range mapValue {
		valueEncodeURI := url.QueryEscape(value)
		mapValue[key] = valueEncodeURI
	}

	return mapValue
}

// Map2UrlQuery 将map格式的请求参数转换为字符串格式的
// mapParams: map格式的参数键值对
// return: 查询字符串
func Map2UrlQuery(mapParams map[string]string) string {
	var strParams string
	for key, value := range mapParams {
		strParams += (key + "=" + value + "&")
	}

	if 0 < len(strParams) {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}

	return strParams
}

// Map2UrlQueryBySort 将map格式的请求参数转换为字符串格式的,并按照Map的key升序排列
// mapParams: map格式的参数键值对
// return: 查询字符串
func Map2UrlQueryBySort(mapParams map[string]string) string {
	var keys []string
	for key := range mapParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var strParams string
	for _, key := range keys {
		strParams += key + "=" + mapParams[key] + "&"
	}

	if 0 < len(strParams) {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}

	return strParams
}

// ComputeHmac256 HMAC SHA256加密
// strMessage: 需要加密的信息
// strSecret: 密钥
// return: BASE64编码的密文
func ComputeHmac256(strMessage string, strSecret string) string {
	key := []byte(strSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strMessage))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
