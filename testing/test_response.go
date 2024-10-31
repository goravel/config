package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	contractstesting "github.com/goravel/framework/contracts/testing"
	"github.com/goravel/framework/support/carbon"
)

type TestResponseImpl struct {
	t        *testing.T
	response *http.Response
	content  string
}

func NewTestResponse(t *testing.T, response *http.Response) contractstesting.TestResponse {
	//content := recorder.Body.String()
	return &TestResponseImpl{t: t, response: response}
}

func (r *TestResponseImpl) AssertStatus(status int) contractstesting.TestResponse {
	actual := r.getStatusCode()
	assert.Equal(r.t, status, actual, fmt.Sprintf("Expected response status code [%d] but received %d.", status, actual))
	return r
}

func (r *TestResponseImpl) AssertOk() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusOK)
}

func (r *TestResponseImpl) AssertCreated() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusCreated)
}

func (r *TestResponseImpl) AssertAccepted() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusAccepted)
}

func (r *TestResponseImpl) AssertNoContent(status ...int) contractstesting.TestResponse {
	expectedStatus := http.StatusNoContent
	if len(status) > 0 {
		expectedStatus = status[0]
	}

	r.AssertStatus(expectedStatus)

	assert.Empty(r.t, r.content)

	return r
}

func (r *TestResponseImpl) AssertMovedPermanently() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusMovedPermanently)
}

func (r *TestResponseImpl) AssertFound() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusFound)
}

func (r *TestResponseImpl) AssertNotModified() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusNotModified)
}

func (r *TestResponseImpl) AssertPartialContent() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusPartialContent)
}

func (r *TestResponseImpl) AssertTemporaryRedirect() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusTemporaryRedirect)
}

func (r *TestResponseImpl) AssertBadRequest() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusBadRequest)
}

func (r *TestResponseImpl) AssertUnauthorized() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusUnauthorized)
}

func (r *TestResponseImpl) AssertPaymentRequired() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusPaymentRequired)
}

func (r *TestResponseImpl) AssertForbidden() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusForbidden)
}

func (r *TestResponseImpl) AssertNotFound() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusNotFound)
}

func (r *TestResponseImpl) AssertMethodNotAllowed() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusMethodNotAllowed)
}

func (r *TestResponseImpl) AssertNotAcceptable() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusNotAcceptable)
}

func (r *TestResponseImpl) AssertConflict() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusConflict)
}

func (r *TestResponseImpl) AssertRequestTimeout() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusRequestTimeout)
}

func (r *TestResponseImpl) AssertGone() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusGone)
}

func (r *TestResponseImpl) AssertUnsupportedMediaType() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusUnsupportedMediaType)
}

func (r *TestResponseImpl) AssertUnprocessableEntity() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusUnprocessableEntity)
}

func (r *TestResponseImpl) AssertTooManyRequests() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusTooManyRequests)
}

func (r *TestResponseImpl) AssertInternalServerError() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusInternalServerError)
}

func (r *TestResponseImpl) AssertServiceUnavailable() contractstesting.TestResponse {
	return r.AssertStatus(http.StatusServiceUnavailable)
}

func (r *TestResponseImpl) AssertHeader(headerName, value string) contractstesting.TestResponse {
	got := r.getHeader(headerName)
	assert.NotEmpty(r.t, got, fmt.Sprintf("Header [%s] not present on response.", headerName))
	if got != "" {
		assert.Equal(r.t, value, got, fmt.Sprintf("Header [%s] was found, but value [%s] does not match [%s].", headerName, got, value))
	}
	return r
}

func (r *TestResponseImpl) AssertHeaderMissing(headerName string) contractstesting.TestResponse {
	got := r.getHeader(headerName)
	assert.Empty(r.t, got, fmt.Sprintf("Unexpected header [%s] is present on response.", headerName))
	return r
}

func (r *TestResponseImpl) AssertCookie(name, value string) contractstesting.TestResponse {
	cookie := r.getCookie(name)
	assert.NotNil(r.t, cookie, fmt.Sprintf("Cookie [%s] not present on response.", name))

	if cookie == nil {
		return r
	}

	assert.Equal(r.t, value, cookie.Value, fmt.Sprintf("Cookie [%s] was found, but value [%s] does not match [%s]", name, cookie.Value, value))

	return r
}

func (r *TestResponseImpl) AssertCookieExpired(name string) contractstesting.TestResponse {
	cookie := r.getCookie(name)
	assert.NotNil(r.t, cookie, fmt.Sprintf("Cookie [%s] not present on response.", name))

	if cookie == nil {
		return r
	}

	expirationTime := carbon.FromStdTime(cookie.Expires)
	assert.True(r.t, !(cookie.MaxAge > 0 || (!expirationTime.IsZero() && expirationTime.Gt(carbon.Now()))), fmt.Sprintf("Cookie [%s] is not expired; it expires at [%s].", name, expirationTime.ToString()))

	return r
}

func (r *TestResponseImpl) AssertCookieNotExpired(name string) contractstesting.TestResponse {
	cookie := r.getCookie(name)
	assert.NotNil(r.t, cookie, fmt.Sprintf("Cookie [%s] not present on response.", name))

	if cookie == nil {
		return r
	}

	expirationTime := carbon.FromStdTime(cookie.Expires)
	assert.True(r.t, cookie.MaxAge > 0 || (!expirationTime.IsZero() && expirationTime.Gt(carbon.Now())), fmt.Sprintf("Cookie [%s] is expired; it expired at [%s].", name, expirationTime))
	return r
}

func (r *TestResponseImpl) AssertCookieMissing(name string) contractstesting.TestResponse {
	assert.Nil(r.t, r.getCookie(name), fmt.Sprintf("Cookie [%s] is present on response.", name))

	return r
}

func (r *TestResponseImpl) getStatusCode() int {
	return r.response.StatusCode
}

func (r *TestResponseImpl) getCookie(name string) *http.Cookie {
	for _, c := range r.response.Cookies() {
		if c.Name == name {
			return c
		}
	}

	return nil
}

func (r *TestResponseImpl) getHeader(name string) string {
	return r.response.Header.Get(name)
}
