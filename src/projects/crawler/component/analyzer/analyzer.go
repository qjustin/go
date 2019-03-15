package analyzer

import (
	"net/http"
)

type Parser func(httpResp *http.Response, respDepth uint32) ([]Data, []error)

type IAnalyzer interface {
	IComponent

	GetParsers() []Parser
}

type Analyzer struct {
}
