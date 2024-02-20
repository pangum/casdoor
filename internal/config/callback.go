package config

type Callback struct {
	Enabled bool `json:"enabled,omitempty" yaml:"enabled" xml:"enabled" toml:"enabled"`
	// nolint:lll
	Path  string `default:"casdoor/callback" json:"path,omitempty" yaml:"path" xml:"path" toml:"path" validate:"required_if=Enabled true,startsnotwith=/"`
	Code  string `default:"code" json:"code,omitempty" yaml:"code" xml:"code" toml:"code"`
	State string `default:"state" json:"state,omitempty" yaml:"state" xml:"state" toml:"state"`
	// 是否设置
	Uid *bool `default:"true" json:"uid,omitempty" yaml:"uid" xml:"uid" toml:"uid"`
}
