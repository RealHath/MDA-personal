package membership

import (
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
