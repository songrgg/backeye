package watcher

import (
	"context"
	"errors"

	"encoding/json"

	"github.com/songrgg/backeye/pkg/common"
	"github.com/songrgg/backeye/pkg/variable"
	"github.com/songrgg/backeye/pkg/watcher/http"
	httpConf "github.com/songrgg/backeye/pkg/watcher/http/conf"
)

type Watcher struct {
	points []Point
}

func (w *Watcher) AddPoint(p Point) {
	w.points = append(w.points, p)
}

func (w *Watcher) Do(ctx context.Context) ([]AssertionResult, error) {
	var assertionResult []AssertionResult
	var variables interface{}
	for _, p := range w.points {
		newCtx := ctx
		if variables != nil {
			newCtx = context.WithValue(newCtx, common.VariablesKey, variables)
		}
		newCtx, err := p.Do(newCtx)

		if ar := newCtx.Value(common.AssertionResultsKey); ar != nil {
			assertionResult = append(assertionResult, ar.([]AssertionResult)...)
		}

		if err != nil {
			return assertionResult, err
		}

		variables = newCtx.Value(common.VariablesKey)
	}
	return assertionResult, nil
}

func NewWatcher(config string) (*Watcher, error) {
	if config == "" {
		return nil, errors.New("empty watcher config")
	}

	watcher := &Watcher{}
	conf := Config{}
	err := json.Unmarshal([]byte(config), &conf)
	if err != nil {
		return nil, err
	}

	for _, p := range conf.Points {
		if p.Type == "http" {
			pointConf := httpConf.Config{}
			err := json.Unmarshal([]byte(p.Conf), &pointConf)
			if err != nil {
				return nil, err
			}

			var vars []variable.Variable
			for _, v := range pointConf.Variables {
				vars = append(vars, &v)
			}

			var assertions []Assertion
			for _, a := range pointConf.Assertions {
				assert, err := http.NewAssertionWithConf(&a)
				if err != nil {
					return nil, err
				}
				assertions = append(assertions, assert)
			}

			watcher.AddPoint(Point{
				Task:       &pointConf.Task,
				Assertions: assertions,
				Variables:  vars,
			})
		} else {
			return nil, errors.New("unsupported watcher point type")
		}
	}
	return watcher, nil
}
