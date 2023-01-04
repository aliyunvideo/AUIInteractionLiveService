package models

import (
	"time"

	"ApsaraLive/pkg/alicloud/live"
)

type RoomInfo struct {
	// 直播Id
	ID string `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at"`
	// 直播开始时间
	StartedAt time.Time `json:"started_at"`
	// 直播结束时间
	StoppedAt time.Time `json:"stopped_at"`

	// 直播标题
	Title string `json:"title"`
	// 直播公告
	Notice string `json:"notice"`
	// 直播封面
	CoverUrl string `json:"cover_url"`

	// 主播Id
	AnchorId string `json:"anchor_id"`

	// 主播Nick
	AnchorNick string `json:"anchor_nick"`

	// 扩展字段
	Extends string `json:"extends" gorm:"type:text"`

	// 直播状态，0-准备中，1-已开始，2-已结束
	Status int `json:"status"`

	// 直播模式 0-普通直播, 1-连麦直播，2-PK直播
	Mode int `json:"mode"`
	// 群组Id
	ChatId string `json:"chat_id"`

	// 连麦Id
	MeetingId string `json:"meeting_id"`

	// 点播Id
	VodId string `json:"vod_id"`

	// 连麦成员信息（json序列化）
	MeetingInfo string `json:"meeting_info" gorm:"type:text"`

	// 直播间统计
	Metrics *Metrics `json:"metrics,omitempty" gorm:"-"`

	// 直播转录制，点播信息
	VodInfo *VodInfo `json:"vod_info,omitempty" gorm:"-"`

	// 用户状态
	UserStatus *UserStatus `json:"user_status,omitempty" gorm:"-"`

	// 推流相关地址信息，动态生成
	PushInfo *PushLiveInfo `json:"push_url_info,omitempty" gorm:"-"`
	// 拉流相关地址信息，动态生成
	PullInfo *live.PullLiveInfo `json:"pull_url_info,omitempty" gorm:"-"`

	// 连麦PK 等信息， 动态生成
	LinkInfo *LinkInfo `json:"link_info,omitempty" gorm:"-"`
}

type Metrics struct {
	// 当前直播间在线人数
	OnlineCount uint64 `json:"online_count" gorm:"-"`
	// 当前直播间点赞总量
	LikeCount uint64 `json:"like_count" gorm:"-"`
	// 当前直播间访问人次
	Pv uint64 `json:"pv" gorm:"-"`
	// 当前直播间独立访问用户数，
	Uv uint64 `json:"uv" gorm:"-"`
}

type UserStatus struct {
	// 用户设备是否静音
	Mute bool `json:"mute" gorm:"-"`
	// 静音来源
	MuteSource []string `json:"mute_source" gorm:"-"`
}

// LinkInfo 前缀字段（artc://）+固定字段（live.aliyun.com）+拉流标识位（play）+roomid（房间ID）+SdkIAppID（连麦应用ID）+UserID（连麦观众ID）+timestamp（有效时长时间戳）+token
// 播放前缀+播流域名+AppName（live） + StreamID（由连麦应用ID_房间ID_主播ID_camera组成）+auth_key
// LinkInfo 直播连麦推拉流
type LinkInfo struct {
	// 推流地址
	RtcPushUrl string `json:"rtc_push_url"`
	// 拉流地址
	RtcPullUrl string `json:"rtc_pull_url"`
	// 普通观众CDN拉流地址
	CdnPullInfo *live.PullLiveInfo `json:"cdn_pull_info"`
}

// PushLiveInfo 推流信息
type PushLiveInfo struct {
	// RTMP协议地址
	RtmpUrl string `json:"rtmp_url"`
	// RTS协议地址
	RtsUrl string `json:"rts_url"`
	// SRT协议地址
	SrtUrl string `json:"srt_url"`
}

type Status struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

const (
	// LiveStatusAll -1 所有状态
	LiveStatusAll = -1
	// LiveStatusPrepare 0 准备中
	LiveStatusPrepare = 0
	// LiveStatusOn 1 已开始
	LiveStatusOn = 1
	// LiveStatusOff 2 已结束
	LiveStatusOff = 2
)

const (
	// LiveModeNormal 0 常规模式
	LiveModeNormal = 0
	// LiveModeLink 1 连麦模式
	LiveModeLink = 1
	// LiveModePk 2 pk模式,暂未实现
	LiveModePk = 2
)
