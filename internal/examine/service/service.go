package service

import "log"

func Examine(files []string) (err error) {
	for _, f := range files {
		log.Println("Service:", f)
	}

	return
}
