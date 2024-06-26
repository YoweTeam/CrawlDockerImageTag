package context

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// Context interface
type Context interface {
	context.Context
	GetRawContext() interface{}
	GetProtocol() ProtocolType
	GetMeshTrace() MeshTrace
	GetRequestTrace() RequestTrace
	GetBizFlowTrace() BizFlowTrace
	GetUserInfo() UserInfo
	GetLang() LangInfo
	Get(key string) (value interface{}, exist bool)
	Set(key string, value interface{})
}

// ProtocolType int
type ProtocolType int32

// LangInfo obj
type LangInfo struct {
	Code string `json:"code"`
}

const (
	_ ProtocolType = iota
	Gin
	Grpc
	Test
	MQ
)

func (p ProtocolType) String() string {
	switch p {
	case Gin:
		return "GIN"
	case Grpc:
		return "GRPC"
	case Test:
		return "TEST"
	case MQ:
		return "MQ"
	default:
		return "UNKNOWN"
	}
}

type BizFlowTrace string

type RequestTrace string

type MeshTrace map[string]string

type UserInfo struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func ToMap(ctx Context) map[string]string {
	var userInfo = ctx.GetUserInfo()
	var fields = make(map[string]string)
	fields[BIZ_USERID_KEY] = userInfo.UserId
	fields[BIZ_USERNAME_KEY] = userInfo.UserName
	fields[REQUEST_TRACE_KEY] = string(ctx.GetRequestTrace())
	fields[BIZ_FLOW_TRACE_KEY] = string(ctx.GetBizFlowTrace())
	fields[LANG_KEY] = ctx.GetLang().Code
	apm := ctx.GetMeshTrace()
	meshKeys := GetMeshKeys()
	for _, k := range meshKeys {
		if apm != nil {
			fields[k] = apm[k]
		} else {
			fields[k] = ""
		}
	}
	return fields
}

func GetMeshKeys() []string {
	var headerKeys []string
	tmp := viper.GetString(MESH_HEADERS_CONFIG_KEY)
	if tmp != "" {
		headerKeys = strings.Split(tmp, ",")
	} else {
		headerKeys = append(headerKeys, DEFAULT_MESH_HEADERS_KEYS...)
	}
	return headerKeys
}

// NewContext new
func NewContext(origin Context) Context {
	ctx := NewEmptyContext()
	ctx.Trace = origin.GetRequestTrace()
	ctx.Biz = BizFlowTrace(uuid.New().String())
	ctx.User = UserInfo{
		UserId:   uuid.New().String(),
		UserName: uuid.New().String(),
	}
	apm := make(map[string]string)
	for _, v := range GetMeshKeys() {
		apm[v] = uuid.New().String()
	}
	ctx.Mesh = apm
	ctx.Lang = LangInfo{Code: DEFAULT_LANG}
	return ctx
}

// NewBackgroundContext new
func NewBackgroundContext() Context {
	ctx := NewEmptyContext()
	ctx.Trace = RequestTrace(uuid.New().String())
	ctx.Biz = BizFlowTrace(uuid.New().String())
	ctx.User = UserInfo{
		UserId:   uuid.New().String(),
		UserName: uuid.New().String(),
	}
	apm := make(map[string]string)
	for _, v := range GetMeshKeys() {
		apm[v] = uuid.New().String()
	}
	ctx.Mesh = apm
	ctx.Lang = LangInfo{Code: DEFAULT_LANG}
	return ctx
}
