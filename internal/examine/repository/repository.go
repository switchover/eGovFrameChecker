package repository

import "log"

func Examine(files []string) (err error) {
	for _, f := range files {
		log.Println("Repository:", f)
	}

	return
}
