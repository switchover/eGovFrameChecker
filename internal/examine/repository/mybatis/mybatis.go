package mybatis

import (
	"github.com/switchover/eGovFrameChecker/internal/examine/common"
	"github.com/switchover/eGovFrameChecker/pkg/java"
)

func Examine(listener *java.Listener) (result bool, superClassName string) {
	result = common.CheckClassAnnotations("repository.mybatis", listener)
	if !result {
		return
	}

	result, superClassName = common.CheckSuperClass("repository.mybatis", listener)
	return
}
