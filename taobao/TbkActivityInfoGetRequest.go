package taobao

import (
	"encoding/json"
	"net/http"
)

/*
TbkActivityInfoGetRequest 淘宝客-推广者-官方活动转链
doc: https://open.taobao.com/api.htm?docId=48340&docType=2
*/
type TbkActivityInfoGetRequest struct {
	ActivityMaterialID string
	ADZoneID           string
	RelationID         string
	SubID              string
	UnionID            string
	result             *TbkActivityInfoGetResponse
}

type TbkActivityInfoGetResponse struct {
	TbkActivityInfoGetResponseWrap struct {
		Data struct {
			ClickURL          string `json:"click_url"`
			ShortClickURL     string `json:"short_click_url"`
			WxMiniprogramPath string `json:"wx_miniprogram_path"`
			WxQrcodeURL       string `json:"wx_qrcode_url"`
		} `json:"data"`
		RequestID string `json:"request_id"`
	} `json:"tbk_activity_info_get_response"`
}

func (r *TbkActivityInfoGetRequest) getAPIName() string {
	return "taobao.tbk.activity.info.get"
}

func (r *TbkActivityInfoGetRequest) addAPIParameters(m map[string]string) {
	if r.ActivityMaterialID != "" {
		m["activity_material_id"] = r.ActivityMaterialID
	}
	if r.ADZoneID != "" {
		m["adzone_id"] = r.ADZoneID
	}
	if r.RelationID != "" {
		m["relation_id"] = r.RelationID
	}
	if r.SubID != "" {
		m["sub_pid"] = r.SubID
	}
	if r.UnionID != "" {
		m["union_id"] = r.UnionID
	}
}

func (r *TbkActivityInfoGetRequest) handleResp(resp *http.Response) {
	result := &TbkActivityInfoGetResponse{}
	json.NewDecoder(resp.Body).Decode(result)
	r.result = result
}

/*
GetResult get api callback
*/
func (r *TbkActivityInfoGetRequest) GetResult() TbkActivityInfoGetResponse {
	return *r.result
}
