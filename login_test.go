package weapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"openid":      "mock-openid",
			"session_key": "mock-session_key",
			"unionid":     "mock-unionid",
			"errcode":     0,
			"errmsg":      "mock-errmsg",
		}
		bts, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		w.Write(bts)
		if r.Method != "GET" {
			t.Fatalf("Except 'GET' got '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiLogin {
			t.Fatalf("Except to path '%s',got '%s'", apiLogin, path)
		}

		r.ParseForm()
		queries := []string{"appid", "secret", "js_code", "grant_type"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}

		}
	}))
	defer ts.Close()

	_, err := login("mock-appid", "mock-secret", "mock-code", ts.URL+apiLogin)
	if err != nil {
		t.Fatal(err)
	}
}