package andbeyondmedia

import (
	"testing"

	"github.com/prebid/prebid-server/adapters/adapterstest"
	"github.com/prebid/prebid-server/config"
	"github.com/prebid/prebid-server/openrtb_ext"
)

func TestJsonSamples(t *testing.T) {
	bidder, buildErr := Builder(openrtb_ext.BidderAndBeyondMedia, config.Adapter{
		Endpoint: "http://backend.andbeyond.media/pserver"}, config.Server{ExternalUrl: "http://hosturl.com", GvlID: 1, Datacenter: "2"})

	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}

	adapterstest.RunJSONBidderTest(t, "andbeyondmediatest", bidder)
}
