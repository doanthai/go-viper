package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig("test")
	appName := viper.Get("app.name")

	assert.Equal(t, appName, "GO-Viper")
}

func TestLoadNumberConfig(t *testing.T)  {
	LoadConfig("test")
	appPort := viper.GetInt("server.port")

	assert.Equal(t, appPort, 8000)
}

func TestEnvFileInCorrect(t *testing.T) {
	isPanic := false
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("This code did not panic")
		} else {
			assert.Equal(t, r, "Env name incorrect")
		}
		isPanic = true
	}()

	LoadConfig("abc")

	assert.True(t, isPanic)
}

func TestNotFoundConfigFile(t *testing.T) {
	isPanic := false
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("This code did not panic")
		} else {
			fmt.Println(r)
		}
		isPanic = true
	}()

	LoadConfig("prod")

	assert.True(t, isPanic)
}