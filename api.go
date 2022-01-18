package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/periskop-dev/periskop-go"
	"github.com/periskop-dev/periskop-pushgateway/adapters"
)

func NewErrorsGatewayHandler(c *periskop.ErrorCollector) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			c.ReportWithHTTPRequest(err, req)
			http.Error(w, err.Error(), 500)
		}
		res := adapters.Payload{}
		err = json.Unmarshal(payload, &res)
		if err != nil {
			c.ReportWithHTTPRequest(err, req)
			http.Error(w, err.Error(), 500)
		}
		for _, errorAggregate := range res.AggregatedErrors {
			for _, err := range errorAggregate.LatestErrors {
				errWithContext := adapters.ToPeriskopErrorWithContext(err)
				c.ReportErrorWithContext(errWithContext, periskop.Severity(errorAggregate.Severity), errorAggregate.AggregationKey)
			}
		}
	})
}
