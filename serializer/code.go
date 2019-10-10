package serializer

// 定义所有状态码
const (
	// 用户不存在
	UserNotFoundError = 40000

	// 用户密码错误
	UserPasswordError = 40001

	// 用户无权限查看此资源 (需要登录)
	UserNotPermissionError = 40002

	// 用户输入不合法
	UserInputError = 40003

	// 用户重复错误
	UserRepeatError = 40004
)

const (
	// 严重的错误
	ServerPanicError = 50000

	// 数据库写入错误
	DatabaseWriteError = 50001

	// 数据库读取错误
	DatabaseReadError = 50002
)
