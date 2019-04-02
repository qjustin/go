package analyzer

import (
	"net/http"
	"projects/crawler/component"
)

type (
	RespParser func(httpResp *http.Response, respDepth uint32) ([]Data, []error)

	IAnalyzer interface {
		component.IComponent
		RespParsers() []RespParser
		Analyze(resp *Response) ([]Data, []error)
	}

	Analyzer struct {
		component.Component
		respParsers []RespParser
	}
)

func (this *Analyzer) RespParsers() []RespParser {
	parsers := make([]RespParser, len(this.respParsers))
	copy(parsers, this.respParsers)
	return parsers
}

func (this *Analyzer) Analyze(resp *Response) ([]Data, []error) {

}
