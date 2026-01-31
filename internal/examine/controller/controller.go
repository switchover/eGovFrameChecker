package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/viper"
	c "github.com/switchover/eGovFrameChecker/internal/constant"
	"github.com/switchover/eGovFrameChecker/internal/examine/common"
	"github.com/switchover/eGovFrameChecker/internal/i18n"
	"github.com/switchover/eGovFrameChecker/internal/json"
	"github.com/switchover/eGovFrameChecker/pkg/csv"
	"github.com/switchover/eGovFrameChecker/pkg/java"
	"github.com/switchover/eGovFrameChecker/pkg/locale"
	"github.com/switchover/eGovFrameChecker/pkg/parser"
)

func Examine(files []string, streamer *json.Streamer) (err error) {
	verbose := viper.GetBool("inspect.verbose")
	output := viper.GetBool("inspect.output")
	skipFileError := viper.GetBool("inspect.skip")

	var writer *csv.Writer
	if output {
		writer, err = csv.NewWriter("controllers.csv",
			[]string{"Total list (*Controller.java)", "Use @Controller/RestController", "Use Annotation HandlerMapping"})
		if err != nil {
			return err
		}
		defer writer.Close()
	}

	total := 0
	violations := 0
	logList := make([]string, 0, 10)
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

		classResult, _ := common.CheckClassAnnotations("controller", listener)
		methodResult := common.CheckMethodAnnotations("controller", listener)

		total++
		target := common.FormatClassName(listener.ClassName, f)
		record := []string{target}
		controller := target
		requestMapping := target
		if !classResult && !methodResult {
			logList = append(logList, fmt.Sprintf("%s- Controller(%s%s%s) violates the class and method rules.%s\n",
				c.Magenta, c.MagentaUnderline, listener.ClassName, c.MagentaNoUnderline, c.Reset))
			violations++
			controller = ""
			requestMapping = ""
			if streamer != nil {
				message, err := i18n.GetErrorMessage("CTR001", locale.GetLanguage())
				if err != nil {
					return err
				}
				err = streamer.AddViolation(json.Controller, json.Violation{
					FilePath:    f,
					PackageName: listener.PackageName,
					ClassName:   listener.ClassName,
					Violation:   message.Message,
					Description: message.Description,
				})
				if err != nil {
					return err
				}
			}
		} else if !classResult {
			logList = append(logList, fmt.Sprintf("%s- Controller(%s%s%s) violates the class rule.%s\n",
				c.Magenta, c.MagentaUnderline, listener.ClassName, c.MagentaNoUnderline, c.Reset))
			violations++
			controller = ""
			if streamer != nil {
				message, err := i18n.GetErrorMessage("CTR002", locale.GetLanguage())
				if err != nil {
					return err
				}
				err = streamer.AddViolation(json.Controller, json.Violation{
					FilePath:    f,
					PackageName: listener.PackageName,
					ClassName:   listener.ClassName,
					Violation:   message.Message,
					Description: message.Description,
				})
				if err != nil {
					return err
				}
			}
		} else if !methodResult {
			logList = append(logList, fmt.Sprintf("%s- Controller(%s%s%s) violates the method rule.%s\n",
				c.Magenta, c.MagentaUnderline, listener.ClassName, c.MagentaNoUnderline, c.Reset))
			violations++
			requestMapping = ""
			if streamer != nil {
				message, err := i18n.GetErrorMessage("CTR003", locale.GetLanguage())
				if err != nil {
					return err
				}
				err = streamer.AddViolation(json.Controller, json.Violation{
					FilePath:    f,
					PackageName: listener.PackageName,
					ClassName:   listener.ClassName,
					Violation:   message.Message,
					Description: message.Description,
				})
				if err != nil {
					return err
				}
			}
		}
		record = append(record, controller, requestMapping)

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
		log.Printf("%s No controller found.\n", c.IconNotOkay)
		log.Println("--------------------------------------------------------------------------------")
		return
	}

	if violations == 0 {
		if total > 1 {
			log.Printf("%s Controllers(%d files) are OK.\n", c.IconOkay, total)
		} else {
			log.Printf("%s Controller(1 file) is OK.\n", c.IconOkay)
		}
	} else {
		if total > 1 {
			log.Printf("%s Controllers(%d files) have %d violation(s).\n", c.IconNotOkay, total, violations)
		} else {
			log.Printf("%s Controller(1 file) has %d violation.\n", c.IconNotOkay, violations)
		}
	}
	for _, message := range logList {
		log.Printf("%s", message)
	}
	log.Println("--------------------------------------------------------------------------------")

	if output {
		log.Println("Output file written: controllers.csv")
	}
	return
}
