package downloader

import (
	"net/http"
	"projects/crawler/component"
)

type IDownloader interface {
	component.IComponent

	Downloader(req *http.Request) *http.Response
}
