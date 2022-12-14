package constants

const (
	STATUS_SUCCESS       = "00"
	STATUS_STILL_PROCESS = "01"
	STATUS_NOT_FOUND     = "02"
	STATUS_CONFLICT      = "03"
	STATUS_FAILED        = "05"
)

const (
	MESSAGE_SUCCESS       = "Success"
	MESSAGE_STILL_PROCESS = "Transaction is being process"
	MESSAGE_FAILED        = "Something went wrong"
	MESSAGE_NOT_FOUND     = "Not Found"
	MESSAGE_CONFLICT      = "Already Exist"
)

const (
	REDIS_STUDENT = "student"
	REDIS_CLASS   = "class"
)
