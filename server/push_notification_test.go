// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package server

import (
	"strings"
	"testing"
)

func TestPushNotification(t *testing.T) {
	msg := PushNotification{Type: "test"}
	json := msg.ToJson()
	result := PushNotificationFromJson(strings.NewReader(json))

	if msg.Type != result.Type {
		t.Fatal("Ids do not match")
	}
}
