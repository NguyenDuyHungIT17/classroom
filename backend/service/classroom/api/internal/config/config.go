package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DataSource string

	//auth
	Auth struct {
		AccessSecret  string
		AccessExpire  int64
		RefreshExpire int64
	}

	SMTPConfig struct {
		ClientOrigin string
		EmailFrom    string
		SMTPHost     string
		SMTPUser     string
		SMTPPass     string
		SMTPPort     int64
	}
}
