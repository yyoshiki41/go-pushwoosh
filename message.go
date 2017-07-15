package pushwoosh

import (
	"context"
	"net/http"
)

type Message struct {
	Content            string   `json:"content"`
	SendDate           string   `json:"send_date,omitempty"`
	Devices            []string `json:"devices,omitempty"`
	Users              []string `json:"users,omitempty"`
	Campaign           string   `json:"campaign,omitempty"`
	IOSBadges          string   `json:"ios_badges,omitempty"`
	TimeZone           string   `json:"timezone,omitempty"`
	IgnoreUserTimezone bool     `json:"ignore_user_timezone,omitempty"`
	Platforms          []int64  `json:"platforms,omitempty"`
	Preset             string   `json:"preset,omitempty"`
	SendRate           int64    `json:"send_rate,omitempty"`
}

// CreateMessage creates a new push message.
// http://docs.pushwoosh.com/docs/createmessage
func (c *Client) CreateMessage(ctx context.Context, messages *[]Message) (*Result, error) {
	var result Result
	notifications := map[string]interface{}{
		"notifications": messages,
	}

	err := c.call(ctx, http.MethodPost, "createMessage", notifications, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
