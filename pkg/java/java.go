package java

import (
	"strings"

	"github.com/switchover/eGovFrameChecker/pkg/parser"
)

// Inner class and interface are not supported

type Listener struct {
	*parser.BaseJavaParserListener
	ClassName         string
	IsInterface       bool
	ClassAnnotations  []string
	MethodAnnotations map[string]bool
	FieldAnnotations  map[string]bool

	// for private use
	currentClass       string
	currentAnnotations []string
	currentMethod      string
	currentField       string
	isInnerClass       bool
}

func (l *Listener) initialize() {
	l.MethodAnnotations = make(map[string]bool)
	l.FieldAnnotations = make(map[string]bool)
}

func (l *Listener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	if l.currentClass != "" { // inner class
		l.isInnerClass = true
		return
	}
	l.initialize()
	l.ClassName = ctx.Identifier().GetText()
	modifiers := ctx.GetParent().GetChild(0) // 클래스 앞의 modifiers 블록
	for _, child := range modifiers.GetChildren() {
		if annotationCtx, ok := child.(*parser.AnnotationContext); ok {
			if strings.HasPrefix(annotationCtx.GetText(), "@") {
				l.ClassAnnotations = append(l.ClassAnnotations, "@"+annotationCtx.QualifiedName().GetText())
			}
		}
	}
	l.currentClass = l.ClassName
}

func (l *Listener) ExitClassDeclaration(_ *parser.ClassDeclarationContext) {
	if l.isInnerClass {
		l.isInnerClass = false
		return
	}
	l.currentClass = ""
}

func (l *Listener) EnterInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	if l.currentClass != "" { // inner interface
		l.isInnerClass = true
		return
	}
	l.initialize()
	l.ClassName = ctx.Identifier().GetText()
	l.IsInterface = true
}

func (l *Listener) ExitInterfaceDeclaration(_ *parser.InterfaceDeclarationContext) {
	if l.isInnerClass {
		l.isInnerClass = false
		return
	}
	l.currentClass = ""
}

func (l *Listener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	if l.isInnerClass {
		return
	}
	l.currentMethod = ctx.Identifier().GetText()
	addAnnotations(l.MethodAnnotations, l.currentAnnotations...)
	l.currentAnnotations = nil
}

func (l *Listener) ExitMethodDeclaration(_ *parser.MethodDeclarationContext) {
	if l.isInnerClass {
		return
	}
	l.currentMethod = ""
}

func (l *Listener) EnterFieldDeclaration(ctx *parser.FieldDeclarationContext) {
	if l.isInnerClass {
		return
	}
	l.currentField = ctx.VariableDeclarators().GetText()
	addAnnotations(l.FieldAnnotations, l.currentAnnotations...)
	l.currentAnnotations = nil
}

func (l *Listener) ExitFieldDeclaration(_ *parser.FieldDeclarationContext) {
	if l.isInnerClass {
		return
	}
	l.currentField = ""
}

func (l *Listener) EnterAnnotation(ctx *parser.AnnotationContext) {
	if l.isInnerClass {
		return
	}
	annotationName := ctx.QualifiedName().GetText()

	if l.currentClass != "" && l.currentMethod == "" && l.currentField == "" {
		l.currentAnnotations = append(l.currentAnnotations, "@"+annotationName)
	}
}

func addAnnotations(annotations map[string]bool, toBeAdded ...string) {
	for _, a := range toBeAdded {
		annotations[a] = true
	}
}
