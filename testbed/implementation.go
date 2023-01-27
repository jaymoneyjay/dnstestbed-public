package testbed

type Implementation interface {
	flushCache() error
	reload() error
	start()
	filterQueries(queryLog []byte) []byte
	SetConfig(qmin, reload bool)
}
