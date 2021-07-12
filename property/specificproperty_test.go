package property

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseListSpecificPropertyPayload(t *testing.T) {

	payloadAsBytes := []byte(payloadMock)
	res, errParsing := parseListSepecificPropertyResponse(payloadAsBytes)
	assert.Nil(t, errParsing)
	assert.Equal(t, "One Bedroom Apartment", res.Name)
	assert.Len(t, res.Images, 2)
	assert.Equal(t, "https://a0.muscache.com/im/pictures/2762835/690c16be_original.jpg?aki_policy=x_large", res.Images[0].URL)

}

const payloadMock = `
<Property>
<Name>One Bedroom Apartment</Name>
<Images>
    <Image ImageTypeID="1" ImageReferenceID="1">https://a0.muscache.com/im/pictures/2762835/690c16be_original.jpg?aki_policy=x_large</Image>
    <Image ImageTypeID="3" ImageReferenceID="2">https://a0.muscache.com/im/pictures/2762748/fe106e44_original.jpg?aki_policy=x_large</Image>
</Images>
</Property>
`
