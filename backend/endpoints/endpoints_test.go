package endpoints

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jakebjorke/sezzle-interview/models"
	"github.com/stretchr/testify/assert"
)

/*
Add tests for the use cases in the email.

If there is time come back and test failure modes
*/

func Test_Endpoints_5plus5(t *testing.T) {
	assert := assert.New(t)

	evalGood("5 + 5", "5 + 5 = 10", assert)
}

func Test_Endpoints_3times4(t *testing.T) {
	assert := assert.New(t)

	evalGood("3 * 4", "3 * 4 = 12", assert)
}

func evalGood(in, expected string, assert *assert.Assertions) {
	input := models.ExpressionRequest{Value: in}
	b, err := json.Marshal(input)
	if err != nil {
		assert.Fail("unable to create json input:  ", err)
	}

	req, err := http.NewRequest("POST", "/expression", bytes.NewBuffer(b))
	if err != nil {
		assert.Fail("unable to create request:  ", err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Expression)
	handler.ServeHTTP(recorder, req)

	assert.Equal(recorder.Code, http.StatusOK)
	var response models.ExpressionResponse
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(expected, response.Statement)
}
