package glapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	RawResponseWriter http.ResponseWriter
}

func (res *Response) Send(r interface{}) {

	switch str := r.(type) {
	case string:
		res.SendText(str)
		return
	}

	res.SendJson(r)
}

func (res *Response) SendText(text string) {
	fmt.Fprintf(res.RawResponseWriter, text)
}

func (res *Response) SendJson(r interface{}) {
	j, _ := json.MarshalIndent(r, "", "  ")
	res.SendText(string(j[:]))
}
