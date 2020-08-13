# validator-zh

中文的参数验证。

# 使用方法

以 gin 为例其他框架类似. echo 框架可以直接复制到 echo.New()的实例上

```go
import(
  zh "github.com/glepnir/validatorzh"
)

var users User
_ = c.ShoudBindBodyWith(users,binding.JSON)
v := new(zh.ValidatorZh)
err := v.Validate(users)
if err !=nil{
  c.json(http.StatusBadRequest,gin.H{
    "message": err.Error(),
  })
}
```
