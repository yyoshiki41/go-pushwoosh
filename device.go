package pushwoosh

import (
	"context"
	"net/http"
)

const (
	// Possible device types
	DeviceTypeIOS          = 1
	DeviceTypeBB           = 2
	DeviceTypeAndroid      = 3
	DeviceTypeWindowsPhone = 5
	DeviceTypeOSX          = 7
	DeviceTypeWindows      = 8
	DeviceTypeAmazon       = 9
	DeviceTypeSafari       = 10
	DeviceTypeChrome       = 11
	DeviceTypeFirefox      = 12
)

// Device is a struct to register device.
type Device struct {
	Type       int64  `json:"device_type"`
	HardWareID string `json:"hwid"`
	Language   string `json:"language,omitempty"`
	Timezone   int64  `json:"timezone,omitempty"`
}

// RegisterDevice registers device for the application.
// http://docs.pushwoosh.com/docs/registerdevice
func (c *Client) RegisterDevice(ctx context.Context, device *Device) (*Result, error) {
	var result Result
	err := c.call(ctx, http.MethodPost, "registerDevice", device, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
