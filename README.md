# validator-zh

[validator](https://github.com/go-playground/validator)的中文提示,新增手机号码和身份证号码.

# 使用方法

以 gin 为例其他框架类似.

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

err := zh.Validate(users)
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
