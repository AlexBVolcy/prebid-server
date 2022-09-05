package smartrtb

import (
	"testing"

	"github.com/prebid/prebid-server/adapters/adapterstest"
	"github.com/prebid/prebid-server/config"
	"github.com/prebid/prebid-server/openrtb_ext"
	"github.com/stretchr/testify/assert"
)

func TestJsonSamples(t *testing.T) {
	bidder, buildErr := Builder(openrtb_ext.BidderSmartRTB, config.Adapter{
		Endpoint: "http://market-east.smrtb.com/json/publisher/rtb?pubid=test"}, config.Server{ExternalUrl: "http://hosturl.com", GdprID: "1", Datacenter: "2"})

	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}

	adapterstest.RunJSONBidderTest(t, "smartrtbtest", bidder)
}

func TestEndpointTemplateMalformed(t *testing.T) {
	_, buildErr := Builder(openrtb_ext.BidderSmartRTB, config.Adapter{
		Endpoint: "{{Malformed}}"}, config.Server{ExternalUrl: "http://hosturl.com", GdprID: "1", Datacenter: "2"})

	assert.Error(t, buildErr)
}
