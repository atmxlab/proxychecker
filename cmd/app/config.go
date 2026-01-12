package app

import "github.com/atmxlab/proxychecker/pkg/validator"

type Config struct {
	Queue  Queue  `atmc:"queue"`
	API    API    `atmc:"api"`
	Logger Logger `atmc:"logger"`
	Env    Env    `atmc:"env"`
}

func (c Config) Validate() error {
	v := validator.New()

	v.WrapErr(c.Env.Validate(), "Env.Validate")

	return v.Err()
}

type Logger struct {
	Level string `atmc:"level"`
}

type Queue struct {
	QueueWorkerCount uint `atmc:"queueWorkerCount"`
	QueueBufferSize  int  `atmc:"queueBufferSize"`
}

type API struct {
	GRPC GRPC `atmc:"grpc"`
	HTTP HTTP `atmc:"http"`
}

type GRPC struct {
	Port uint16 `atmc:"port"`
}

type HTTP struct {
	Port    uint16  `atmc:"port"`
	Swagger Swagger `atmc:"swagger"`
}

type Swagger struct {
	UI   SwaggerUI   `atmc:"ui"`
	JSON SwaggerJSON `atmc:"json"`
}
type SwaggerUI struct {
	URL        string `atmc:"url"`
	SourcePath string `atmc:"sourcePath"`
}
type SwaggerJSON struct {
	URL  string `atmc:"url"`
	Path string `atmc:"path"`
}

type Env struct {
	ServerIP string `atmc:"serverIP"`
	ENV      string `atmc:"env"`
}

func (env Env) Validate() error {
	v := validator.New()

	if env.ServerIP == "" {
		v.Failed("server ip is empty")
	}

	if env.ENV == "" {
		v.Failed("env is empty")
	}

	return v.Err()
}
