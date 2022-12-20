package errors

type Fault string

const (
	FaultUnknown Fault = ""
	FaultClient  Fault = "FaultClient"
	FaultLib     Fault = "FaultLib"
)

type LibErr struct {
	Fault Fault
	Err   error
	Msg   string
}

func (e LibErr) Error() string {
	msg := e.Msg
	if e.Err != nil {
		if msg != "" {
			msg += ": " + e.Err.Error()
		} else {
			msg = e.Err.Error()
		}
	}
	return msg
}

func (e LibErr) Unwrap() error {
	return e.Err
}
