package membership

import (
	maa "github.com/MaaXYZ/maa-framework-go/v4"
	"github.com/rs/zerolog/log"
)

// MembershipCheckAction is a placeholder action that no longer enforces membership checks.
type MembershipCheckAction struct{}

var _ maa.CustomActionRunner = &MembershipCheckAction{}

func (a *MembershipCheckAction) Run(ctx *maa.Context, arg *maa.CustomActionArg) bool {
	log.Info().Msg("MembershipCheck: check bypassed, allowing all users")
	return true
}
