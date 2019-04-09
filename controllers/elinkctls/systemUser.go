package elinkctls

import (
	"github.com/thinkgos/gogate/models"
	"github.com/thinkgos/gogate/protocol/elinkch/ctrl"
	"github.com/thinkgos/gomo/elink"

	"github.com/json-iterator/go"
)

type SysMultiUserPy struct {
	Uid []int64 `json:"uid"`
}

type SysUserPy struct {
	Uid int64 `json:"uid"`
}

type SysMultiUserRequest struct {
	ctrl.BaseRequest
	Payload SysMultiUserPy `json:"payload,omitempty"`
}

type SysUserController struct {
	ctrl.Controller
}

func (this *SysUserController) Get() {
	out, err := jsoniter.Marshal(SysMultiUserPy{Uid: models.GetUsers()})
	if err != nil {
		this.ErrorResponse(elink.CodeErrSysOperationFailed)
		return
	}
	this.WriteResponse(elink.CodeSuccess, out)
}

func (this *SysUserController) Post() {
	this.userDeal(false)
}

func (this *SysUserController) Delete() {
	this.userDeal(true)
}

func (this *SysUserController) userDeal(isDel bool) {
	var uid []int64
	var isArray bool

	code := elink.CodeSuccess
	defer func() {
		this.ErrorResponse(code)
	}()

	juid := jsoniter.Get(this.Input.Payload, "payload", "uid")
	switch juid.ValueType() {
	case jsoniter.NumberValue:
		uid = append(uid, juid.ToInt64())
	case jsoniter.ArrayValue:
		isArray = true
		req := &SysMultiUserRequest{}
		if err := jsoniter.Unmarshal(this.Input.Payload, req); err != nil {
			code = elink.CodeErrSysInvalidParameter
			return
		}
		uid = req.Payload.Uid
	default:
		code = elink.CodeErrSysInvalidParameter
		return
	}

	uidSuc := []int64{}
	for _, v := range uid {
		if isDel {
			if err := models.DeleteUser(v); err != nil {
				continue
			}
		} else {
			if err := models.AddUser(v); err != nil {
				continue
			}
		}
		uidSuc = append(uidSuc, v)
	}
	if len(uidSuc) == 0 {
		code = elink.CodeErrSysOperationFailed
		return
	}
	var out []byte
	var err error
	if isArray {
		out, err = jsoniter.Marshal(SysMultiUserPy{Uid: uidSuc})
		if err != nil {
			code = elink.CodeErrSysOperationFailed
			return
		}
	} else {
		out, err = jsoniter.Marshal(SysUserPy{Uid: uidSuc[0]})
		if err != nil {
			code = elink.CodeErrSysOperationFailed
			return
		}
	}

	this.WriteResponse(elink.CodeSuccess, out)
}