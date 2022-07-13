package config

type Config struct {
	Application Application
	Database    Database
	ThirdParty  ThirdParty
}

type Application struct {
	Name string `yaml:"name,omitempty"`
	Port string `yaml:"port,omitempty"`
}

type Database struct {
	Host string `yaml:"host,omitempty"`
	Port string `yaml:"port,omitempty"`
	User string `yaml:"user,omitempty"`
	Pass string `yaml:"pass,omitempty"`
}

type ThirdParty struct {
	Twitter Twitter `yaml:"twitter,omitempty"`
}

type Twitter struct {
	Host      string `yaml:"host,omitempty"`
	AuthToken string `yaml:"auth-token,omitempty"`
}
