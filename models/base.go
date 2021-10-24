package models

import (
	"time"

	"github.com/s-pos/go-utils/config"
)

const (
	enabled       = true
	defaultBool   = false
	defaultInt    = 0
	defaultString = ""

	// type store
	Offline = "offline"
	Online  = "online"
	// Source store
	Tokopedia = "tokopedia"
	Shopee    = "shopee"
)

var timezone = config.Timezone()

func convertTimezone(t time.Time) time.Time {
	return t.UTC().In(timezone)
}
