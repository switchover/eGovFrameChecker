package jpa

import (
	"github.com/switchover/eGovFrameChecker/internal/examine/common"
	"github.com/switchover/eGovFrameChecker/pkg/java"
)

func Examine(listener *java.Listener) (result bool) {
	result = common.CheckClassAnnotations("repository.jpa", listener)
	if !result {
		return
	}

	result, _ = common.CheckSuperClass("repository.jpa", listener)
	return
}
