package httpResponse

type HttpResponse struct {
	StatusCode    string
	ResuluNessage string
	Data          interface{}
}

func SuccessResponse(Data interface{} )  HttpResponse  {
	return HttpResponse{
		StatusCode: "200",
		ResuluNessage: "success",
		Data: Data,
	}
}

func ErrorResponse(errorMessage string) HttpResponse  {
	return HttpResponse{
		StatusCode: "500",
		ResuluNessage: errorMessage,
		Data: nil,
	}
}

func NotFoundResponse(errorMessage string) HttpResponse  {
	return HttpResponse{
		StatusCode: "404",
		ResuluNessage: errorMessage,
		Data: nil,
	}
}