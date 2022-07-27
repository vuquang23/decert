package constants

const (
	//HTTP 200 - OK
	RestCodeOK = 0
	RestMsgOK  = "Succeeded"

	//HTTP 400 - Bad Request
	RestErrCodeRequired = 4000
	RestErrMsgRequired  = "Missing required fields"

	RestErrCodeNotAcceptedValue = 4001
	RestErrMsgNotAcceptedValue  = "Input is not in the accepted values"

	RestErrCodeInvalidFormat = 4002
	RestErrMsgInvalidFormat  = "Input has an invalid format"

	//HTTP 401 - Unauthorized
	RestErrCodeUnauthenticated = 4010
	RestErrMsgUnauthenticated  = "Unauthenticated"

	//HTTP 404 - Not found
	RestErrCodeNotFound = 4040
	RestErrMsgNotFound  = "Not found"

	//HTTP 500 - Internal Server Error
	RestErrCodeInternal = 5000
	RestErrMsgInternal  = "Internal Server Error"
)
