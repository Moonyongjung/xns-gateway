package types

type XnsError struct {
	code uint64
	desc string
}

const (
	ErrMsgInvalidDataType = "invalid data type of the message"
	ErrMsgDuplicateExpire = "invalid duplicated expire type"
)

var (
	Success           = new(0, "success")
	ErrInit           = new(101, "error initialize")
	ErrParseConfig    = new(102, "error parsing config.yaml file")
	ErrParseApp       = new(103, "error parsing app file")
	ErrContract       = new(104, "error contract")
	ErrKey            = new(105, "error key")
	ErrHaveNoAddress  = new(106, "error no contract address")
	ErrBroadcast      = new(107, "error broadcast")
	ErrTx             = new(108, "error create transaction")
	ErrQuery          = new(109, "error query")
	ErrNewMsg         = new(110, "error generate new message")
	ErrParseData      = new(111, "error parse data")
	ErrGW             = new(112, "error gateway")
	ErrHttpServer     = new(113, "error HTTP server")
	ErrInvalidRequest = new(114, "error invalid request")
)

func new(code uint64, desc string) (x XnsError) {
	x.code = code
	x.desc = desc

	return x
}

func (x XnsError) Code() uint64 {
	return x.code
}

func (x XnsError) Desc() string {
	return x.desc
}
