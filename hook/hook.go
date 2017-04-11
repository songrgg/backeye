package hook

import (
	"github.com/songrgg/backeye/assertion"
)

// Hook defines the hooks called in schedules & assertions
type Hook interface {
	AssertionResult(*assertion.AssertionResult) error
}
