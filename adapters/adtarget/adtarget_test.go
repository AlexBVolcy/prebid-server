package adtarget

import (
	"testing"

	"github.com/prebid/prebid-server/adapters/adapterstest"
	"github.com/prebid/prebid-server/config"
	"github.com/prebid/prebid-server/openrtb_ext"
)

func TestJsonSamples(t *testing.T) {
	bidder, buildErr := Builder(openrtb_ext.BidderAdtarget, config.Adapter{
		Endpoint: "http://ghb.console.adtarget.com.tr/pbs/ortb"}, config.Server{ExternalUrl: "http://hosturl.com", GdprID: "1", Datacenter: "2"})

	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}

	adapterstest.RunJSONBidderTest(t, "adtargettest", bidder)
}
