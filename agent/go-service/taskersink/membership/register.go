package membership

import maa "github.com/MaaXYZ/maa-framework-go/v4"

// Register registers membership quota checks and runtime tracking.
func Register() {
	maa.AgentServerRegisterCustomAction("RuntimeQuotaCheck", &RuntimeQuotaCheckAction{})
<<<<<<< HEAD
=======
	maa.AgentServerRegisterCustomAction("QuotaDisplayAction", &QuotaDisplayAction{})
	maa.AgentServerAddTaskerSink(runtimeTracker)
	maa.AgentServerAddContextSink(runtimeTracker)
>>>>>>> upstream/main
}
