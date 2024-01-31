package cnst



//init error
//get error message , return 

var errMap = make(map[int]string)


func InitErr(){
	errMap[4004] = "The user was not found or the password was incorrect. The response will be empty."
}


func GetErrMsg(errcode int) string{
	msg := ""

	errStr, ok := errMap[errcode]
	if ok {
		msg = errStr
	} 
	return msg
}