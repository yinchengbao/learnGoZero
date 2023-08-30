package config

import "github.com/zeromicro/go-zero/rest"

type DB struct {
	DataSource string
}

type Auth struct {
	AccessSecret string
	AccessExpire int64
}

type Config struct {
	rest.RestConf
	Db   DB
	Auth Auth
}
