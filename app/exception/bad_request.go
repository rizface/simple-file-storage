package exception

type BadRequest struct {
	Error interface{}
}

func PanicIfBadRequest(err interface{}) {
	if err != nil {
		panic(BadRequest{Error: err.(error).Error()})
	}
}
