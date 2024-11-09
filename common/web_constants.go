package common

// web常用常量
const (
	ErrInvalidParam          = "请求参数错误！"
	ErrUserExists            = "用户已存在！"
	ErrCartExists            = "购物车已存在！"
	ErrOrderExists           = "订单已存在！"
	ErrUserNotFound          = "用户不存在！"
	ErrOrderNotFound         = "订单不存在！"
	ErrPasswordMismatch      = "密码不一致！"
	ErrGRPCConnFailed        = "grpc连接失败！"
	ErrUserCreationFailed    = "用户注册失败！"
	ErrLoginFailed           = "登录失败！"
	ErrEncryptionFailed      = "密码加密失败！"
	ErrInternalServer        = "服务器内部错误！"
	ErrDeletionFailed        = "删除失败！"
	ErrUpdateFailed          = "修改失败！"
	ErrGenerateJWTFailed     = "生成令牌失败！"
	ErrVerifyJWTFailed       = "验证令牌失败！"
	ErrDiscoverServiceFailed = "发现服务失败！"
	ErrFindEmailFailed       = "邮箱未找到！"

	MsgRegistrationSuccess = "注册成功！"
	MsgSelectionSuccess    = "查询成功！"
	MsgDeletionSuccess     = "删除成功！"
	MsgUpdateSuccess       = "修改成功！"
	MsgLoginSuccess        = "登录成功！"
	MsgSuccess             = "success"
)
