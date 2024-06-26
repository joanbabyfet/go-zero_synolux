// Code generated by goctl. DO NOT EDIT.
package types

type CommonResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int         `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type SaveRequest struct {
	Mobile  string `form:"mobile"`
	Name    string `form:"name"`
	Email   string `form:"email"`
	Content string `form:"content"`
}
