package support

const Version string = "v1.12.7"

const (
	EnvRuntime = "runtime"
	EnvArtisan = "artisan"
	EnvTest    = "test"
)

var (
	Env          = EnvRuntime
	EnvPath      string
	RelativePath string
	RootPath     string
)
