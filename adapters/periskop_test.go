package adapters

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/periskop-dev/periskop-go"
)

func TestPeriskopAdapter(t *testing.T) {
	uuid, _ := uuid.Parse("5d9893c6-51d6-11ea-8aad-f894c260afe5")
	timestamp := time.Now().UTC()
	body := "body"
	errWithContext := ErrorWithContext{
		Error: ErrorInstance{
			Class:      "Test",
			Message:    "error!",
			Stacktrace: []string{"line 12:", "syntax error"},
		},
		UUID:      uuid.String(),
		Severity:  "error",
		Timestamp: timestamp,
		HTTPContext: &HTTPContext{
			RequestMethod:  "POST",
			RequestURL:     "http://example.com",
			RequestBody:    body,
			RequestHeaders: map[string]string{"Cache-Control": "no-cache"},
		},
	}
	expectedPeriskopErrWithContext := periskop.ErrorWithContext{
		Error: periskop.ErrorInstance{
			Class:      errors.New("Test").Error(),
			Message:    "error!",
			Stacktrace: []string{"line 12:", "syntax error"},
		},
		UUID:      uuid,
		Timestamp: timestamp,
		Severity:  periskop.SeverityError,
		HTTPContext: &periskop.HTTPContext{
			RequestMethod:  "POST",
			RequestURL:     "http://example.com",
			RequestHeaders: map[string]string{"Cache-Control": "no-cache"},
			RequestBody:    &body,
		},
	}
	generatedPeriskopErrWithContext := ToPeriskopErrorWithContext(errWithContext)
	// override the values since the values are auto generated
	generatedPeriskopErrWithContext.UUID = uuid
	generatedPeriskopErrWithContext.Timestamp = timestamp
	if !reflect.DeepEqual(expectedPeriskopErrWithContext, generatedPeriskopErrWithContext) {
		expected, _ := json.MarshalIndent(expectedPeriskopErrWithContext, "", "\t")
		generated, _ := json.MarshalIndent(generatedPeriskopErrWithContext, "", "\t")
		t.Errorf("error generating Periskop error:\nexpected %s\ngot %s",
			expected, generated)
	}
}
