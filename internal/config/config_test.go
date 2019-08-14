package config

import (
	"os"
	"testing"
)

func setupEnvVars(t *testing.T) {
	err := os.Setenv("DB_PASSWORD", "secret")
	if err != nil {
		t.Errorf("Could not set env vars")
	}

	err = os.Setenv("DB_NAME", "boilerplate")
	if err != nil {
		t.Errorf("Could not set env vars")
	}
}

func TestGetConfig(t *testing.T) {
	setupEnvVars(t)

	c := GetConfig()

	if c.DbPassword != "secret" {
		t.Errorf("Config c.%v = %v; want %v", "DbPassword", c.DbPassword, "secret")
	}
	if c.DbSslMode != "enable" {
		t.Errorf("Config c.%v = %v; want %v", "DbSslMode", c.DbSslMode, "enable")
	}
	if c.DbName != "boilerplate" {
		t.Errorf("Config c.%v = %v; want %v", "DbName", c.DbName, "boilerplate")
	}
	if c.DbUser != "postgres" {
		t.Errorf("Config c.%v = %v; want %v", "DbUser", c.DbUser, "postgres")
	}
	if c.DbPort != 5432 {
		t.Errorf("Config c.%v = %v; want %v", "DbPort", c.DbPort, 5432)
	}
	if c.DbHost != "localhost" {
		t.Errorf("Config c.%v = %v; want %v", "DbHost", c.DbHost, "localhost")
	}
	if c.LogLevel != "info" {
		t.Errorf("Config c.%v = %v; want %v", "LogLevel", c.LogLevel, "info")
	}
	if c.Port != 3000 {
		t.Errorf("Config c.%v = %v; want %v", "Port", c.Port, 3000)
	}
}
