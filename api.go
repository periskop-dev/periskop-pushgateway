package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/soundcloud/periskop-go"
	"github.com/soundcloud/periskop-pushgateway/adapters"
)

func NewErrorsGatewayHandler(c *periskop.ErrorCollector) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		res := adapters.Payload{}
		json.Unmarshal(payload, &res)
		for _, errorAggregate := range res.AggregatedErrors {
			for _, err := range errorAggregate.LatestErrors {
				errWithContext := adapters.ToPeriskopErrorWithContext(err)
				c.ReportErrorWithContext(errWithContext, periskop.Severity(errorAggregate.Severity), errorAggregate.AggregationKey)
			}
		}
	})
}
