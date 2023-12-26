package config

import "github.com/fast/pkg/conf"

var Instance *Config

func Init(configFile string, fs ...func()) error {
	Instance = &Config{}
	return conf.Parse(configFile, Instance, fs...)
}

func Show(hiddenFields ...string) string {
	return conf.Show(Instance, hiddenFields...)
}

type Config struct {
	Server Server `yaml:"server" json:"server"`
	HTTP   HTTP   `yaml:"http" json:"http"`
	OTEL   OTEL   `yaml:"otel" json:"otel"`
	Logger Logger `yaml:"logger" json:"logger"`
	Mysql  Mysql  `yaml:"mysql" json:"mysql"`
	Redis  Redis  `yaml:"redis" json:"redis"`
}

type Server struct {
	CacheType             string  `yaml:"cacheType" json:"cacheType"`
	EnableCircuitBreaker  bool    `yaml:"enableCircuitBreaker" json:"enableCircuitBreaker"`
	EnableHTTPProfile     bool    `yaml:"enableHTTPProfile" json:"enableHTTPProfile"`
	EnableLimit           bool    `yaml:"enableLimit" json:"enableLimit"`
	EnableMetrics         bool    `yaml:"enableMetrics" json:"enableMetrics"`
	EnableStat            bool    `yaml:"enableStat" json:"enableStat"`
	EnableTrace           bool    `yaml:"enableTrace" json:"enableTrace"`
	Env                   string  `yaml:"env" json:"env"`
	Host                  string  `yaml:"host" json:"host"`
	Name                  string  `yaml:"name" json:"name"`
	RegistryDiscoveryType string  `yaml:"registryDiscoveryType" json:"registryDiscoveryType"`
	TracingSamplingRate   float64 `yaml:"tracingSamplingRate" json:"tracingSamplingRate"`
	Version               string  `yaml:"version" json:"version"`
}

type OTEL struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
}

type Mysql struct {
	ConnMaxLifetime int      `yaml:"connMaxLifetime" json:"connMaxLifetime"`
	Dsn             string   `yaml:"dsn" json:"dsn"`
	EnableLog       bool     `yaml:"enableLog" json:"enableLog"`
	MastersDsn      []string `yaml:"mastersDsn" json:"mastersDsn"`
	MaxIdleConns    int      `yaml:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int      `yaml:"maxOpenConns" json:"maxOpenConns"`
	SlavesDsn       []string `yaml:"slavesDsn" json:"slavesDsn"`
	SlowThreshold   int      `yaml:"slowThreshold" json:"slowThreshold"`
}

type Redis struct {
	DialTimeout  int    `yaml:"dialTimeout" json:"dialTimeout"`
	Dsn          string `yaml:"dsn" json:"dsn"`
	ReadTimeout  int    `yaml:"readTimeout" json:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout" json:"writeTimeout"`
}

type Logger struct {
	Format string `yaml:"format" json:"format"`
	IsSave bool   `yaml:"isSave" json:"isSave"`
	Level  string `yaml:"level" json:"level"`
}

type HTTP struct {
	Port         int `yaml:"port" json:"port"`
	ReadTimeout  int `yaml:"readTimeout" json:"readTimeout"`
	WriteTimeout int `yaml:"writeTimeout" json:"writeTimeout"`
}
