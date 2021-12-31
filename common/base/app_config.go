// Package base
// @Author: Nick Bi
// @Date:
// @Desc: todo...
///**
package base

type Config struct {
	AutoLoadConfig bool
	Server         Server
	Database       Database
	Redis          Redis
	Log            Log
}

type Server struct {
	Port        string
	Name        string
	ContextPath string
	PublicUrl   []string
}

type Database struct {
	Mysql Mysql
}

type Mysql struct {
	Url string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	Db       string
}

type Log struct {
	Level      string
	Path       string
	Sugar      bool
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	PackLog    bool
}
