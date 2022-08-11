package constants

const (
	DomainErrCodeRequired = "DOMAIN:REQUIRED"
	DomainErrMsgRequired  = "Domain error: Missing required fields"

	DomainErrCodeNotAcceptedValue = "DOMAIN:NOT_ACCEPTED_VALUE"
	DomainErrMsgNotAcceptedValue  = "Domain error: Input is not in the accepted values"

	DomainErrCodeInvalidFormat = "DOMAIN:INVALID_FORMAT"
	DomainErrMsgInvalidFormat  = "Domain error: Input has an invalid format"

	DomainErrCodeUnauthenticated = "DOMAIN:UNAUTHENTICATED"
	DomainErrMsgUnauthenticated  = "Domain error: Unauthenticated"

	DomainErrCodeNotFound = "DOMAIN:NOT_FOUND"
	DomainErrMsgNotFound  = "Domain error: Not found"

	DomainErrCodeUnknown = "DOMAIN:UNKNOWN"
	DomainErrMsgUnknown  = "Domain error: Internal Server Error"
)
