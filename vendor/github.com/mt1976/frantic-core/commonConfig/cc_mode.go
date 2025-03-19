package commonConfig

type MODE struct {
	name string
}

var (
	MODE_DEVELOPMENT MODE = MODE{name: "development"}
	MODE_PRODUCTION  MODE = MODE{name: "production"}
	MODE_TEST        MODE = MODE{name: "test"}
)
