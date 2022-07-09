package gateway_test

import (
	"encoding/xml"
	"github.com/semirm-dev/seeba/gateway"
	"github.com/semirm-dev/seeba/internal/web"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type apiResponse string

type mockedSearch struct {
}

func (srch *mockedSearch) All() (interface{}, error) {
	return "mocked music data", nil
}

func TestGetMusic(t *testing.T) {
	router := web.NewRouter()
	router.GET("music", gateway.GetMusic(&mockedSearch{}))

	req, _ := http.NewRequest("GET", "/music", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	respBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}

	var resp apiResponse
	if err = xml.Unmarshal(respBody, &resp); err != nil {
		t.Fail()
	}
	assert.NotNil(t, resp)
	assert.Equal(t, apiResponse("mocked music data"), resp)
}
