package config

import "time"

type Server struct {
	Env          string
	Name         string
	Version      string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
