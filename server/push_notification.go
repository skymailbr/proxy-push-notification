// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package server

import (
	"encoding/json"
	"io"
)

const (
	PushNotifyApple   = "apple"
	PushNotifyAndroid = "android"

	PushTypeMessage     = "message"
	PushTypeClear       = "clear"
	PushTypeUpdateBadge = "update_badge"
	PushTypeSession     = "session"
	PushTypeTest        = "test"

	PushMessageV2 = "v2"

	PushSoundNone = "none"
)

type PushNotification struct {
	DeviceID      string `json:"deviceIdentifier"`
	PushTokenHash string `json:"pushTokenHash"`
	Subject       string `json:"subject"`
	Signature     string `json:"signature"`
	Type          string `json:"type"`
	Priority      string `json:"priority"`
}

type PushNotificationDecryptedData struct {
	NId     string `json:"nid"`
	App     string `json:"app"`
	Subject string `json:"subject"`
	Type    string `json:"type"`
	Id      string `json:"id"`
}

func (me *PushNotification) ToJson() string {
	b, err := json.Marshal(me)
	if err != nil {
		return ""
	}
	return string(b)
}

func PushNotificationFromJson(data io.Reader) *PushNotification {
	decoder := json.NewDecoder(data)
	var me PushNotification
	err := decoder.Decode(&me)
	if err == nil {
		return &me
	} else {
		return nil
	}
}
