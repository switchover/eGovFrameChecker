package hibernate

import (
	"github.com/switchover/eGovFrameChecker/internal/examine/common"
	"github.com/switchover/eGovFrameChecker/pkg/java"
)

func Examine(listener *java.Listener) (result bool) {
	result = common.CheckClassAnnotations("repository.hibernate", listener)
	if !result {
		return
	}

	result = common.CheckFieldTypes("repository.hibernate", listener)
	return
}
