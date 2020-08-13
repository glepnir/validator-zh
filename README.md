# validator-zh

中文的参数验证。

# 使用方法

以 gin 为例其他框架类似. echo 框架可以直接复制到 echo.New()的实例上

```go
import(
  zh "github.com/glepnir/validatorzh"
)

type CreateUserSchema struct {
	UserName       string `json:"username" validate:"required" label:"用户姓名"`
	Phone          string `json:"phone" validate:"required,mobile" label:"联系电话"`
	IdCard         string `json:"phone" validate:"required,idcard" label:"身份证号码"`
}

var users User
_ = c.ShoudBindBodyWith(users,binding.JSON)
v := new(zh.ValidatorZh)
err := v.Validate(users)
if err !=nil{
  c.json(http.StatusBadRequest,gin.H{
    "message": err.Error(),
  })
}
---------output:
"用户姓名为必填字段"
"联系电话格式错误"
"身份证号码格式错误"
```
