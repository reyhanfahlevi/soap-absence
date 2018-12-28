package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tokopedia/affiliate/pkg/env"
	yaml "gopkg.in/yaml.v2"
)

// Config container
type Config struct {
	App       AppSettings    `yaml:"app"`
	DB        DatabaseConfig `yaml:"db"`
	Scheduler Scheduler      `yaml:"scheduler"`
}

// DatabaseConfig db config
type DatabaseConfig struct {
	Master string `yaml:"master"`
	Slave  string `yaml:"slave"`
}

// AppSettings General App Settings
type AppSettings struct {
	ApiPort      string   `yaml:"api_port"`
	DeviceIPList []string `yaml:"device_ip_list"`
}

type Scheduler struct {
	SyncUserData      string `yaml:"sync_user_data"`
	PullAttendanceLog string `yaml:"pull_attendance_log"`
}

var (
	config *Config
)

// option defines configuration option
type option struct {
	configFile string
}

// Init initializes `config` from the default config file.
// use `WithConfigFile` to specify the location of the config file
func Init(opts ...Option) error {
	opt := &option{
		configFile: getDefaultConfigFile(),
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}

	out, err := ioutil.ReadFile(opt.configFile)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(out, &config)
}

// Option define an option for config package
type Option func(*option)

// WithConfigFile set `config` to use the given config file
func WithConfigFile(file string) Option {
	return func(opt *option) {
		opt.configFile = file
	}
}

// getDefaultConfigFile get default config file.
// - files/etc/soap-absence/soap-absence.development.yaml in dev
// - otherwise /etc/soap-absence/soap-absence.{TKPENV}.yaml
func getDefaultConfigFile() string {
	var (
		repoPath   = filepath.Join(os.Getenv("GOPATH"), "src/github.com/reyhanfahlevi/soap-absence")
		configPath = filepath.Join(repoPath, "files/etc/soap-absence/soap-absence.development.yaml")
		tkpEnv     = env.Get()
	)

	if tkpEnv != env.EnvDevelopment {
		configPath = fmt.Sprintf("/etc/soap-absence/soap-absence.%s.yaml", tkpEnv)
	}
	return configPath
}

// Get config
func Get() *Config {
	return config
}
