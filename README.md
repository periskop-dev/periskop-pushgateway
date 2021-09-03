# periskop-pushgateway

## API

### Push errors

`POST /errors/{target}`


# Get list of errors

`GET /errors/`

```json
[
  {
    "target_uuid": "5d9893c6-51d6-11ea-8aad-f894c260afe5",
    "aggregated_errors":[
		  {
			"aggregation_key":"test",
			"total_count":1,
			"severity":"error",
			"created_at":"2020-02-17T22:42:45Z",
			"latest_errors":[
			  {
				"error":{
				  "class":"testing",
				  "message":"",
				  "stacktrace":[
					"line 12:",
					"syntax error"
				  ],
				  "cause":null
				},
				"uuid":"5d9893c6-51d6-11ea-8aad-f894c260afe5",
				"timestamp":"2020-02-17T22:42:45Z",
				"severity":"error",
				"http_context":{
				  "request_method":"GET",
				  "request_url":"http://example.com",
				  "request_headers":{
					"Cache-Control":"no-cache"
				  },
				  "request_body": null
				}
			  }
			]
		  }
    ]
  },
    {
    "target_uuid": "5d9893c6-51d6-11ea-8aad-f894c260afe5",
    "aggregated_errors":[
		  {
			"aggregation_key":"test",
			"total_count":1,
			"severity":"error",
			"created_at":"2020-02-17T22:42:45Z",
			"latest_errors":[
			  {
				"error":{
				  "class":"testing",
				  "message":"",
				  "stacktrace":[
					"line 12:",
					"syntax error"
				  ],
				  "cause":null
				},
				"uuid":"5d9893c6-51d6-11ea-8aad-f894c260afe5",
				"timestamp":"2020-02-17T22:42:45Z",
				"severity":"error",
				"http_context":{
				  "request_method":"GET",
				  "request_url":"http://example.com",
				  "request_headers":{
					"Cache-Control":"no-cache"
				  },
				  "request_body": null
				}
			  }
			]
		  }
    ]
  }
]
```
