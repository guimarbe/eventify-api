package utils

import "log"

func HandleError(err error) error {
	if err != nil {
		log.Printf("Error: %v", err)
	}
	return err
}