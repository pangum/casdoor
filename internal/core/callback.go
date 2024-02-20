package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pangum/casdoor/internal/config"
)

type Callback struct {
	config *config.Callback
	mux    *http.ServeMux
}

func NewCallback(config *config.Callback, mux *http.ServeMux) *Callback {
	return &Callback{
		config: config,
		mux:    mux,
	}
}

func (c *Callback) Handle() {
	c.mux.HandleFunc(c.config.Path, c.serve)
}

func (c *Callback) serve(writer http.ResponseWriter, req *http.Request) {
	code := ""
	state := ""
	switch req.Method {
	case "GET", "DELETE":
		code, state = c.parseForm(&req.Form)
	case "POST", "PUT":
		code, state = c.parse(req)
	}
	fmt.Println(code, state)
}

func (c *Callback) parse(req *http.Request) (code string, state string) {
	if body, gbe := req.GetBody(); nil != gbe {
		code, state = c.parseForm(&req.PostForm)
	} else {
		code, state = c.parseBody(body)
	}

	return
}

func (c *Callback) parseBody(body io.ReadCloser) (code string, state string) {
	if bytes, rae := io.ReadAll(body); nil != rae {
		// 不处理，后续步骤处理
	} else {
		code, state = c.parseJSON(&bytes)
	}

	return
}

func (c *Callback) parseJSON(bytes *[]byte) (code string, state string) {
	content := make(map[string]string)
	if ue := json.Unmarshal(*bytes, &content); nil == ue {
		code = content[c.config.Code]
		state = content[c.config.State]
	}

	return
}

func (c *Callback) parseForm(form *url.Values) (code string, state string) {
	if form.Has(c.config.Code) {
		code = form.Get(c.config.Code)
	}
	if form.Has(c.config.State) {
		state = form.Get(c.config.State)
	}

	return
}
