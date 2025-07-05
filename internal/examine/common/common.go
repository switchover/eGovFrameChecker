package common

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/viper"
	"github.com/switchover/eGovFrameChecker/internal/target"
	"github.com/switchover/eGovFrameChecker/pkg/java"
	"github.com/switchover/eGovFrameChecker/pkg/parser"
)

var cache = make(map[string]*java.Listener)

var toBeCheckedSuperClasses = make([]string, 0)

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
		if listener.MethodAnnotations[strings.TrimSpace(annotation)] {
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

func CheckFieldTypes(section string, listener *java.Listener, checkSuperClass bool) bool {
	fieldTypes := viper.GetString(fmt.Sprintf("%s.%s", section, "fieldTypes"))
	for _, field := range listener.FieldTypes {
		for _, fieldType := range strings.Split(fieldTypes, ",") {
			if field == strings.TrimSpace(fieldType) {
				return true
			}
		}
	}
	// recursive check
	if checkSuperClass {
		if listener.SuperClassName != "" {
			check, _ := recursiveFieldTypesCheck(fieldTypes, listener, listener.SuperClassName)
			if check && !slices.Contains(toBeCheckedSuperClasses, listener.SuperClassName) {
				toBeCheckedSuperClasses = append(toBeCheckedSuperClasses, listener.SuperClassName)
			}
			return check
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
	check, _ := recursiveSuperClassCheck(superClasses, listener, listener.SuperClassName)
	if check && !slices.Contains(toBeCheckedSuperClasses, listener.SuperClassName) {
		toBeCheckedSuperClasses = append(toBeCheckedSuperClasses, listener.SuperClassName)
	}
	// return current super class name not recursive super class name
	return check, listener.SuperClassName
}

func recursiveFieldTypesCheck(expectedFieldTypes string, currentListener *java.Listener, superClassName string) (bool, string) {
	listener, err := getListenerOfSuperClass(superClassName, currentListener)
	if err != nil {
		return false, ""
	}

	// check an expected field type
	for _, field := range listener.FieldTypes {
		for _, fieldType := range strings.Split(expectedFieldTypes, ",") {
			if field == strings.TrimSpace(fieldType) {
				return true, field
			}
		}
	}

	// recursive function call
	return recursiveFieldTypesCheck(expectedFieldTypes, listener, listener.SuperClassName)
}

func recursiveSuperClassCheck(expectedSuperClasses string, currentListener *java.Listener, superClassName string) (bool, string) {
	listener, err := getListenerOfSuperClass(superClassName, currentListener)
	if err != nil {
		return false, ""
	}

	// check expected super classes
	for _, superClass := range strings.Split(expectedSuperClasses, ",") {
		if listener.SuperClassName == strings.TrimSpace(superClass) {
			return true, listener.SuperClassName
		}
	}

	// recursive function call
	return recursiveSuperClassCheck(expectedSuperClasses, listener, listener.SuperClassName)
}

func GetToBeCheckedSuperClasses() []string {
	return toBeCheckedSuperClasses
}

func FormatClassName(className string, filePath string) string {
	target := strings.ReplaceAll(viper.GetString("inspect.target"), "\\", "/") + "/" // Windows OS 처리
	filePath = strings.TrimPrefix(filePath, target)
	return fmt.Sprintf("%s.java - %s", className, filePath)
}

func getListenerOfSuperClass(superClassName string, currentListener *java.Listener) (*java.Listener, error) {
	if superClassName == "" {
		return nil, fmt.Errorf("super class name is empty")
	}
	fqcn := currentListener.GetFqcnFromImports(superClassName)
	if fqcn == "" {
		return nil, fmt.Errorf("could not find FQCN for class: %s", superClassName)
	}

	if cache[fqcn] == nil {
		f := target.GetSourceFile(fqcn)
		if f == "" {
			return nil, fmt.Errorf("source file not found for FQCN: %s", fqcn)
		}
		data, err := os.ReadFile(f)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %v", err)
		}

		input := antlr.NewInputStream(string(data))

		lexer := parser.NewJavaLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewJavaParser(stream)

		listener := &java.Listener{}
		antlr.ParseTreeWalkerDefault.Walk(listener, p.CompilationUnit())

		cache[fqcn] = listener
	}

	return cache[fqcn], nil
}
