package pushwoosh

import (
	"context"
	"net/http"
)

// Message is a struct to create messages.
type Message struct {
	Content            string                 `json:"content"`
	SendDate           string                 `json:"send_date,omitempty"`
	Devices            []string               `json:"devices,omitempty"`
	Users              []string               `json:"users,omitempty"`
	Data               map[string]interface{} `json:"data,omitempty"` // key/value pair
	Campaign           string                 `json:"campaign,omitempty"`
	TimeZone           string                 `json:"timezone,omitempty"`
	IgnoreUserTimezone bool                   `json:"ignore_user_timezone,omitempty"`
	Platforms          []int64                `json:"platforms,omitempty"`
	Preset             string                 `json:"preset,omitempty"`
	SendRate           int64                  `json:"send_rate,omitempty"`
	// iOS related
	IOSBadges string `json:"ios_badges,omitempty"` // integer or use "+n" or "-n" to increment/decrement the badge value by n
	// Android related
	AndroidBadges string `json:"android_badges,omitempty"` // integer or use "+n" or "-n" to increment/decrement the badge value by n
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
