package membership

import (
	"fmt"

	maa "github.com/MaaXYZ/maa-framework-go/v4"
	"github.com/rs/zerolog/log"
)

type MembershipCheckAction struct{}

type RuntimeQuotaCheckAction struct{}

var (
	_ maa.CustomActionRunner = &MembershipCheckAction{}
	_ maa.CustomActionRunner = &RuntimeQuotaCheckAction{}
)

func (a *MembershipCheckAction) Run(ctx *maa.Context, arg *maa.CustomActionArg) bool {
	log.Info().Msg("MembershipCheck: check bypassed, allowing all users")
	return true
}

func (a *RuntimeQuotaCheckAction) Run(ctx *maa.Context, arg *maa.CustomActionArg) bool {
	log.Info().Msg("RuntimeQuotaCheck: quota check bypassed, allowing all users")
	return true
}

// formatQuotaDeniedMessage is required by runtime_tracker.go (dead code since RuntimeTracker is not registered).
func formatQuotaDeniedMessage(snapshot QuotaSnapshot) string {
	return fmt.Sprintf("Quota exhausted. Tier: %s, Limit: %d min, Sponsor: %s",
		snapshot.TierName, FormatMinutes(snapshot.LimitSeconds), snapshot.SponsorURL)
}
