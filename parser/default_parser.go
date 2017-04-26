package parser

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	nethttp "net/http"
	"strconv"
	"time"

	"github.com/songrgg/backeye/assertion"
	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/std"
	"github.com/songrgg/backeye/task"
	"github.com/songrgg/backeye/watch"
	"github.com/songrgg/backeye/watch/http"
	modelmapper "gopkg.in/jeevatkm/go-model.v0"
)

type DefaultParser struct {
	Task *Task
}

// Task defines the JSON
type Task struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Desc    string  `json:"desc"`
	Cron    string  `json:"cron"`
	Watches []Watch `json:"watches"`
}

type Watch struct {
	Name          string            `json:"name"`
	Desc          string            `json:"desc"`
	Timeout       int32             `json:"timeout"`
	Interval      int               `json:"interval"`
	Path          string            `json:"path"`
	Method        string            `json:"method"`
	PathVariables []PathVar         `json:"path_variables"`
	Headers       map[string]string `json:"headers"`
	Assertions    []Assertion       `json:"assertions"`
}

type PathVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Assertion struct {
	Source   string `json:"source"`
	Operator string `json:"operator"`
	Left     string `json:"left"`
	Right    string `json:"right"`
}

func (p *DefaultParser) load(data []byte) error {
	task := Task{}
	err := json.Unmarshal(data, &task)
	if err != nil {
		log.Println("json Unmarshal error: ", err)
		return err
	}
	p.Task = &task
	return nil
}

// TranslateModel translates model to task
func (p *DefaultParser) ParseTask(t *model.Task) (*task.Task, error) {
	task := task.Task{}
	modelmapper.Copy(&task, t)

	watches := make([]watch.Watch, 0)
	for i := range t.Watches {
		watch, err := p.ParseWatch(&t.Watches[i])
		if err != nil {
			std.LogErrorc("default_parser", err, "failed to parse watch")
		}
		watches = append(watches, watch)
	}
	task.Watches = watches
	return &task, nil
}

func (p *DefaultParser) ParseWatch(w *model.Watch) (watch.Watch, error) {

	watch := &http.Watch{
		Name: w.Name,
		Desc: w.Desc,
		Path: w.Path,
	}

	// parse method
	switch w.Method {
	case "GET":
		watch.Method = http.GET
	case "POST":
		watch.Method = http.POST
	case "PUT":
		watch.Method = http.PUT
	case "HEAD":
		watch.Method = http.HEAD
	}

	// parse path variables
	// pathVars := make(map[string]string)
	// for _, pathvar := range w.PathVariables {
	// 	pathVars[pathvar.Name] = pathvar.Value
	// }
	// watch.PathVariables = pathVars

	if w.Timeout > 0 {
		watch.Timeout = time.Duration(w.Timeout) * time.Millisecond
	}

	assertions := make([]assertion.AssertionFunc, 0)
	for i := range w.Assertions {
		assertions = append(assertions, p.parseAssertion(&w.Assertions[i]))
	}
	watch.Assertions = assertions

	return watch, nil
}

func (p *DefaultParser) parseAssertion(t *model.Assertion) assertion.AssertionFunc {
	return func(ctx context.Context, resp *nethttp.Response) assertion.AssertionResult {
		body := ctx.Value(watch.ResponseBody)

		v := make(map[string]interface{})
		err := json.Unmarshal(body.([]byte), &v)
		if err != nil {
			return assertion.AssertionResult{
				Success: false,
				Error:   err,
			}
		}

		left := ""
		if t.Source == "header" {
			if t.Left == "status_code" {
				left = strconv.Itoa(resp.StatusCode)
			} else {
				left = resp.Header.Get(t.Left)
			}
		} else if t.Source == "body" {
			switch v[t.Left].(type) {
			case float64:
				left = strconv.FormatFloat(v[t.Left].(float64), 'g', -1, 64)
			default:
				left = v[t.Left].(string)
			}
		} else {
			return assertion.AssertionResult{
				Success: false,
				Error:   errors.New("invalid source"),
			}
		}

		right := t.Right

		success := false

		if t.Operator == "equal" {
			success = left == right
			err = errors.New("not equal")
		} else if t.Operator == "not_empty" {
			success = left != ""
			err = errors.New("not empty")
		}
		return assertion.AssertionResult{
			Success: success,
			Error:   err,
		}
	}
}

// Translate translates JSON to task
func (p *DefaultParser) Translate(data interface{}) (*task.Task, error) {
	return p.ParseTask(data.(*model.Task))
	// if err := p.load(data); err != nil {
	// 	return nil, err
	// }

	// task := task.Task{}
	// if p.Task != nil {
	// 	task.Name = p.Task.Name
	// 	task.Desc = p.Task.Desc
	// 	task.CronSpec = p.Task.Cron
	// }

	// task.Watches = make([]watch.Watch, 0)
	// for _, watch := range p.Task.Watches {
	// 	task.Watches = append(task.Watches, parseWatch(&watch))
	// }
	// return &task, nil
}