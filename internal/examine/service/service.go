package service

import (
	"fmt"
	"log"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/viper"
	c "github.com/switchover/eGovFrameChecker/internal/constant"
	"github.com/switchover/eGovFrameChecker/internal/examine/common"
	"github.com/switchover/eGovFrameChecker/pkg/csv"
	"github.com/switchover/eGovFrameChecker/pkg/java"
	"github.com/switchover/eGovFrameChecker/pkg/parser"
)

func Examine(files []string) (err error) {
	verbose := viper.GetBool("inspect.verbose")
	output := viper.GetBool("inspect.output")
	skipFileError := viper.GetBool("inspect.skip")

	var writer *csv.Writer
	if output {
		writer, err = csv.NewWriter("services.csv",
			[]string{"Total list (*ServiceImpl.java)", "Extends EgovAbstractServiceImpl", "Use Interface", "Super Class"})
		if err != nil {
			return err
		}
		defer writer.Close()
	}

	total := 0
	violations := 0
	for i, f := range files {
		if verbose {
			log.Printf("%d: %s", i+1, f)
		}

		data, err := os.ReadFile(f)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
			if skipFileError {
				continue
			}
			os.Exit(1)
		}

		input := antlr.NewInputStream(string(data))

		lexer := parser.NewJavaLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewJavaParser(stream)

		listener := &java.Listener{}
		antlr.ParseTreeWalkerDefault.Walk(listener, p.CompilationUnit())

		classResult := common.CheckClassAnnotations("service", listener)
		extendsResult, superClass := common.CheckSuperClass("service", listener)
		implementsResult := common.CheckImplementation("service", listener)

		if listener.IsInterface {
			log.Printf("%s- Service(%s) excluded because it's an interface.%s\n",
				c.Yellow, listener.ClassName, c.Reset)
			continue
		}

		total++
		target := common.FormatClassName(listener.ClassName, f)
		record := []string{target}
		service := target
		implement := target
		if !(classResult && extendsResult) && !implementsResult {
			log.Printf("%s- Service(%s%s%s) violates the class and interface rules.%s\n",
				c.Magenta, c.MagentaUnderline, listener.ClassName, c.MagentaNoUnderline, c.Reset)
			violations++
			service = ""
			implement = ""
		} else if !(classResult && extendsResult) {
			log.Printf("%s- Service(%s%s%s) violates the class rule.%s\n",
				c.Magenta, c.MagentaUnderline, listener.ClassName, c.MagentaNoUnderline, c.Reset)
			violations++
			service = ""
		} else if !implementsResult {
			log.Printf("%s- Service(%s%s%s) violates the interface rule.%s\n",
				c.Magenta, c.MagentaUnderline, listener.ClassName, c.MagentaNoUnderline, c.Reset)
			violations++
			implement = ""
		}
		record = append(record, service, implement, superClass)

		if writer != nil {
			err = writer.Write(record)
			if err != nil {
				if skipFileError {
					log.Printf("Failed to write record but skipped: %v\n", err)
					continue
				}
				return err
			}
		}
	}

	log.Println("--------------------------------------------------------------------------------")
	if total == 0 {
		log.Printf("%s No service found.\n", c.IconNotOkay)
		log.Println("--------------------------------------------------------------------------------")
		return
	}

	if violations == 0 {
		if total > 1 {
			log.Printf("%s Services(%d files) are OK.\n", c.IconOkay, total)
		} else {
			log.Printf("%s Service(1 file) is OK.\n", c.IconOkay)
		}
	} else {
		if total > 1 {
			log.Printf("%s Services(%d files) have %d violation(s).\n", c.IconNotOkay, total, violations)
		} else {
			log.Printf("%s Service(1 file) has %d violation.\n", c.IconNotOkay, violations)
		}
	}
	log.Println("--------------------------------------------------------------------------------")

	if output {
		log.Println("Output file written: services.csv")
	}
	return
}
