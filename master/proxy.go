package main

import (
	"net/http"
	"net/url"
)

// direct proxy
func proxyDirector(req *http.Request) {
	// get key from header
	targetURL := req.Header.Get("X-Target-URL")
	if targetURL == "" {

	}

	// 解析透传地址
	target, err := url.Parse(targetURL)
	if err != nil {
		http.Error(req.Response, "Failed to parse target URL", http.StatusInternalServerError)
		return
	}

	// 设置代理的目标地址
	req.URL.Scheme = target.Scheme
	req.URL.Host = target.Host
	req.URL.Path = target.Path

	// 可选：修改其他请求头
	req.Host = target.Host
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
}
