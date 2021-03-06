package inputdockerlog

import (
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/tsaikd/gogstash/config"
)

var (
	logger = config.Logger
)

func init() {
	logger.Level = logrus.DebugLevel
	config.RegistInputHandler(ModuleName, InitHandler)
}

func Test_main(t *testing.T) {
	require := require.New(t)
	require.NotNil(require)

	conf, err := config.LoadFromString(`{
		"input": [{
			"type": "dockerlog",
			"dockerurl": "unix:///var/run/docker.sock",
			"sincepath": "sincedb-%{USER}"
		}]
	}`)
	require.NoError(err)

	err = conf.RunInputs()
	require.NoError(err)

	waitsec := 10
	logger.Debugf("Wait for %d seconds", waitsec)
	time.Sleep(time.Duration(waitsec) * time.Second)
}
