package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goexl/gox"
	"github.com/pangum/casdoor/internal/config"
	"github.com/pangum/casdoor/internal/internal/constant"
)

type Callback struct {
	config *config.Callback
	mux    *http.ServeMux
	client *Client
}

func NewCallback(config *config.Callback, mux *http.ServeMux, client *Client) *Callback {
	return &Callback{
		config: config,
		mux:    mux,
		client: client,
	}
}

func (c *Callback) Handle() {
	c.mux.HandleFunc(gox.StringBuilder(constant.Slash, c.config.Path).String(), c.serve)
}

func (c *Callback) serve(writer http.ResponseWriter, req *http.Request) {
	code := ""
	state := ""
	if pfe := req.ParseForm(); nil != pfe {
		// TODO 日志
	} else {
		code = req.Form.Get(c.config.Code)
		state = req.Form.Get(c.config.State)
	}

	if "" == code || "" == state {
		// TODO 日志
	} else {
		c.parse(writer, code, state)
	}

	return
}

func (c *Callback) parse(writer http.ResponseWriter, code string, state string) {
	if token, gte := c.client.GetOAuthToken(code, state); nil != gte {
		// TODO 日志
		fmt.Println(gte)
	} else if claims, pte := c.client.ParseJwtToken(token.AccessToken); nil != pte {
		// TODO 日志
		fmt.Println(pte)
	} else {
		bytes, _ := json.Marshal(claims)
		writer.Write(bytes)
	}
}
