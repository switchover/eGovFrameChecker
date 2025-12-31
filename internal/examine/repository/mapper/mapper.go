package mapper

import (
	"strings"

	"github.com/switchover/eGovFrameChecker/internal/examine/common"
	"github.com/switchover/eGovFrameChecker/pkg/java"
)

func Examine(listener *java.Listener) (result bool) {
	result, annotation := common.CheckClassAnnotations("repository.mapper", listener)
	if !result {
		return
	}

	if strings.HasPrefix(annotation, "@") {
		annotation = annotation[1:]
	}
	result = common.CheckConditionalImports("repository.mapper", annotation, listener)
	if !result {
		return
	}

	result = common.CheckInterface("repository.mapper", listener)
	return
}
