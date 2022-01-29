package model

type Response struct {
	Code int         `json:"code" yaml:"code" xml:"code"`
	Msg  string      `json:"msg" yaml:"msg" xml:"msg"`
	Data interface{} `json:"data" yaml:"data" xml:"data"`
}
