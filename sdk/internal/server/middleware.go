package server

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	v1 "sdk/api/sdk/v1"
	"sort"
	"strconv"

	"sdk/internal/service"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func MiddlewareCheckSign(sdk *service.SdkService) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if _, ok := transport.FromServerContext(ctx); ok {
				// Do something on entering
				// 中间件验证签名
				if !CheckSign(ctx, req, sdk) {
					return nil, v1.ErrorSignNotMatch("sign not match")
				}

				defer func() {
					// Do something on exiting
				}()
			}
			return handler(ctx, req)
		}
	}
}

// 验证签名
// 请求签名方式：appid + service_name + data字符串 + login_key
// service_name可以是sdk.game.initsdk或其他
// data字符串结构：url.QueryEscape(param1=a&param2=b)
// login_key是后台给每个游戏配置的唯一秘钥
func CheckSign(ctx context.Context, req interface{}, sdk *service.SdkService) bool {
	// 将interface的参数转成map
	reqJson, _ := json.Marshal(req)
	var reqMap map[string]interface{}
	json.Unmarshal([]byte(reqJson), &reqMap)

	dataStr := HttpBuildQuery(reqMap)
	makeSign := MakeSign(ctx, reqMap, dataStr, sdk)

	if fmt.Sprintf("%v", reqMap["sign"]) == makeSign {
		return true
	}
	return false
}

// data参数拼接：url.QueryEscape(param1=a&param2=b)
func HttpBuildQuery(reqMap map[string]interface{}) string {
	// 将interface的参数转成map
	dataJson, _ := json.Marshal(reqMap["data"])
	var dataMap map[string]interface{}
	json.Unmarshal([]byte(dataJson), &dataMap)

	var keys []string
	for k := range dataMap {
		keys = append(keys, k)
	}
	sort.Strings(keys) // 等同于ksort，注意首字母一致的要按照次字母排序

	dataStr := ""
	for _, k := range keys {
		if dataStr != "" {
			dataStr += "&"
		}
		switch reflect.TypeOf(dataMap[k]).String() {
		case "string":
			dataStr += fmt.Sprintf("%s=%s", k, dataMap[k])
		case "float64":
			// uint32参数会被req转成float64
			dataStr += fmt.Sprintf("%s=%s", k, strconv.FormatFloat(dataMap[k].(float64), 'f', -1, 64))
		}
	}

	return url.QueryEscape(dataStr) // urlencode()
}

// 根据map生成签名
// 当cpLoginKey=b5d6ad09709d5eac958e8fb8ad511c76时，数据和签名如下：
// {
//   "service": "sdk.game.initsdk",
//   "appId": 1000000,
//   "data": {
//     "udid": "12-34-56-78-9100",
//     "channel": 1000001
//   },
//   "sign": "ad4522e0fcad8c671d679e0e19c75968"
// }
func MakeSign(ctx context.Context, reqMap map[string]interface{}, dataStr string, sdk *service.SdkService) (signStr string) {
	// uint32参数会被req转成float64
	signStr += fmt.Sprintf("%s", strconv.FormatFloat(reqMap["appId"].(float64), 'f', -1, 64))
	signStr += fmt.Sprintf("%s", reqMap["service"])
	signStr += dataStr
	// signStr += "b5d6ad09709d5eac958e8fb8ad511c76" // TODO: 从db/cache获取
	cpLoginKey, err := GetLoginKey(ctx, reqMap, sdk)

	if err != nil {
		panic("game unavailable.")
	}
	signStr += cpLoginKey
	return fmt.Sprintf("%x", md5.Sum([]byte(signStr)))
}

func GetLoginKey(ctx context.Context, reqMap map[string]interface{}, sdk *service.SdkService) (string, error) {
	res, err := sdk.GetGameInfo(ctx, uint32(reqMap["appId"].(float64)))
	return res.CpLoginKey, err
}
