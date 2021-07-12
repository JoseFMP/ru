package ru

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanParseRespSimple(t *testing.T) {

	respID1 := "123123123asdfasdf"
	resoPayload1 := mockResponsePayload(0, "Success", respID1)
	testResponseWellFormed(t, resoPayload1, 0)

	respID2 := "123123123asdfasdf"
	resoPayload2 := mockResponsePayload(-1, "The XML contains not implemented method", respID2)
	testResponseWellFormed(t, resoPayload2, -1)
}

func TestCanParseResponseID(t *testing.T) {
	respID1 := "123123123asdfasdf"
	resoPayload1 := mockResponsePayload(0, "Success", respID1)
	parsedResponse := testResponseWellFormed(t, resoPayload1, 0)
	if string(parsedResponse.ResponseID) != respID1 {
		t.Log("Response id not matching.")
		t.FailNow()
	}
}

func testResponseWellFormed(t *testing.T, payload string, expectedStatusID StatusID) *GenericCommandResponse {

	payloadAsByteString := []byte(payload)
	parsedPayload, err := parseCommandResponse(&payloadAsByteString)

	if err != nil {
		t.Logf("Failed to parse the payload: %v", err)
		t.FailNow()
	}

	if parsedPayload == nil {
		t.Log("The payload parsed to nil.")
		t.FailNow()
	}

	if &parsedPayload.Status == nil {
		t.Log("The status was not parsed correctly.")
		t.FailNow()
	}

	if parsedPayload.Status.ID != expectedStatusID {
		t.Logf("The status id should be %d but is %d.", expectedStatusID, parsedPayload.Status.ID)
		t.FailNow()
	}
	return parsedPayload
}

func parseBrokenResponse(t *testing.T) {
	response := []byte("tralara")

	_, err := parseCommandResponse(&response)

	assert.NotNil(t, err, "Parsing a wrong response did not return an error.")
}

func mockResponsePayload(statusID StatusID, statusMessage string, responseID string) string {
	return fmt.Sprintf(respTemplate, fmt.Sprintf("%d ", statusID), statusMessage, responseID)
}

func TestCreateSimpleReq(t *testing.T) {

	clientConfig := ClientConfig{
		UserName: "pepito",
		Password: "strange-password",
	}

	request := clientConfig.createRequestPullPayload(commands.ListStatuses, "")

	assert.NotEmpty(t, request, "Request string is empty")
	assert.Contains(t, request, "Pull_", "Request does not contain keyword Pull_")
	assert.Contains(t, request, "_RQ", "Request does not contain keyword _RQ")
}

func TestCanParseStatusMessage(t *testing.T) {
	respMock := mockResponsePayload(0, "Success", "123123")
	asByteSlice := ([]byte)(respMock)
	parsed, err := parseCommandResponse(&asByteSlice)
	assert.Nil(t, err)
	assert.NotNil(t, parsed)
	assert.Equal(t, parsed.Status.Status, "Success")

	respMock2 := mockResponsePayload(0, "Success3", "123123")
	asByteSlice2 := ([]byte)(respMock2)
	parsed2, err2 := parseCommandResponse(&asByteSlice2)
	assert.Nil(t, err2)
	assert.NotNil(t, parsed2)
	assert.Equal(t, parsed2.Status.Status, "Success3")
}

func TestCanParseErrors(t *testing.T) {

	errorPayload := fmt.Sprintf(errorTemplate, "-4", "Incorrect login or password")
	asByteArray := []byte(errorPayload)
	parsed, err := parseCommandResponse(&asByteArray)
	assert.Nil(t, err)
	assert.NotNil(t, parsed)
	assert.Equalf(t, StatusID(-4), parsed.Status.ID, "Status id missmatches in object: %v", formatJSON(parsed))
}

func formatJSON(object interface{}) string {
	asJSONProperty, _ := json.MarshalIndent(object, "", " ")
	return string(asJSONProperty)
}

func TestPayloadStripped(t *testing.T) {
	respMock := mockResponsePayload(0, "Success", "123123")
	asByteArray := []byte(respMock)
	parsed, errParsing := parseCommandResponse(&asByteArray)
	assert.Nil(t, errParsing)
	assert.NotNil(t, parsed)
	assert.True(t, len(parsed.Payload.Payload) > 0)
	payloadAsString := string(parsed.Payload.Payload)
	payloadAsString = strings.ReplaceAll(payloadAsString, " ", "")
	assert.True(t, len(payloadAsString) > 10)

	assert.False(t, strings.Contains(payloadAsString, "ResponseID"))
	assert.False(t, strings.Contains(payloadAsString, "Pull_ListStatuses_RS"))
	assert.True(t, string(parsed.Payload.XMLName.Local) == "Statuses")
}

