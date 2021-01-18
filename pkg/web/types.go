package web

type Any = interface{}

type Filler = func(Any) error

type FillerError struct {
	FillerError error
}

func (ferr FillerError) Error() string {
	return ferr.FillerError.Error()
}

type ErrorResult struct {
	StatusCode       int
	StatusPhrase     string
	DeveloperMessage string
	ErrorString      string
	TraceId          string
	ParentSpanId     string
	SpanId           string
	Cause            map[string]string
	Args             []Any
	Details          []ErrorDetail
}

type ErrorDetail struct {
	error string
	Args  []string
	Path  string
}

type RequestContext struct {
}
