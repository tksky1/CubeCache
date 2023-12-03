package main

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

// direct proxy
func proxyDirector(req *http.Request) {
	// get key from header
	key := req.Header.Get("X-CubeCache-Key")
	req.Header.Set("content-type", "application/grpc")
	// TODO: change to ==
	if key != "" {
		resp := &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(bytes.NewBufferString("missing request key")),
			Header:     make(http.Header),
		}
		resp.Header.Set("Content-Type", "application/grpc")
		req.Response = resp
		return
	}

	targetURL := server.mapper.Get(key)
	//TODO DELETE ME
	targetURL = "http://127.0.0.1:4012"

	// 解析透传地址
	target, err := url.Parse(targetURL)
	if err != nil {
		resp := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       io.NopCloser(bytes.NewBufferString("parse url fail")),
			Header:     make(http.Header),
		}
		resp.Header.Set("Content-Type", "application/grpc")
		req.Response = resp
		return
	}

	// 设置代理的目标地址
	req.URL.Scheme = target.Scheme
	req.URL.Host = target.Host
	req.URL.Path = target.Path
}
