package app

type Config struct {
	Queue Queue `atmc:"queue"`
	API   API   `atmc:"api"`
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
