package mapper

import (
	"github.com/switchover/eGovFrameChecker/internal/examine/common"
	"github.com/switchover/eGovFrameChecker/pkg/java"
)

func Examine(listener *java.Listener) (result bool, err error) {
	result = common.CheckClassAnnotations("repository.mapper", listener)
	if !result {
		return
	}

	result = common.CheckInterface("repository.mapper", listener)
	return true, nil
}
