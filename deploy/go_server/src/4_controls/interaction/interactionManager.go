package interaction

import (
	"0_common/commonConst"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/goinggo/mapstructure"
	"io/ioutil"
	"net/http"
)


type interactionManager struct {
}

var itm = &interactionManager{}

func GetInteractionManger() *interactionManager{
	return itm
}

// 向前端响应
func (im *interactionManager)Response(w http.ResponseWriter, rsp *ResponseProto) {
	var responseBytes []byte
	var err error
	if responseBytes, err = json.Marshal(rsp); err != nil {
		panic(err)
	}
	if _, err = w.Write(responseBytes); err != nil {
		panic(err)
	}
}


func (im *interactionManager)ResponseError(w http.ResponseWriter, err error) {
	im.Response(w, &ResponseProto{
		Status: commonConst.ErrorFlag,
		Msg:    err.Error(),
	})
}


func (im *interactionManager)ResponseSuccess(w http.ResponseWriter, successMsg string) {
	im.Response(w, &ResponseProto{
		Status: commonConst.SuccessFlag,
		Msg:    successMsg,
	})
}



// 解析请求, 取得请求协议结构 (请求协议 即 请求体)
func parseRequest(r *http.Request) (reqProto ReqProto, err error) {
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
func (im *interactionManager)GetDataFormRequest(r *http.Request, dataPointer interface{}) error {

	// 获取请求结构体
	var requestBody ReqProto
	var err error
	if requestBody, err = parseRequest(r); err != nil {
		logs.Error(err)
		return err
	}

	// 获取请求结构体的 data
	var data map[string]interface{}
	var ok bool
	if data, ok = requestBody.Data.(map[string]interface{}); !ok {
		logs.Error(err)
		return err
	}

	// 通过data，形成相应的结构体
	if err = mapstructure.Decode(data, dataPointer); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// 生成图片响应Data
func (im *interactionManager)GetPhotoRspData(photoBase64 string) ([]byte, error) {
	var bytes []byte
	var err error
	if bytes, err = json.Marshal(GetPhotoRspDTO{PhotoBase64: photoBase64});err != nil {
		logs.Error(err)
		return nil,err
	}
	return bytes, nil
}
