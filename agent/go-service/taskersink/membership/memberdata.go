package membership

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

var appVersion string

// SetVersion sets the application version for debug-mode detection.
func SetVersion(v string) {
	appVersion = v
}

// MembershipStatus represents the calculated current membership state.
type MembershipStatus struct {
	MembershipType  string
	UserLevel       int
	RemainingValue  float64
	VirtualExpiry   string
	IsMember        bool
	UserID          string
	UnsupportedTier bool
	DeviceCode      DeviceCodeV6
}

var (
	cachedStatus   *MembershipStatus
	cachedStatusMu sync.RWMutex
	cachedDataTime time.Time
)

const cacheExpiry = 1 * time.Hour

// GetMembershipStatus returns the current membership status, using cache if available.
func GetMembershipStatus() *MembershipStatus {
	cachedStatusMu.RLock()
	if cachedStatus != nil && time.Since(cachedDataTime) < cacheExpiry {
		status := cachedStatus
		cachedStatusMu.RUnlock()
		return status
	}
	cachedStatusMu.RUnlock()

	return checkMembership()
}

// checkMembership performs the full membership check flow.
func checkMembership() *MembershipStatus {
	log.Info().Msg("Membership check bypassed, allowing all users")
	return &MembershipStatus{
		MembershipType: "金Doro会员",
		UserLevel:      3,
		IsMember:       true,
		VirtualExpiry:  "99991231",
	}
}
