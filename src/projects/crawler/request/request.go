package request

import "net/http"

type Request struct {
	// httpReq 代表HTTP请求。
	httpReq *http.Request
	// depth 代表请求的深度。
	depth uint32
}
