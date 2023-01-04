package live

type PullLiveInfo struct {
	// RTMP协议拉流地址
	RtmpUrl string `json:"rtmp_url"`
	// RTS协议拉流地址
	RtsUrl string `json:"rts_url"`
	// FLV协议拉流地址
	FlvUrl string `json:"flv_url"`
	// HLS协议拉流地址
	HlsUrl string `json:"hls_url"`
}
