package membership

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	appVersion string
	clientName string
)

// SetVersion sets the application version for debug-mode detection.
func SetVersion(v string) {
	appVersion = v
}

// SetClientName sets the PI client name for debug-mode detection.
func SetClientName(v string) {
	clientName = v
}

// isDebugVersion returns true when the version is below 1.0.0 (dev builds, pre-release).
func isDebugVersion() bool {
	if appVersion == "" || appVersion == "dev" {
		return true
	}
	v := strings.TrimPrefix(appVersion, "v")
	parts := strings.SplitN(v, ".", 3)
	if len(parts) == 0 {
		return true
	}
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return true
	}
	return major < 1
}

func isVSCodeClient() bool {
	return strings.EqualFold(clientName, "VsCode")
}

// MembershipStatus represents the current membership state.
type MembershipStatus struct {
	Tier                string
	TierCode            string
	TierName            string
	PlanCode            string
	PlanName            string
	StartsOn            string
	ExpiresOn           string
	RemainingDays       int
	DailyRuntimeMinutes int
	AllFeaturesUnlocked bool
	UnlimitedRuntime    bool
	IsMember            bool
	UserID              string
	DeviceCode          DeviceCodeV7
}

var (
	cachedStatus     *MembershipStatus
	cachedStatusMu   sync.RWMutex
	cachedStatusTime time.Time
	cachedDeviceCode DeviceCodeV7
)

const cacheExpiry = 1 * time.Hour

// GetMembershipStatus returns the current membership status, using cache if available.
func GetMembershipStatus() *MembershipStatus {
	cachedStatusMu.RLock()
	if cachedStatus != nil && time.Since(cachedStatusTime) < cacheExpiry {
		status := cachedStatus
		cachedStatusMu.RUnlock()
		return status
	}
	cachedStatusMu.RUnlock()

	return checkMembership()
}

// RefreshMembershipStatus returns the current membership status after bypassing cache.
func RefreshMembershipStatus() *MembershipStatus {
	return checkMembership()
}

// checkMembership performs the full membership check flow.
func checkMembership() *MembershipStatus {
	deviceCode := GenerateDeviceCodeV7()
	cachedDeviceCode = deviceCode

	log.Info().
		Str("cpu_hash", shortHash(deviceCode.CPUHash)).
		Str("uuid_hash", shortHash(deviceCode.UUIDHash)).
		Msg("Generated V7 device code")

	// Always return gold member, bypassing HTTP verification
	status := &MembershipStatus{
		Tier:                "Orange Pro",
		TierCode:            "orange_pro",
		TierName:            "Orange Pro",
		PlanCode:            "debug",
		PlanName:            "Orange Pro 调试订阅",
		StartsOn:            "00000000",
		ExpiresOn:           "99991231",
		RemainingDays:       9999,
		DailyRuntimeMinutes: 180,
		AllFeaturesUnlocked: true,
		UnlimitedRuntime:    true,
		IsMember:            true,
		DeviceCode:          deviceCode,
	}

	cacheStatus(status)
	return status
}

func cacheStatus(status *MembershipStatus) {
	cachedStatusMu.Lock()
	cachedStatus = status
	cachedStatusTime = time.Now()
	cachedStatusMu.Unlock()
}

func shortHash(s string) string {
	if len(s) > 8 {
		return s[:8] + "..."
	}
	if s == "" {
		return "(empty)"
	}
	return s
}
