package constant

const (
	Success = iota + 200
	DuplicatedUser
	UserNotFound
	WrongPassword
	ParamError  // 参数错误
	ServerError // 服务器错误
	NotFound    // 资源不存在
)
