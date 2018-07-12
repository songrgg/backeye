package variable

import "github.com/songrgg/backeye/pkg/response"

// Variable extracts the specified value from point response and register them in the variables section.
type Variable interface {
	Extract(ctx response.Response) (string, interface{}, error)
}
