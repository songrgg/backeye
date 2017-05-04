package http

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	nethttp "net/http"
	"strings"
	"time"

	"github.com/robertkrimen/otto"

	"regexp"

	"github.com/songrgg/backeye/assertion"
	"github.com/songrgg/backeye/std"
	"github.com/songrgg/backeye/watch"
)

// Method indicates the HTTP methods
type Method int

// AssertionFunc indicates the assertion
// type AssertionFunc func(context.Context, *nethttp.Response) AssertionResult

const (
	// GET indicates the HTTP Get method
	GET Method = iota
	// POST indicates the HTTP Get method
	POST Method = iota
	// PUT indicates the HTTP Get method
	PUT Method = iota
	// DELETE indicates the HTTP Get method
	DELETE Method = iota
	// HEAD indicates the HTTP head method
	HEAD Method = iota
)

// Watch defines API sampling
type Watch struct {
	Name          string
	Desc          string
	Method        Method
	Path          string
	Variables     []watch.Variable
	PathVariables map[string]string
	Headers       map[string]string
	Timeout       time.Duration
	ExtractFunc   func(string) map[string]string
	Assertions    []assertion.AssertionFunc
}

// Run executes this watch
func (w *Watch) Run(ctx context.Context) (watch.Result, error) {
	var taskID int64
	if ctx.Value("task_id") != nil {
		taskID = ctx.Value("task_id").(int64)
	} else {
		taskID = 0
	}
	result := watch.Result{
		TaskID:        taskID,
		WatchName:     w.Name,
		ExecutionTime: time.Now(),
	}

	var vars map[string]string
	if ctx.Value("variables") != nil {
		vars = ctx.Value("variables").(map[string]string)
	} else {
		vars = make(map[string]string)
	}

	// resolve the path
	path := resolvePath(vars, w.Path)

	std.LogInfoc("http_watch", fmt.Sprintf("start to watch %s", path))

	switch w.Method {
	case GET:
		resp, err := nethttp.Get(path)
		if err != nil {
			return result, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return result, err
		}

		vm := initVM(body)
		newvars := renderVars(vm, w.Variables)
		for k, v := range vars {
			newvars[k] = v
		}

		ctx = context.WithValue(ctx, watch.ResponseBody, body)

		assertionResults := make([]assertion.Result, 0)
		for i, assertion := range w.Assertions {
			assertionResults = append(assertionResults, assertion(ctx, resp))
			if !assertionResults[i].Success {
				break
			}
		}
		result.Assertions = assertionResults
		result.ExtractValues = newvars
		return result, nil

	case POST:
	case PUT:
		return result, nil
	}

	return result, errors.New("unsupported HTTP method")
}

func initVM(body []byte) *otto.Otto {
	// Create vm for variable rendering
	vm := otto.New()
	vm.Set("$RESPONSE_RAW", string(body))
	vm.Run(`$RESPONSE_JSON = JSON.parse($RESPONSE_RAW)`)
	_, err := vm.Run(`JSON.parse($RESPONSE_RAW)`)
	if err != nil {
		std.LogErrorc("watch", err, "failed to parse response json")
	}
	return vm
}

func renderVars(vm *otto.Otto, vars []watch.Variable) map[string]string {
	m := make(map[string]string)
	for i := range vars {
		name, val := vars[i].Name, vars[i].Value
		v, err := vm.Run(val)
		if err != nil {
			std.LogErrorc("watch", err, "failed to render variable ["+name+"]")
			continue
		}
		std.LogDebugLn("success to render variable ["+name+"]: ", v)

		vm.Set(name, v)
		m[name], _ = v.ToString()
	}
	return m
}

func resolvePath(vars map[string]string, path string) string {
	pathVar, err := regexp.Compile(`\$\{(\w+)\}`)
	if err != nil {
		std.LogErrorc("http_watch", err, "failed to parse path")
		return path
	}

	newpath := pathVar.ReplaceAllFunc([]byte(path), func(p []byte) []byte {
		key := strings.Trim(string(p[2:len(p)-1]), " ")
		val, ok := vars[key]
		if ok {
			return []byte(val)
		}
		return p
	})
	return string(newpath)
}