const errorTemplate = `
<error ID="%s">%s</error>
`

const respTemplate = `
<Pull_ListStatuses_RS>
	<Status ID="%s">%s</Status>
	<ResponseID>%s</ResponseID>
	<Statuses>
			<StatusInfo ID="-6">This request was rate limited. Maximum number of requests allowed for this API method with specified parameters is {0} per {1} sliding window. You have reached this limit.</StatusInfo>
			<StatusInfo ID="-5">This request was rate limited. Maximum concurrent requests allowed for this API method with specified parameters is {0}. You have reached this limit.</StatusInfo>
			<StatusInfo ID="-4">Incorrect login or password</StatusInfo>
			<StatusInfo ID="-3">Invalid request. {0}</StatusInfo>
			<StatusInfo ID="-2">This request cannot be processed. {0}</StatusInfo>
			<StatusInfo ID="-1">The XML contains not implemented method</StatusInfo>
			<StatusInfo ID="0">Success</StatusInfo>
			<StatusInfo ID="1">Property is not available for a given dates</StatusInfo>
			<StatusInfo ID="2">Nothing available for a given dates</StatusInfo>
			<StatusInfo ID="3">Property has no price settings for a given dates</StatusInfo>
			<StatusInfo ID="4">Wrong destination id:{0}</StatusInfo>
			<StatusInfo ID="5">Wrong distance unit id:{0}</StatusInfo>
			<StatusInfo ID="6">Wrong composition room id:{0}</StatusInfo>
			<StatusInfo ID="7">Wrong amenity id:{0}</StatusInfo>
			<StatusInfo ID="8">Wrong arrival instructions</StatusInfo>
			<StatusInfo ID="9">Could not insert late arrival fee, From:{0} To:{1} Fee:{2}</StatusInfo>
			<StatusInfo ID="10">Could not insert early departure fee, From:{0} To:{1} Fee:{2}</StatusInfo>
			<StatusInfo ID="11">Wrong payment method id:{0}</StatusInfo>
			<StatusInfo ID="12">Wrong deposit type id:{0}</StatusInfo>
			<StatusInfo ID="13">Cancallation policies overlaps</StatusInfo>
			<StatusInfo ID="14">Owner does not exist</StatusInfo>
			<StatusInfo ID="15">Apartment name ({0}) already exist in database.</StatusInfo>
			<StatusInfo ID="16">You already defined apartment with PUID:{0}</StatusInfo>
			<StatusInfo ID="17">Unexpected error, contact IT or try again</StatusInfo>
			<StatusInfo ID="18">Property with given ID does not exist.</StatusInfo>
			<StatusInfo ID="19">Dates mishmash</StatusInfo>
			<StatusInfo ID="20">Past dates</StatusInfo>
			<StatusInfo ID="21">Weird block dates for property: {0} - {1} - {2}. Whole block is {3} - {4}</StatusInfo>
			<StatusInfo ID="22">We have confirmed reservation for those dates. Please cancel the reservation instead of marking the dates as available.</StatusInfo>
			<StatusInfo ID="23">Wrong ImageTypeID:{0}</StatusInfo>
			<StatusInfo ID="24">Your are not the owner of the apartment.</StatusInfo>
			<StatusInfo ID="25">The value of "Bigger" must be smaller than the value of "Smaller".</StatusInfo>
			<StatusInfo ID="26">Warning! Look at Notifs collection.</StatusInfo>
			<StatusInfo ID="27">DaysToArrivalFrom and DaysToArrivalTo requires positive values.</StatusInfo>
			<StatusInfo ID="28">Reservation does not exist.</StatusInfo>
			<StatusInfo ID="29">Requested stay, cost details do not match with property on reservation on hold.</StatusInfo>
			<StatusInfo ID="30">Element ignored because of other errors.</StatusInfo>
			<StatusInfo ID="31">Error occured. All changes rolled back.</StatusInfo>
			<StatusInfo ID="32">Bigger and Smaller requires positive values.</StatusInfo>
			<StatusInfo ID="33">Smaller is smaller than Bigger.</StatusInfo>
			<StatusInfo ID="34">RUPrice is not valid. Correct price is:{0}</StatusInfo>
			<StatusInfo ID="35">AlreadyPaid is bigger than ClientPrice.</StatusInfo>
			<StatusInfo ID="36">Wrong DetailedLocationID. City or district precision is required.</StatusInfo>
			<StatusInfo ID="37">Property name is too long (max 150).</StatusInfo>
			<StatusInfo ID="38">Property has missing data and cannot be offered.</StatusInfo>
			<StatusInfo ID="39">Location does not exist.</StatusInfo>
			<StatusInfo ID="40">You cannot define discounts before the prices. The property has missing prices in given dates.</StatusInfo>
			<StatusInfo ID="41">The reservation was created by the other user.</StatusInfo>
			<StatusInfo ID="42">The reservation is expired.</StatusInfo>
			<StatusInfo ID="43">You cannot confirm this reservation. It's broken.</StatusInfo>
			<StatusInfo ID="44">The apartments are not in the same city.</StatusInfo>
			<StatusInfo ID="45">Data validation error.</StatusInfo>
			<StatusInfo ID="46">The property is not active. PropertyID:{0}</StatusInfo>
			<StatusInfo ID="47">Property is not available for a given dates. PropertyID:{0}</StatusInfo>
			<StatusInfo ID="48">The reservation is not on Put On Hold status.</StatusInfo>
			<StatusInfo ID="49">CountryID does not exist.</StatusInfo>
			<StatusInfo ID="50">Guest name is required.</StatusInfo>
			<StatusInfo ID="51">Guest surname is required.</StatusInfo>
			<StatusInfo ID="52">Guest email is required.</StatusInfo>
			<StatusInfo ID="53">This method is deprecated. Use Push_PutConfirmedReservationMulti_RS</StatusInfo>
			<StatusInfo ID="54">This method is deprecated. Use Push_PutPropertiesOnHold_RQ</StatusInfo>
			<StatusInfo ID="55">Negative values in price elements is not allowed.</StatusInfo>
			<StatusInfo ID="56">Property does not exist.</StatusInfo>
			<StatusInfo ID="57">The request contains both types of composition definitions: composition and composition with amenities. Please use only one type.</StatusInfo>
			<StatusInfo ID="58">This amenity: {0} is not allowed in room type: {1}</StatusInfo>
			<StatusInfo ID="59">Positive value is required</StatusInfo>
			<StatusInfo ID="60">Duplicate value in LOSS element</StatusInfo>
			<StatusInfo ID="61">Duplicate value in EGPS element</StatusInfo>
			<StatusInfo ID="62">Missing Text or Image value.</StatusInfo>
			<StatusInfo ID="63">Wrong laguage id:{0}.</StatusInfo>
			<StatusInfo ID="64">DayOfWeek attribute must be between {0} and {1}.</StatusInfo>
			<StatusInfo ID="65">No permission to property {0}.</StatusInfo>
			<StatusInfo ID="66">Coordinates are invalid or missing.</StatusInfo>
			<StatusInfo ID="67">Duplicate value in LOSPS element</StatusInfo>
			<StatusInfo ID="68">NumberOfGuests in LOSP element has to be greather than 0</StatusInfo>
			<StatusInfo ID="69">Building does not exist</StatusInfo>
			<StatusInfo ID="70">Some properties not updated:{0}</StatusInfo>
			<StatusInfo ID="71">Wrong security deposit type id: {0}</StatusInfo>
			<StatusInfo ID="72">Discount value can't be lower than 0.</StatusInfo>
			<StatusInfo ID="73">At least one PropertyID element is required.</StatusInfo>
			<StatusInfo ID="74">DateFrom has to be earlier than DateTo.</StatusInfo>
			<StatusInfo ID="75">DateFrom has to be earlier or equal to DateTo.</StatusInfo>
			<StatusInfo ID="76">StandardGuests must be smaller than CanSleepMax.</StatusInfo>
			<StatusInfo ID="77">NOP: positive value required.</StatusInfo>
			<StatusInfo ID="78">Minimum stay is not valid (X nights).</StatusInfo>
			<StatusInfo ID="79">Stay period doesn't match with minimum stay</StatusInfo>
			<StatusInfo ID="80">Cannot activate archived property</StatusInfo>
			<StatusInfo ID="81">You don't have permission to modify this owner</StatusInfo>
			<StatusInfo ID="82">Apartment is Archived or no longer available or not Active</StatusInfo>
			<StatusInfo ID="83">Mixed owners in the request. Contact IT.</StatusInfo>
			<StatusInfo ID="84">Too many properties in your request (max 100).</StatusInfo>
			<StatusInfo ID="85">Invalid time value. Allowed values 00:00 - 23:59</StatusInfo>
			<StatusInfo ID="86">Operation has reached the maximum limit of time. The results are not complete.</StatusInfo>
			<StatusInfo ID="87">Wrong page URL Type</StatusInfo>
			<StatusInfo ID="88">Wrong date format for parameter {0}</StatusInfo>
			<StatusInfo ID="89">Stay period doesn't match with changeover</StatusInfo>
			<StatusInfo ID="90">Enqueued</StatusInfo>
			<StatusInfo ID="91">Not found</StatusInfo>
			<StatusInfo ID="92">Duplicate value in distances.</StatusInfo>
			<StatusInfo ID="93">Unauthorized</StatusInfo>
			<StatusInfo ID="94">Some of required fields were not filled.</StatusInfo>
			<StatusInfo ID="95">Email already exists.</StatusInfo>
			<StatusInfo ID="96">Password must be at least 8 characters long.</StatusInfo>
			<StatusInfo ID="97">Standard number of guests must be of positive value.</StatusInfo>
			<StatusInfo ID="98">Deposit amount can't exceed value of 214,748.3647</StatusInfo>
			<StatusInfo ID="99">Technical error - missing file</StatusInfo>
			<StatusInfo ID="100">Property description is required</StatusInfo>
			<StatusInfo ID="101">Pets not allowed</StatusInfo>
			<StatusInfo ID="102">Currency doesn't match with city currency</StatusInfo>
			<StatusInfo ID="103">Properties collection cannot be empty</StatusInfo>
			<StatusInfo ID="104">You need provide at least one value to modify stay.</StatusInfo>
			<StatusInfo ID="105">Some periods overlap. Periods must be separable.</StatusInfo>
			<StatusInfo ID="106">You can only modify stay in confirmed reservation.</StatusInfo>
			<StatusInfo ID="107">No reserved apartment found.</StatusInfo>
			<StatusInfo ID="108">Client Price cannot be negative</StatusInfo>
			<StatusInfo ID="109">Already Paid cannot be negative.</StatusInfo>
			<StatusInfo ID="110">Cannot use OwnerID created by other users.</StatusInfo>
			<StatusInfo ID="111">Only property owner can add reviews.</StatusInfo>
			<StatusInfo ID="112">Review rating value must be between 0-5</StatusInfo>
			<StatusInfo ID="113">Submitted date must be later than arrival date</StatusInfo>
			<StatusInfo ID="114">Cannot remove confirmed reservation. Some periods ignored.</StatusInfo>
			<StatusInfo ID="115">MinStays not satisfied, collection cannot be empty.</StatusInfo>
			<StatusInfo ID="116">LocationID {0} is not proper city location.</StatusInfo>
			<StatusInfo ID="117">Only one description allowed per language.</StatusInfo>
			<StatusInfo ID="118">Max number of guests must be of positive value.</StatusInfo>
			<StatusInfo ID="119">Property name is not defined.</StatusInfo>
			<StatusInfo ID="120">Check-in / check-out details are incorrect.</StatusInfo>
			<StatusInfo ID="121">Reservation not mapped in PMS. Contact IT support with Rentals United ID and your PMS Reservation ID.</StatusInfo>
			<StatusInfo ID="122">Failed to modify reservation in PMS. Try again or contact IT support for more information.</StatusInfo>
			<StatusInfo ID="123">Failed to cancel reservation in PMS. Try again or contact IT support for more information.</StatusInfo>
			<StatusInfo ID="124">Failed to insert reservation in PMS. Try again or contact IT support for more information.</StatusInfo>
			<StatusInfo ID="125">Wrong quantity of amenities. It should be between 0 - 32767.</StatusInfo>
			<StatusInfo ID="126">Invalid URL.</StatusInfo>
			<StatusInfo ID="127">Missing mandatory element: {0}.</StatusInfo>
			<StatusInfo ID="128">Cancellation policy text cannot be empty.</StatusInfo>
			<StatusInfo ID="129">Only reservations for apartments from same city are allowed</StatusInfo>
			<StatusInfo ID="130">Cannot change apartment from city other than initial reservation. Cancel this reservation and create new one.</StatusInfo>
			<StatusInfo ID="131">Bad request.</StatusInfo>
			<StatusInfo ID="132">This functionality is forbidden for you.</StatusInfo>
			<StatusInfo ID="133">Too many images. Images limit is 100.</StatusInfo>
			<StatusInfo ID="134">Invalid currency.</StatusInfo>
			<StatusInfo ID="135">Request rejected by partner.</StatusInfo>
			<StatusInfo ID="136">Customer info is required.</StatusInfo>
			<StatusInfo ID="137">PMSID is not valid.</StatusInfo>
			<StatusInfo ID="138">Provide a not empty PUID.</StatusInfo>
			<StatusInfo ID="139">Cannot create new property as archived.</StatusInfo>
			<StatusInfo ID="140">PropertyID : {0} cannot be archived / activated. The whole request was cancelled. Use Push_PutProperty_RQ instead.</StatusInfo>
			<StatusInfo ID="141">PUID already exists for another property.</StatusInfo>
			<StatusInfo ID="142">Property already has PMSID assigned. You cannot omit PMSID while update.</StatusInfo>
			<StatusInfo ID="143">Cannot archive this property. {0}</StatusInfo>
			<StatusInfo ID="144">Invalid Additional fees collection. {0}</StatusInfo>
			<StatusInfo ID="145">There are not enough units in this property.</StatusInfo>
			<StatusInfo ID="146">Multi unit functionality is disabled.</StatusInfo>
			<StatusInfo ID="147">Changeover is invalid. Use number 1, 2, 3 or 4.</StatusInfo>
			<StatusInfo ID="148">Number of units do not match with reservation on hold.</StatusInfo>
			<StatusInfo ID="149">Units must be a positive number</StatusInfo>
			<StatusInfo ID="150">Invalid invoice ID</StatusInfo>
			<StatusInfo ID="151">CancelUrl is missing</StatusInfo>
			<StatusInfo ID="152">ReturnUrl is missing</StatusInfo>
			<StatusInfo ID="153">PayPal transaction failed</StatusInfo>
			<StatusInfo ID="154">Single unit apartment cannot be converted to multi unit</StatusInfo>
			<StatusInfo ID="155">ResApaID has to be provided for multi unit properties</StatusInfo>
			<StatusInfo ID="156">ResApaID is not valid</StatusInfo>
			<StatusInfo ID="157">Card registration failed</StatusInfo>
			<StatusInfo ID="158">Request rejected. Check PMS synchronization settings</StatusInfo>
			<StatusInfo ID="159">Failed to insert reservation for property {0}({1}) to PMS. {2}</StatusInfo>
			<StatusInfo ID="160">The location and city already added to database</StatusInfo>
			<StatusInfo ID="161">You are not permitted to list this user's reservations</StatusInfo>
			<StatusInfo ID="162">Pass valid email</StatusInfo>
			<StatusInfo ID="163">Future dates</StatusInfo>
			<StatusInfo ID="164">Invalid date format</StatusInfo>
			<StatusInfo ID="165">You cannot archive a confirmed reservation. If this is a cancellation please cancel it first.</StatusInfo>
			<StatusInfo ID="166">Reservation is archived. Unarchive first before performing this operation.</StatusInfo>
			<StatusInfo ID="167">Reservation is for properties you are not the owner of</StatusInfo>
			<StatusInfo ID="168">Reservation is not canceled.</StatusInfo>
			<StatusInfo ID="169">Cannot process a request because apartment is connected to RMS</StatusInfo>
			<StatusInfo ID="170">Invalid property type id</StatusInfo>
			<StatusInfo ID="171">Invalid xml format in AdditionalData node</StatusInfo>
			<StatusInfo ID="172">Preparation time before arrival is blocking requested stay</StatusInfo>
			<StatusInfo ID="173">Discount ignored or trimmed. Maximum length of discounted stay is {0} nights.</StatusInfo>
			<StatusInfo ID="174">Invalid field LicenceInfo/FrenchLicenceInfo/CityTaxCategory. Allowed values: 11-19</StatusInfo>
			<StatusInfo ID="175">Invalid field LicenceInfo/FrenchLicenceInfo/TypeOfResidence. Allowed values: 1-3</StatusInfo>
			<StatusInfo ID="176">Request confirmation in external system failed</StatusInfo>
			<StatusInfo ID="177">Property is not connected to external system</StatusInfo>
			<StatusInfo ID="178">This reservation was made in external system and cannot be cancelled in Rentals United. Please cancel it directly in the sales channel</StatusInfo>
			<StatusInfo ID="179">Additional fee value must be greater or equal zero</StatusInfo>
			<StatusInfo ID="180">LOS pricing and full stay pricing cannot be defined in the same season</StatusInfo>
			<StatusInfo ID="181">Number of units cannot be decreased</StatusInfo>
	</Statuses>
</Pull_ListStatuses_RS>
`
