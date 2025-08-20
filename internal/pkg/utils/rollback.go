package utils

import (
	"fmt"
	"images-service/internal/app/ent"
	"log"
)

func Rollback(tx *ent.Tx, originalErr error) error {
	if rbErr := tx.Rollback(); rbErr != nil {
		log.Printf("rollback failed: %v", rbErr)
		return fmt.Errorf("%v | rollback failed: %v", originalErr, rbErr)
	}
	return originalErr
}
