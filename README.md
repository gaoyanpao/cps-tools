# 集成各cps平台的api接口

## 未完成，其他api待补充
example: 
```golang
import "github.com/gaoyanpao/cps-tools/taobao"

func main() {
	taobao.SetAppInfo("your appkey", "your secret")
	req := &taobao.TbkActivityInfoGetRequest{
		ActivityMaterialID: "",
		ADZoneID:           "",
		SubID:              "",
	}
	taobao.Call(req)
	respData := req.GetResult()
}

```