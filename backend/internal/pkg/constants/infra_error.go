package constants

const (
	InfraErrCodeDBConnect = "INFRA:DATABASE:CONNECT"
	InfraErrMsgDBConnect  = "Infra error: Failed to connect to database"

	InfraErrCodeDBNotFound = "INFRA:DATABASE:NOT_FOUND"
	InfraErrMsgDBNotFound  = "Infra error: Not found resource in database"

	InfraErrCodeDBSelect = "INFRA:DATABASE:SELECT"
	InfraErrMsgDBSelect  = "Infra error: Failed to select resources from database"

	InfraErrCodeDBInsert = "INFRA:DATABASE:INSERT"
	InfraErrMsgDBInsert  = "Infra error: Failed to insert into database"

	InfraErrCodeDBUpdate = "INFRA:DATABASE:UPDATE"
	InfraErrMsgDBUpdate  = "Infra error: Failed to update resources in database"

	InfraErrCodeDBDelete = "INFRA:DATABASE:DELETE"
	InfraErrMsgDBDelete  = "Infra error: Failed to delete resources from database"

	InfraErrCodeDBUnknown = "INFRA:DATABASE:UNKNOWN"
	InfraErrMsgDBUnknown  = "Infra error: Unknown database error"
)
