package component

type Minio struct {
	ApiURL          string `json:"apiURL"`
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"accessKeyID"`
	SecretAccessKey string `json:"secretAccessKey"`
	SignEndpoint    string `json:"signEndpoint"`
	UseSSL          string `json:"useSSL"`
}

type Zookeeper struct {
	Schema   string   `json:"schema"`
	ZkAddr   []string `json:"zkAddr"`
	Username string   `json:"username"`
	Password string   `json:"password"`
}

type MySQL struct {
	Address  []string `json:"address"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Database string   `json:"database"`
}

type Kafka struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Addr     []string `json:"addr"`
}
