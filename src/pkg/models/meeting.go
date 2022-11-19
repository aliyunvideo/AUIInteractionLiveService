package models

import "encoding/json"

type MeetingMember struct {
	// 会议(连麦)Id
	MeetingId string `json:"meeting_id"`
	// 用户Id
	UserId string `json:"user_id"`
	// 用户Nick
	UserNick string `json:"user_nick"`
	// 用户头像
	UserAvatar string `json:"user_avatar"`
	// 摄像头状态
	CameraOpened bool `json:"camera_opened"`
	// 麦克风状态
	MicOpened bool `json:"mic_opened"`
}

type MeetingInfo struct {
	Members []MeetingMember `json:"members"`
}

func (m *MeetingInfo) ParseJsonString(str string) error {
	if str == "" {
		return nil
	}
	return json.Unmarshal([]byte(str), m)
}

func (m *MeetingInfo) ToJsonString() (string, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
