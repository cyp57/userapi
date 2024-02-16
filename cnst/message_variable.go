package cnst

// contain key in env file

const (
	SignupOk      = "Sign up success."
	LoginError    = "Authentication failed."
	LoginSuccess  = "Login successful."
	UpdateSuccess = "Update successful."
	DeleteSuccess = "Delete successful."
	RequestSuccess	= "The request was successful."

	ForgotPasswordSuccess = "Password recovery request received. Please check your email ."
	ChangePasswordSuccess = "Password successfully changed."

	ErrParseConfigYaml  = "Error on parsing configuration file."
	ErrReqPathParamUuid = "Require path param <uuid>."
	ErrTokenExpired     = "Token expired."
	ErrReqRole          = "User require role."
	ErrAuthorizeRole          = "Error authorize role."
	ErrReqToken         = "Authentication header is missing, require token."
	ErrSortKeyReq       = "Error , sortKey require value of sort."
	ErrValidUuid        = "uuid and token must be same user."
	ErrInvalidRole      = "Invalid roles type"
	ErrForbidden		= "You don't have permission to access this resource."
	ErrappIdNotFound  = "Fusionauth applicationId not found."
)
