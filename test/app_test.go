package test

import (
	"github.com/orangbus/axios/bootstrap"
	"github.com/orangbus/axios/facades"
	"testing"
)

func init() {
	bootstrap.Boot()
}

func TestGet(t *testing.T) {
	param := map[string]interface{}{
		"page":  1,
		"limit": 10,
	}
	res, err := facades.Axios().Get("https://httpbin.org/get", param)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(res))
}

func TestPost(t *testing.T) {
	param := map[string]interface{}{
		"name": "orangbus",
	}
	res, err := facades.Axios().Post("https://httpbin.org/post", param)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(res))
}
func TestPostForm(t *testing.T) {
	param := map[string]interface{}{
		"name":     "orangbus",
		"password": 123456,
	}
	res, err := facades.Axios().PostForm("https://httpbin.org/post", param)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(res))
}
