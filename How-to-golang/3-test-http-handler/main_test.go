package __test_http_handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetFoo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handGetFoo))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("got %d, want %d", resp.StatusCode, http.StatusOK)
	}
	defer resp.Body.Close()
	expected := "satyanarayan jadhav"
	b, err := ioutil.ReadAll(resp.Body)
	if string(b) != expected {
		t.Errorf("expected %s got %v", expected, string(b))
	}
}
