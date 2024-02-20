package plugin

import (
	"net/http"

	"github.com/goexl/exception"
	"github.com/pangum/casdoor/internal/config"
	"github.com/pangum/casdoor/internal/core"
	"github.com/pangum/pangu"
)

type Constructor struct {
	// 解决命名空间问题
	mux *http.ServeMux
}

func (c *Constructor) New(app *pangu.App, config *pangu.Config) (client *core.Client, err error) {
	wrapper := new(Wrapper)
	if ge := config.Build().Get(wrapper); nil != ge {
		err = ge
	} else {
		client, err = c.new(app, wrapper.Casdoor)
	}

	return
}

func (c *Constructor) new(app *pangu.App, config *core.Config) (client *core.Client, err error) {
	client = core.NewClient(config)
	if nil != config.Callback {
		err = c.initCallback(app, config.Callback)
	}

	return
}

func (c *Constructor) initCallback(app *pangu.App, config *config.Callback) (err error) {
	if gie := app.Dependency().Get(c.httpServer).Build().Build().Inject(); nil != gie {
		err = exception.New().Message("未引入服务器依赖，可以使用pangum/grpc等插件来引入").Build()
	} else {
		callback := core.NewCallback(config, c.mux)
		callback.Handle()
	}

	return
}

func (c *Constructor) httpServer(mux *http.ServeMux) {
	c.mux = mux
}
