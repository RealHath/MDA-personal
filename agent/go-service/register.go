package main

import (
	"github.com/RealHath/MDA-personal/agent/go-service/common/myaction"
	"github.com/RealHath/MDA-personal/agent/go-service/common/myreco"
	"github.com/RealHath/MDA-personal/agent/go-service/pkg/resource"
	"github.com/RealHath/MDA-personal/agent/go-service/taskersink/aspectratio"
	"github.com/RealHath/MDA-personal/agent/go-service/taskersink/hdrcheck"
	"github.com/RealHath/MDA-personal/agent/go-service/taskersink/membership"
	"github.com/RealHath/MDA-personal/agent/go-service/taskersink/processcheck"
	"github.com/rs/zerolog/log"
)

func registerAll() {
	// Resource Sink
	resource.EnsureResourcePathSink()

	// Pre-Check Custom
	aspectratio.Register()
	hdrcheck.Register()
	processcheck.Register()
	membership.Register()

	// Custom Actions
	myaction.Register()

	// Custom Recognitions
	myreco.Register()

	log.Info().
		Msg("All custom components and sinks registered successfully")
}
