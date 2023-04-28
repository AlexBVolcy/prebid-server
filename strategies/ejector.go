package strategies

import (
	"errors"
	"math"
	"time"

	"github.com/prebid/prebid-server/usersync"
)

type Ejector interface {
	Choose(uids map[string]usersync.UidWithExpiry)
}

type OldestEjector struct {
	nonPriorityKeys []string
}

func (o OldestEjector) Choose(uids map[string]usersync.UidWithExpiry) string {
	var oldestElem string = ""
	var oldestDate int64 = math.MaxInt64

	for _, key := range o.nonPriorityKeys {
		value := uids[key]
		timeUntilExpiration := time.Until(value.Expires)
		if timeUntilExpiration < time.Duration(oldestDate) {
			oldestElem = key
			oldestDate = int64(timeUntilExpiration)
		}
	}
	return oldestElem
}

type PriorityBidderEjector struct {
	PriorityGroups [][]string
	SyncerKey      string
	OldestEjector  OldestEjector
}

func (p PriorityBidderEjector) Choose(uids map[string]usersync.UidWithExpiry) (string, error) {
	p.OldestEjector.nonPriorityKeys = get(uids, p.PriorityGroups)

	// There are non priority keys present, let's eject one of those
	if len(p.OldestEjector.nonPriorityKeys) > 0 {
		return p.OldestEjector.Choose(uids), nil
	}

	// There are only priority keys left, check if the syncer is apart of the priority groups
	if isSyncerPriority(p.SyncerKey, p.PriorityGroups) {
		// Eject Oldest Priority Element
		var oldestElem string = ""
		var oldestDate int64 = math.MaxInt64
		for key, value := range uids {
			timeUntilExpiration := time.Until(value.Expires)
			if timeUntilExpiration < time.Duration(oldestDate) {
				oldestElem = key
				oldestDate = int64(timeUntilExpiration)
			}
		}
		return oldestElem, nil
	}
	return "", errors.New("syncer key " + p.SyncerKey + " is not in priority groups")
}

func isSyncerPriority(syncer string, priorityGroups [][]string) bool {
	for _, group := range priorityGroups {
		for _, bidder := range group {
			if syncer == bidder {
				return true
			}
		}
	}
	return false
}

func get(uids map[string]usersync.UidWithExpiry, priorityGroups [][]string) []string {
	nonPriorityKeys := []string{}
	for key := range uids {
		for _, group := range priorityGroups {
			isPriority := false
			for _, bidder := range group {
				if key == bidder {
					isPriority = true
					break
				}
			}
			if !isPriority {
				nonPriorityKeys = append(nonPriorityKeys, key)
			}
		}
	}
	return nonPriorityKeys
}
