package ru

import (
	"strings"
	"testing"
	"time"

	"dev.azure.com/noon-homa/RentalsUnitedGoClient/_git/ru/utils"
	"github.com/stretchr/testify/assert"
)

func TestCanCreatePayload(t *testing.T) {
	from, _ := time.Parse(utils.DateLayout, "2020-01-01")
	to, _ := time.Parse(utils.DateLayout, "2020-02-01")
	res := getAvailabilityCommandPayload(from, to, BasePropertyID(123))
	assert.NotNil(t, res)
	assert.True(t, strings.Contains(res, "<PropertyID>"))
	assert.True(t, strings.Contains(res, "</PropertyID>"))
	assert.True(t, strings.Contains(res, "123"))
	assert.True(t, strings.Contains(res, "<DateFrom>"))
	assert.True(t, strings.Contains(res, "</DateFrom>"))
	assert.True(t, strings.Contains(res, "<DateTo>"))
	assert.True(t, strings.Contains(res, "</DateTo>"))
}
