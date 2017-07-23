# go-pushwoosh

The Pushwoosh Remote APIs Client for Go

[![godoc](https://godoc.org/github.com/yyoshiki41/go-pushwoosh?status.svg)](https://godoc.org/github.com/yyoshiki41/go-pushwoosh)
[![build](https://travis-ci.org/yyoshiki41/go-pushwoosh.svg?branch=master)](https://travis-ci.org/yyoshiki41/go-pushwoosh)
[![go report](https://goreportcard.com/badge/github.com/yyoshiki41/go-pushwoosh)](https://goreportcard.com/report/github.com/yyoshiki41/go-pushwoosh)

- [Official Documentation](http://docs.pushwoosh.com/docs/createmessage)

## Requirements

- Go v1.7+

## Usage

```go
package main

import (
	"context"
	"log"

	"github.com/yyoshiki41/go-pushwoosh"
)

func main() {
	// Set config
	conf := pushwoosh.Config{
		Endpoint:        "https://cp.pushwoosh.com/json",
		ApplicationCode: "aaaaa-00000",
		AccessToken:     "xxx",
	}

	// Create a new client
	client, err := pushwoosh.NewClient(&conf)
	if err != nil {
		log.Fatal(err)
	}

	// Create a Message struct
	msg := pushwoosh.Message{
		Content:   "content is here",
		SendDate:  "now",
		Devices:   []string{"Hardware ID is here"},
		IOSBadges: "+1",
	}
	msgs := []pushwoosh.Message{msg}

	// Call pushwoosh API
	_, err = client.CreateMessage(context.Background(), &msgs)
	if err != nil {
		log.Fatal(err)
	}
}
```

## REMOTE APIs

✅ /createMessage
- [ ] /deleteMessage
- [ ] /getMessageDetails

✅ /registerDevice
- [ ] /unregisterDevice
- [ ] /addTag
- [ ] /deleteTag
- [ ] /listTags
- [ ] /setTags
- [ ] /getTags
- [ ] /setBadge
- [ ] /applicationOpen
- [ ] /pushStat
- [ ] /messageDeliveryEvent
- [ ] /setPurchase
- [ ] /getNearestZone
- [ ] /createTargetedMessage
- [ ] /addGeoZone
- [ ] /addGeoZoneCluster
- [ ] /deleteGeoZone
- [ ] /deleteGeoZoneCluster
- [ ] /listGeoZoneClusters
- [ ] /listGeoZones
- [ ] /updateGeoZone
- [ ] /listTestDevices
- [ ] /getApplicationSubscribersStats
- [ ] /getTagStats
- [ ] /getAppStats
- [ ] /getCampaignStats
- [ ] /getMsgStats
- [ ] /getMsgPlatformsStats
- [ ] /getResults
- [ ] /configureApplication for iOS
- [ ] /configureApplication for Android
- [ ] /configureApplication for Windows Phone
- [ ] /configureApplication for Mac OS X
- [ ] /configureApplication for Blackberry
- [ ] /configureApplication for Windows 8
- [ ] /configureApplication for Amazon
- [ ] /configureApplication for Safari and Chrome
- [ ] /createApplication
- [ ] /updateApplication
- [ ] /deleteApplication
- [ ] /createCampaign
- [ ] /deleteCampaign
- [ ] /createTestDevice
- [ ] /getApplication
- [ ] /getApplications
- [ ] /getApplicationFile
- [ ] /postEvent
- [ ] /registerUser
- [ ] /getPushHistory
- [ ] /getPreset
- [ ] /createFilter
- [ ] /listFilters
- [ ] /deleteFilter
