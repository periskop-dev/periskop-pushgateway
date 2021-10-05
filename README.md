# periskop-pushgateway

[![Build Status](https://api.cirrus-ci.com/github/soundcloud/periskop-pushgateway.svg)](https://cirrus-ci.com/github/soundcloud/periskop-pushgateway)

Pushgateway for [Periskop](https://github.com/soundcloud/periskop).

![2021-10-04_10-28](https://user-images.githubusercontent.com/280193/135818707-ad1d62b1-e65c-4878-a6e0-19d2f37e8022.png)


## API

### Push errors

`POST /errors`

```
curl -v -X POST -H "Content-Type:application/json" -d '
{
   "target_uuid":"5d9893c6-51d6-11ea-8aad-f894c260afe5",
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
                  "request_body":null
               }
            }
         ]
      }
   ]
}                                           
' localhost:7878/errors
```

### Get exported errors

`GET /-/errors/`

```json
{
  "aggregated_errors": [
    {
      "aggregation_key": "testing@test",
      "total_count": 2,
      "severity": "error",
      "latest_errors": [
        {
          "error": {
            "class": "testing",
            "message": "",
            "stacktrace": [
              "line 12:",
              "syntax error"
            ],
            "cause": null
          },
          "uuid": "41b2c83c-60b3-4e21-802c-c01966fd6ee8",
          "timestamp": "2021-09-03T13:58:15.932289427Z",
          "severity": "error",
          "http_context": {
            "request_method": "",
            "request_url": "http://example.com",
            "request_headers": {
              "Cache-Control": "no-cache"
            },
            "request_body": ""
          }
        },
        {
          "error": {
            "class": "testing",
            "message": "",
            "stacktrace": [
              "line 12:",
              "syntax error"
            ],
            "cause": null
          },
          "uuid": "9d3020b9-00c5-4be4-a75a-20751071efbb",
          "timestamp": "2021-09-03T13:58:30.187672797Z",
          "severity": "error",
          "http_context": {
            "request_method": "",
            "request_url": "http://example.com",
            "request_headers": {
              "Cache-Control": "no-cache"
            },
            "request_body": ""
          }
        }
      ],
      "created_at": "2021-09-03T13:58:15.932294302Z"
    }
  ],
  "target_uuid": "6134c27f-77da-4809-8a2e-a6fa32b8014b"
}
```
