Backeye
========
Another API accuracy monitor tool.

Try to fetch & check HTTP API & RPC cronly.

### Sample API monitor definition

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

