# axios

下载
```bash
go get -u github.com/orangbus/axios@latest
```
2、Register service provider
```go
// config/app.go
import "github.com/orangbus/axios"

"providers": []foundation.ServiceProvider{
    &axios.ServiceProvider{},
},

3、 Publish Configuration
```bash
go run . artisan vendor:publish --package=github.com/orangbus/axios
```
## 快速开始
```go
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
```
## 主要功能
- [x] get请求
- [x] post请求
- [x] postform请求
- [ ] header,base_auth,proxy等快捷设置
