Backeye
========
[![Build Status](https://travis-ci.org/songrgg/backeye.png?branch=master)](https://travis-ci.org/songrgg/backeye)
[![Go Report Card](https://goreportcard.com/badge/github.com/songrgg/backeye)](https://goreportcard.com/report/github.com/songrgg/backeye)

Another API accuracy monitor tool.

Try to fetch & check HTTP API & RPC cronly.

### Sample API monitor definition

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
