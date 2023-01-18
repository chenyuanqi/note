package consts

import "os"

func init() {
	switch os.Getenv("ENV_MODE") {
	case EnvModeProd:
		EnvMode = EnvModeProd
	case EnvModeTest:
		EnvMode = EnvModeTest
	}

	EnvModeIsDev = EnvMode == EnvModeDev
	EnvModeIsTest = EnvMode == EnvModeTest
	EnvModeIsRelease = EnvMode == EnvModeProd
}

const (
	EnvModeProd = "release"
	EnvModeTest = "test"
	EnvModeDev  = "dev"
)

var EnvMode = EnvModeDev
var EnvModeIsDev = true
var EnvModeIsTest bool
var EnvModeIsRelease bool
