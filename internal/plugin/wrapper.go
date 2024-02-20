package plugin

import (
	"github.com/pangum/casdoor/internal/core"
)

type Wrapper struct {
	Casdoor *core.Config `json:"casdoor" yaml:"casdoor" xml:"casdoor" toml:"casdoor" validate:"required"`
}
