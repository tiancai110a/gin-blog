package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/golang/glog"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		glog.Info(err.Key, err.Message)
	}

	return
}