package mailer

import (
	"fmt"
	"log"
	"time"
)

func sendMail(fn func() (int, error)) (int, error) {
	for i := 0; i < maxRetries; i++ {
		status, err := fn()
		if err != nil {
			log.Printf("Attempt %d failed: %v", i+1, err)
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		return status, nil
	}
	return -1, fmt.Errorf("Failed after %d attempts", maxRetries)
}
