package cnst

// contain key in env file

const (
	SignupOk      = "sign up success"
	LoginError    = "authentication failed"
	LoginSuccess  = "login successful"
	UpdateSuccess = "update successful"
	DeleteSuccess = "delete successful"

	ErrParseConfigYaml  = "error on parsing configuration file"
	ErrReqPathParamUuid = "require path param <uuid>"
	ErrTokenExpired = "token expired"
	ErrReqToken = "Authentication header is missing, require token"
	ErrSortKeyReq = "error , sortKey require value of sort"
)
