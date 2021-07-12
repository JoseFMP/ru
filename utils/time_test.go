package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestYearDayAsTime(t *testing.T) {

	yearDayMock1 := YearDay{
		Year: 2021,
		Day:  40,
	}

	asTime := yearDayMock1.asTime()

	assert.Equal(t, 2021, asTime.Year())
	assert.Equal(t, time.February, asTime.Month())
	assert.Equal(t, 10, asTime.Day())

}
