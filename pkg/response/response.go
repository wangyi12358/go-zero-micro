package response

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *Response) Ok(data interface{}) *Response {
	return &Response{
		Code: 0,
		Msg:  "",
		Data: data,
	}
}

func Error() {

}
