package adapters

import (
	"time"

	"github.com/periskop-dev/periskop-go"
)

type Payload struct {
	AggregatedErrors []ErrorAggregate `json:"aggregated_errors"`
	TargetUUID       string           `json:"target_uuid"` // ignored
}

type ErrorAggregate struct {
	AggregationKey string             `json:"aggregation_key"`
	TotalCount     int                `json:"total_count"`
	Severity       string             `json:"severity"`
	LatestErrors   []ErrorWithContext `json:"latest_errors"`
	CreatedAt      time.Time          `json:"created_at"` // ignored
}

type ErrorWithContext struct {
	Error       ErrorInstance `json:"error"`
	UUID        string        `json:"uuid"`      // ignored
	Timestamp   time.Time     `json:"timestamp"` // ignored
	Severity    string        `json:"severity"`
	HTTPContext *HTTPContext  `json:"http_context"`
}

type ErrorInstance struct {
	Class      string         `json:"class"`
	Message    string         `json:"message"`
	Stacktrace []string       `json:"stacktrace"`
	Cause      *ErrorInstance `json:"cause"`
}

type HTTPContext struct {
	RequestMethod  string            `json:"request_method"`
	RequestURL     string            `json:"request_url"`
	RequestHeaders map[string]string `json:"request_headers"`
	RequestBody    string            `json:"request_body"`
}

// generateErrorInstance generates a periskop.ErrorInstance from ErrorInstance
func generateErrorInstance(errInstance ErrorInstance) periskop.ErrorInstance {
	periskopErrInstance := periskop.ErrorInstance{
		Class:      errInstance.Class,
		Message:    errInstance.Message,
		Stacktrace: errInstance.Stacktrace,
	}
	// propagates the exception cause in case the exception has it
	if errInstance.Cause != nil {
		cause := generateErrorInstance(*errInstance.Cause)
		periskopErrInstance.Cause = &cause
	}
	return periskopErrInstance
}

func ToPeriskopErrorWithContext(errWithContext ErrorWithContext) periskop.ErrorWithContext {
	errInstance := generateErrorInstance(errWithContext.Error)

	httpContext := periskop.HTTPContext{}
	if errWithContext.HTTPContext != nil {
		httpContext = periskop.HTTPContext{
			RequestMethod:  errWithContext.HTTPContext.RequestMethod,
			RequestURL:     errWithContext.HTTPContext.RequestURL,
			RequestHeaders: errWithContext.HTTPContext.RequestHeaders,
			RequestBody:    &errWithContext.HTTPContext.RequestBody,
		}
	}

	return periskop.NewErrorWithContext(errInstance, periskop.Severity(errWithContext.Severity), &httpContext)
}
