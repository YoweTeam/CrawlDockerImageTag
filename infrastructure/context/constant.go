package context

const (
	DEFUALT_PREFIX          = "x-"
	REQUEST_TRACE_KEY       = DEFUALT_PREFIX + "trace-id"
	BIZ_FLOW_TRACE_KEY      = DEFUALT_PREFIX + "biz-id"
	BIZ_USER_KEY            = DEFUALT_PREFIX + "user"
	BIZ_USERID_KEY          = DEFUALT_PREFIX + "user-id"
	BIZ_USERNAME_KEY        = DEFUALT_PREFIX + "user-name"
	MESH_HEADERS_CONFIG_KEY = "mesh.headers"
	LANG_KEY                = DEFUALT_PREFIX + "lang"
)

var DEFAULT_MESH_HEADERS_KEYS = []string{"x-request-id", "x-request-real-ip", "x-request-real-ip"}

// refer: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Language
var DEFAULT_LANG = "zh-CN"
