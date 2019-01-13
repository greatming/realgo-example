package errors

const (
	SUCCESS_ERRNO = 0
	PARAMS_ERRNO = 2
	DB_ERRNO = 14
	UNKNOWN_ERRNO = 18000
)

type AppError struct {
	Status int    `json:"-"`
	Errno  int    `json:"errno"`
	ErrMsg string `json:"error"`
}
func (t *AppError) Error() string {
	return t.ErrMsg
}
var (
	SUCCESS = AppError{200, SUCCESS_ERRNO, ""}
	PARAMS_ERROR  = AppError{200, PARAMS_ERRNO, "Param error"}
	DB_ERROR      = AppError{200, DB_ERRNO, "db error"}
	UNKNOWN_ERROR = AppError{200, UNKNOWN_ERRNO, "unknown error"}
)
func New(msg string) *AppError {
	return &AppError{200, UNKNOWN_ERRNO, msg}
}
func NewError(errno int, msg string) *AppError {
	return &AppError{200, errno, msg}
}
func FormatNetdiskError(err interface{}) AppError {
	result, ok := err.(*AppError)
	if !ok {
		return UNKNOWN_ERROR
	}
	return *result
}

