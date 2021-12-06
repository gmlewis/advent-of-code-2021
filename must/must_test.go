package must

var fakeFatalErr error

func fakeFatal(v ...interface{}) {
	if err, ok := v[0].(error); ok {
		fakeFatalErr = err
	}
}
