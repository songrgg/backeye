package hook

import (
	"github.com/songrgg/backeye/assertion"
	"github.com/songrgg/backeye/std"
)

type LoggerHook struct {
}

func (lh *LoggerHook) AssertionResult(ar *assertion.AssertionResult) error {
	if !ar.Success {
		std.LogErrorc("assertion_logger_hook", ar.Error, "assertion result catched")
	} else {
		std.LogInfoc("assertion_logger_hook", "assertion result catched")
	}
	return nil
}
