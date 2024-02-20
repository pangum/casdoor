package core

import (
	"github.com/pangum/casdoor/internal/config"
)

type Config struct {
	Endpoint     string           `json:"endpoint,omitempty" yaml:"endpoint" xml:"endpoint" toml:"endpoint" validate:"required,url"`
	Id           string           `json:"id,omitempty" yaml:"id" xml:"id" toml:"id" validate:"required,url"`
	Secret       string           `json:"secret,omitempty" yaml:"secret" xml:"secret" toml:"secret" validate:"required,url"`
	Certificate  string           `json:"certificate,omitempty" yaml:"certificate" xml:"certificate" toml:"certificate" validate:"required,url"`
	Organization string           `json:"organization,omitempty" yaml:"organization" xml:"organization" toml:"organization" validate:"required,url"`
	Application  string           `json:"application,omitempty" yaml:"application" xml:"application" toml:"application" validate:"required,url"`
	Callback     *config.Callback `json:"serve,omitempty" yaml:"serve" xml:"serve" toml:"serve"`
}
