package config

import "github.com/zeromicro/go-zero/rest"

type DB struct {
	DataSource string
}

type Config struct {
	rest.RestConf
	Db DB
}
