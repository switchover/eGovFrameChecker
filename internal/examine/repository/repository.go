package repository

import (
	"log"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/viper"
	c "github.com/switchover/eGovFrameChecker/internal/constant"
	"github.com/switchover/eGovFrameChecker/internal/examine/common"
	"github.com/switchover/eGovFrameChecker/internal/examine/repository/hibernate"
	"github.com/switchover/eGovFrameChecker/internal/examine/repository/ibatis"
	"github.com/switchover/eGovFrameChecker/internal/examine/repository/jpa"
	"github.com/switchover/eGovFrameChecker/internal/examine/repository/mapper"
	"github.com/switchover/eGovFrameChecker/internal/examine/repository/mybatis"
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
		writer, err = csv.NewWriter("repositories.csv",
			[]string{"Total list (*DAO.java or *Mapper.java)", "Extends EgovAbstract* or Use @Mapper", "Super Class"})
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

		superClassName := ""
		result, listener, superClassName, err := check(f)
		if err != nil {
			if skipFileError {
				log.Printf("Failed to examine file but skipped: %v\n", err)
				continue
			}
			return err
		}

		total++
		target := common.FormatClassName(listener.ClassName, f)
		record := []string{target}
		criteria := target
		if !result {
			log.Printf("%s- Repository(%s%s%s) violates the repository rule.%s\n",
				c.Magenta, c.MagentaUnderline, listener.ClassName, c.MagentaNoUnderline, c.Reset)
			violations++
			criteria = ""
		}
		record = append(record, criteria, superClassName)

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
		log.Printf("%s No repository found.\n", c.IconNotOkay)
		log.Println("--------------------------------------------------------------------------------")
		return
	}

	if violations == 0 {
		if total > 1 {
			log.Printf("%s Repositories(%d files) are OK.\n", c.IconOkay, total)
		} else {
			log.Printf("%s Repository(1 file) is OK.\n", c.IconOkay)
		}
	} else {
		if total > 1 {
			log.Printf("%s Repositories(%d files) have %d violation(s).\n", c.IconNotOkay, total, violations)
		} else {
			log.Printf("%s Repository(1 file) has %d violation.\n", c.IconNotOkay, violations)
		}
	}
	log.Println("--------------------------------------------------------------------------------")

	if output {
		log.Println("Output file written: repositories.csv")
	}
	return
}

func check(f string) (result bool, listener *java.Listener, superClassName string, err error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return
	}

	input := antlr.NewInputStream(string(data))

	lexer := parser.NewJavaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewJavaParser(stream)

	listener = &java.Listener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.CompilationUnit())

	result, superClassName = ibatis.Examine(listener)
	if result {
		return
	}

	result, superClassName = mybatis.Examine(listener)
	if result {
		return
	}

	result = mapper.Examine(listener)
	if result {
		superClassName = "<@Mapper>"
		return
	}

	result = jpa.Examine(listener)
	if result {
		superClassName = "<JPA>"
		return
	}

	result = hibernate.Examine(listener)
	return
}
