package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/viper"
	"github.com/switchover/eGovFrameChecker/internal/target"
	"github.com/switchover/eGovFrameChecker/pkg/java"
	"github.com/switchover/eGovFrameChecker/pkg/parser"
)

var cache = make(map[string]*java.Listener)

func CheckClassAnnotations(section string, listener *java.Listener) bool {
	annotations := viper.GetString(fmt.Sprintf("%s.%s", section, "classAnnotations"))
	for _, classAnnotation := range listener.ClassAnnotations {
		for _, annotation := range strings.Split(annotations, ",") {
			if classAnnotation == strings.TrimSpace(annotation) {
				return true
			}
		}
	}

	return false
}

func CheckMethodAnnotations(section string, listener *java.Listener) bool {
	annotations := viper.GetString(fmt.Sprintf("%s.%s", section, "methodAnnotations"))
	for _, annotation := range strings.Split(annotations, ",") {
		if listener.MethodAnnotations[strings.TrimSpace(annotation)] == true {
			return true
		}
	}

	return false
}

func CheckInterface(section string, listener *java.Listener) bool {
	isInterface := viper.GetBool(fmt.Sprintf("%s.%s", section, "interface"))
	return isInterface == listener.IsInterface
}

func CheckImplementation(section string, listener *java.Listener) bool {
	isImplemented := viper.GetBool(fmt.Sprintf("%s.%s", section, "implementation"))
	if isImplemented {
		return listener.HasImplementation
	}
	return !listener.HasImplementation
}

func CheckFieldTypes(section string, listener *java.Listener) bool {
	fieldTypes := viper.GetString(fmt.Sprintf("%s.%s", section, "fieldTypes"))
	for _, field := range listener.FieldTypes {
		for _, fieldType := range strings.Split(fieldTypes, ",") {
			if field == strings.TrimSpace(fieldType) {
				return true
			}
		}
	}
	return false
}

func CheckSuperClass(section string, listener *java.Listener) (bool, string) {
	if listener.SuperClassName == "" {
		return false, ""
	}
	superClasses := viper.GetString(fmt.Sprintf("%s.%s", section, "superClasses"))
	for _, superClass := range strings.Split(superClasses, ",") {
		if listener.SuperClassName == strings.TrimSpace(superClass) {
			return true, listener.SuperClassName
		}
	}
	// recursive check
	return recursiveSuperClassCheck(superClasses, listener, listener.SuperClassName)
}

func recursiveSuperClassCheck(superClasses string, currentListener *java.Listener, currentClassName string) (bool, string) {
	fqcn := currentListener.GetFqcnFromImports(currentClassName)
	if fqcn == "" {
		return false, ""
	}

	// get super class's listener
	var listener *java.Listener
	if cache[fqcn] == nil {
		f := target.GetSourceFile(fqcn)
		if f == "" {
			return false, ""
		}
		data, err := os.ReadFile(f)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
			return false, ""
		}

		input := antlr.NewInputStream(string(data))

		lexer := parser.NewJavaLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewJavaParser(stream)

		listener = &java.Listener{}
		antlr.ParseTreeWalkerDefault.Walk(listener, p.CompilationUnit())

		cache[fqcn] = listener
	} else {
		listener = cache[fqcn]
	}

	// check expected super classes
	for _, superClass := range strings.Split(superClasses, ",") {
		if listener.SuperClassName == strings.TrimSpace(superClass) {
			return true, listener.SuperClassName
		}
	}

	// recursive function call
	return recursiveSuperClassCheck(superClasses, listener, listener.SuperClassName)
}

func FormatClassName(className string, filePath string) string {
	target := strings.ReplaceAll(viper.GetString("inspect.target"), "\\", "/") + "/" // Windows OS 처리
	if strings.HasPrefix(filePath, target) {
		filePath = filePath[len(target):]
	}
	return fmt.Sprintf("%s.java - %s", className, filePath)
}
