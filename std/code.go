package std

type Err struct {
	Code    int
	Message string
}

var (
	ErrOK          = newError(20000, "OK")
	ErrIllegalJson = newError(50001, "Illegal JSON")
)

func newError(code int, message string) *Err {
	return &Err{Code: code, Message: message}
}
