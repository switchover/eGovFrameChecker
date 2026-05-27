package jpa

import (
	"github.com/switchover/eGovFrameChecker/internal/examine/common"
	"github.com/switchover/eGovFrameChecker/pkg/java"
)

func Examine(listener *java.Listener) (result bool, isRepository bool) {
	result, _ = common.CheckExtendsInterface("repository.jpa", listener)
	if !result {
		return
	}
	isRepository = true

	result, _ = common.CheckClassAnnotations("repository.jpa", listener)
	return
}
