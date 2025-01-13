package target

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gobwas/glob"
	"github.com/spf13/viper"
)

type fullQualifiedClassName string

var sourceFiles map[fullQualifiedClassName]string

var controllerFiles []string
var serviceFiles []string
var repositoryFiles []string

func init() {
	sourceFiles = make(map[fullQualifiedClassName]string)
}

func GetSourceFile(fqcn string) string {
	return sourceFiles[fullQualifiedClassName(fqcn)]
}

func GatherSourceFiles(target string, packages []string) (totalFiles int, err error) {
	err = filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".java") {
			for _, pkg := range packages {
				toPackage := convertFilePathToPackageStyle(path)
				index := strings.Index(toPackage, strings.TrimSpace(pkg))
				if index >= 0 {
					checkController(path)
					checkService(path)
					checkRepository(path)

					fqcn := fullQualifiedClassName(toPackage[index:])
					sourceFiles[fqcn] = path
					totalFiles++
					break // 패키지 중 처음 매칭되는 것만 처리
				}
			}
		}
		return nil
	})
	return
}

func checkController(path string) {
	pattern := viper.GetString("controller.fileNameGlobPattern")
	g := glob.MustCompile("**/" + pattern + ".java")
	path = strings.ReplaceAll(path, "\\", "/") // Windows OS 처리
	if g.Match(path) {
		controllerFiles = append(controllerFiles, path)
	}
}

func checkService(path string) {
	pattern := viper.GetString("service.fileNameGlobPattern")
	g := glob.MustCompile("**/" + pattern + ".java")
	path = strings.ReplaceAll(path, "\\", "/") // Windows OS 처리
	if g.Match(path) {
		serviceFiles = append(serviceFiles, path)
	}
}

func checkRepository(path string) {
	pattern := viper.GetString("repository.fileNameGlobPattern")
	g := glob.MustCompile("**/" + pattern + ".java")
	path = strings.ReplaceAll(path, "\\", "/") // Windows OS 처리
	if g.Match(path) {
		repositoryFiles = append(repositoryFiles, path)
	}
}

func GetControllerFiles() []string {
	return controllerFiles
}

func GetServiceFiles() []string {
	return serviceFiles
}

func GetRepositoryFiles() []string {
	return repositoryFiles
}

func convertFilePathToPackageStyle(filePath string) string {
	// Remove target directory
	target := viper.GetString("inspect.target") + string(filepath.Separator)
	if strings.HasPrefix(filePath, target) {
		filePath = filePath[len(target):]
	}
	// Remove file extension
	filePath = strings.TrimSuffix(filePath, filepath.Ext(filePath))
	// Replace file separators with dots
	packagePath := strings.ReplaceAll(filePath, string(filepath.Separator), ".")
	return packagePath
}
