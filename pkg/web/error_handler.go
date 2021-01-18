package web

import "github.com/sirupsen/logrus"

func DefaultErrorHandler(errorContents Any, ctx RequestContext) ErrorResult {
	logrus.Info("Error caught in web handler:", errorContents)
	errResult := ErrorResult{}
	errResult.DeveloperMessage = "Dummy error handler implementation."
	errResult.StatusCode = 500
	return errResult
}
