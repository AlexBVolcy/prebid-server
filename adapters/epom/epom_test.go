package epom

import (
	"testing"

	"github.com/prebid/prebid-server/adapters/adapterstest"
	"github.com/prebid/prebid-server/config"
	"github.com/prebid/prebid-server/openrtb_ext"
)

func TestJsonSamples(t *testing.T) {
	bidder, buildErr := Builder(openrtb_ext.BidderEpom, config.Adapter{
		Endpoint: "https://an.epom.com/ortb"}, config.Server{ExternalUrl: "http://hosturl.com", GdprID: "1", Datacenter: "2"})

	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}

	adapterstest.RunJSONBidderTest(t, "epomtest", bidder)
}
