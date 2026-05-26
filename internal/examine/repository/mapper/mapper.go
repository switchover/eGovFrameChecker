package mapper

import (
	"fmt"
	"strings"

	"github.com/switchover/eGovFrameChecker/internal/examine/common"
	"github.com/switchover/eGovFrameChecker/pkg/java"
)

func Examine(listener *java.Listener) (result bool, superClassName string) {
	result, annotation := common.CheckClassAnnotations("repository.mapper", listener)
	if !result {
		return
	}

	superClassName = fmt.Sprintf("<%s>", annotation)

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
