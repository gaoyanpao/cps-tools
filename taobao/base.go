package taobao

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	mAppKey string = ""
	mSecret string = ""
)

/*
RESTAPI 淘宝sdk
*/
type RESTAPI interface {
	getAPIName() string
	addAPIParameters(map[string]string)
	handleResp(*http.Response)
}

/*
SetAppInfo 设置全局参数
*/
func SetAppInfo(appKey string, secret string) {
	mAppKey = appKey
	mSecret = secret
}

/*
Call 调用淘宝api
*/
func Call(api RESTAPI) error {
	params := map[string]string{
		P_FORMAT:      "json",
		P_APPKEY:      mAppKey,
		P_SIGN_METHOD: "md5",
		P_VERSION:     "2.0",
		P_TIMESTAMP:   strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
		P_API:         api.getAPIName(),
		// P_PARTNER_ID:  SYSTEM_GENERATE_VERSION,
	}
	api.addAPIParameters(params)

	params[P_SIGN] = makeSign(params)
	//new request
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s", URL_TAOBAO_OPEN_API), nil)
	if err != nil {
		log.Println(err)
		return errors.New("new request is fail ")
	}
	req.URL.Path = N_REST
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "Keep-Alive")

	// TODO handle MultiPartForm

	//http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	api.handleResp(resp)
	return nil
}
