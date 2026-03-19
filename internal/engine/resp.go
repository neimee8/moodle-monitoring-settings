package engine

type Resp struct {
	Status int
	Msg    string
	Data   any
}

func NewResp(msg string, data any, status ...int) Resp {
	s := 200

	if len(status) > 0 {
		s = status[0]
	}

	return Resp{
		Status: s,
		Msg:    msg,
		Data:   data,
	}
}
