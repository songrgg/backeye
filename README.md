Backeye
========
[![Build Status](https://travis-ci.org/songrgg/backeye.png?branch=master)](https://travis-ci.org/songrgg/backeye)
[![Go Report Card](https://goreportcard.com/badge/github.com/songrgg/backeye?refresh=1)](https://goreportcard.com/report/github.com/songrgg/backeye)
[![Coverage Status](https://coveralls.io/repos/github/songrgg/backeye/badge.svg?branch=feature%2Fmultiple-watches)](https://coveralls.io/github/songrgg/backeye?branch=feature%2Fmultiple-watches)

Another API accuracy monitor tool.

Try to fetch & check HTTP API & RPC cronly.

### Sample API monitor definition
1. Add single watch to one task

```json
{
    "name": "Post API",
    "desc": "post API",
    "cron": "*/2 * * * *",
    "watches": [
        {
            "name": "post list",
            "desc": "post list",
            "interval": 0,
            "path": "https://api-prod.wallstreetcn.com/apiv1/content/articles",
            "method": "GET",
            "headers": {
                "User-Agent": "backeye"
            },
            "assertions": [
                {
                    "source": "header",
                    "operator": "equal",
                    "left": "status_code",
                    "right": "200"
                },
                {
                    "source": "header",
                    "operator": "not_empty",
                    "left": "X-Ivanka-Trace-Id",
                    "right": ""
                },
                {
                    "source": "body",
                    "operator": "equal",
                    "left": "code",
                    "right": "20000"
                }
            ]
        }
    ]
}
```

### Multiple API watches
1. Add multiple watches to one task
2. Share the variables exported by last watch

```json
{
    "name": "Post API",
    "desc": "post API",
    "cron": "*/2 * * * *",
    "watches": [
        {
            "name": "post list",
            "desc": "post list",
            "interval": 0,
            "path": "https://api-prod.wallstreetcn.com/apiv1/content/articles",
            "method": "GET",
            "headers": {
                "User-Agent": "backeye"
            },
            "assertions": [
                {
                    "source": "header",
                    "operator": "equal",
                    "left": "status_code",
                    "right": "200"
                },
                {
                    "source": "header",
                    "operator": "not_empty",
                    "left": "X-Ivanka-Trace-Id",
                    "right": ""
                },
                {
                    "source": "body",
                    "operator": "equal",
                    "left": "code",
                    "right": "20000"
                }
            ],
            "variables": [
                {
                    "name": "postID",
                    "value": "$RESPONSE.data.items[0].id"
                }
            ]
        },
        {
            "name": "post detail",
            "desc": "post detail",
            "interval": 0,
            "path": "https://api-prod.wallstreetcn.com/apiv1/content/articles/${postID}?extract=0",
            "method": "GET",
            "headers": {
                "User-Agent": "backeye"
            },
            "assertions": [
                {
                    "source": "header",
                    "operator": "equal",
                    "left": "status_code",
                    "right": "200"
                },
                {
                    "source": "body",
                    "operator": "equal",
                    "left": "code",
                    "right": "20000"
                }
            ]
        }
    ]
}
```
