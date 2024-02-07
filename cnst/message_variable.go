package cnst

// contain key in env file

const (
	SignupOk      = "sign up success"
	LoginError    = "authentication failed"
	LoginSuccess  = "login successful"
	UpdateSuccess = "update successful"

	ErrParseConfigYaml  = "error on parsing configuration file"
	ErrReqPathParamUuid = "require path param <uuid>"
	ErrTokenExpired = "token expired"
	ErrreqToken = "Authentication header is missing, require token"
)
