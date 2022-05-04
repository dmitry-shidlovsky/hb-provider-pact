package provider

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dmitry-shidlovsky/TestPact/model"
)

func TestGetUser_unit(t *testing.T) {
	userRepository.TestInit()
	server := httptest.NewServer(commonMiddleware(GetUser))
	defer server.Close()

	u, _ := url.Parse(server.URL)
	req, err := http.NewRequest("GET", u.String()+"/user/1", nil)
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)

	decoder := json.NewDecoder(res.Body)

	var user *model.User
	err = decoder.Decode(&user)
	assert.NoError(t, err)

	pUser, err := userRepository.ByID(1)
	assert.NoError(t, err)
	assert.Equal(t, user, pUser)
}
