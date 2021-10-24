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
)

var timezone = config.Timezone()

func convertTimezone(t time.Time) time.Time {
	return t.UTC().In(timezone)
}
