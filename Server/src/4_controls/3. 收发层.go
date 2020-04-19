package controls

import (
	"0_common/commonConst"
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"io/ioutil"
	"net/http"
)


// 向前端响应
func Response(w http.ResponseWriter, rsp *ResponseProto) (err error) {
	responseBytes := make([]byte, 0)
	if responseBytes, err = json.Marshal(rsp); err != nil {
		return err
	}
	if _, err = w.Write(responseBytes); err != nil {
		return err
	}
	return nil
}
// 向前端响应错误
func ResponseError(w http.ResponseWriter, err error) error {
	if err := Response(w, &ResponseProto{
		Status: commonConst.ErrorFlag,
		Msg:    err.Error(),
	}); err != nil {
		return err
	}
	return nil
}


// 解析请求, 取得请求协议结构 (请求协议 即 请求体)
func ParseRequest(r *http.Request) (reqProto ReqProto, err error) {
	// 获取请求体字节流
	reqBodyBytes := make([]byte, 0)
	if reqBodyBytes, err = ioutil.ReadAll(r.Body); err != nil {
		return reqProto, err
	}

	// 获取请求协议结构
	reqProto = ReqProto{}
	if err = json.Unmarshal(reqBodyBytes, &reqProto); err != nil {
		return reqProto, err
	}
	return reqProto, err
}


// 获取请求中的Data
func ParseRequestData(r *http.Request,dataStructPointer interface{}) error{
	var requestBody ReqProto
	var err error
	if 	requestBody, err = ParseRequest(r);err != nil {
		return err
	}
	data, ok := make(map[string]interface{}), false
	if data, ok = requestBody.Data.(map[string]interface{}); !ok {
		return fmt.Errorf("the struct of data is error")
	}

	// 获取注册表单
	if err = mapstructure.Decode(data, dataStructPointer); err != nil {

		return err
	}
	return nil
}


// 生成图片响应Data
func FormatPhotoRspData(photoBase64 string) ([]byte,error){
	return json.Marshal(GetPhotoRspData{PhotoBase64:photoBase64})
}
