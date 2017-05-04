package http

import (
	"context"
	"net/http"
	"regexp"
	"testing"
	"time"

	"strings"

	"github.com/songrgg/backeye/assertion"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	watch := Watch{
		Method:  GET,
		Path:    "https://api-prod.wallstreetcn.com/apiv1/content/articles",
		Timeout: time.Second,
		Assertions: []assertion.AssertionFunc{
			func(ctx context.Context, resp *http.Response) assertion.Result {
				return assertion.Result{
					Passed: resp.StatusCode == 200,
				}
			},
		},
	}

	result, err := watch.Run(context.Background())
	assert.Equal(t, err, nil)
	assert.Equal(t, len(result.Assertions), 1)
	assert.Equal(t, result.Assertions[0].Passed, true)
}

func TestMultipleWatch(t *testing.T) {
	watch := Watch{
		Method:  GET,
		Path:    "https://api-prod.wallstreetcn.com/apiv1/content/articles",
		Timeout: time.Second,
		Assertions: []assertion.AssertionFunc{
			func(ctx context.Context, resp *http.Response) assertion.Result {
				return assertion.Result{
					Passed: resp.StatusCode == 200,
				}
			},
			func(ctx context.Context, resp *http.Response) assertion.Result {
				return assertion.Result{
					Passed: resp.Status == "200 OK",
				}
			},
		},
	}

	result, err := watch.Run(context.Background())
	assert.Equal(t, err, nil)
	assert.Equal(t, len(result.Assertions), 2)
	assert.Equal(t, result.Assertions[0].Passed, true)
	assert.Equal(t, result.Assertions[1].Passed, true)
}

func TestPathRender(t *testing.T) {
	m := map[string]string{
		"postID": "2",
	}
	path := "https://api-prod.wallstreetcn.com/apiv1/content/articles/${postID}"
	pathVar := regexp.MustCompile(`\$\{(\w+)\}`)
	newpath := pathVar.ReplaceAllFunc([]byte(path), func(p []byte) []byte {
		key := strings.Trim(string(p[2:len(p)-1]), " ")
		return []byte(m[key])
	})
	assert.Equal(t, "https://api-prod.wallstreetcn.com/apiv1/content/articles/2", string(newpath))

	// vm := otto.New()
	// vm.Set("task", `{"name":"test"}`)
	// val, _ := vm.Run(`JSON.parse(task).name`)
	// fmt.Println(val)
}
