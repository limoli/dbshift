package lib

import (
	"os"
)

type IDbConfig interface {
	GetUser() string
	GetPassword() string
	GetHost() string
	GetPort() string
	GetName() string
}

type Configuration struct {
	Db Db `yaml:"db"`
}

type Db struct {
	Type       DatabaseType  `yaml:"type"`
	Migration  MigrationPath `yaml:"migration"`
	Connection Connection    `yaml:"connection"`
}

type MigrationPath struct {
	Path    string `yaml:"path"`
	PathEnv string `yaml:"pathEnv"`
}

type Connection struct {
	Name     ConnectionType `yaml:"name"`
	User     ConnectionType `yaml:"user"`
	Password ConnectionType `yaml:"password"`
	Host     ConnectionType `yaml:"host"`
	Port     ConnectionType `yaml:"port"`
}

type ConnectionType struct {
	Env   string `yaml:"env"`
	Value string `yaml:"value"`
}

func (c *Connection) GetName() string {
	value := os.Getenv(c.Name.Env)
	if value != "" {
		return value
	}
	return c.Name.Value
}

func (c *Connection) GetUser() string {
	value := os.Getenv(c.User.Env)
	if value != "" {
		return value
	}
	return c.User.Value
}

func (c *Connection) GetPassword() string {
	value := os.Getenv(c.Password.Env)
	if value != "" {
		return value
	}
	return c.Password.Value
}

func (c *Connection) GetHost() string {
	value := os.Getenv(c.Host.Env)
	if value != "" {
		return value
	}
	return c.Host.Value
}

func (c *Connection) GetPort() string {
	value := os.Getenv(c.Port.Env)
	if value != "" {
		return value
	}
	return c.Port.Value
}
