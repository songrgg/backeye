Backeye
========
[![Build Status](https://travis-ci.org/songrgg/backeye.svg?branch=master)](https://travis-ci.org/songrgg/backeye)
[![Go Report Card](https://goreportcard.com/badge/github.com/songrgg/backeye?refresh=1)](https://goreportcard.com/report/github.com/songrgg/backeye)
[![Coverage Status](https://coveralls.io/repos/github/songrgg/backeye/badge.svg?branch=master)](https://coveralls.io/github/songrgg/backeye?branch=master)

Another API accuracy monitor tool.

Try to fetch & check HTTP API & RPC cronly.


### Quick start
Take MySQL ready, and make sure it's accessible, update the `conf/backeye.yaml`.

Start the server
```shell
go run main.go
```

### API document
Access [http://localhost:9876/swagger-ui/index.html](http://localhost:9876/swagger-ui/index.html).

![API swagger](https://raw.githubusercontent.com/songrgg/backeye/feature/swagger-doc/public/swagger-ui/images/backeye-swagger.png)

### API testing cases

#### Sample API monitor definition
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

#### Multiple API watches
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

### Run With Docker

Get the configuration ready, for example in the test environment  
**backeye.test.yaml**
```yaml
bind: ":9876"

log:
  level: 5

mysql:
    username: "<MYSQL_USERNAME>"
    password: "<MYSQL_PASSWORD>"
    host: "<MYSQL_HOST>"
    port: 3306
    db_name: "backeye"
    max_idle: 50
    max_conn: 100
```

Start the backeye with docker, set `CONFIGOR_ENV` to `test` and mount the configuration file.
```shell
docker run -d -e CONFIGOR_ENV=test -v conf:/usr/src/app/conf songrgg/backeye
```
