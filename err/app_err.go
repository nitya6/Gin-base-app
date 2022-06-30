package err

import (
	"fmt"
)

type AppError struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Type    string `json:"type"`
	Details error  `json:"error"`
}

func (err *AppError) Error() string {

	return fmt.Sprintf("Type: %s  Message: %s  Status: %s Error: %v", err.Type, err.Message, err.Status, err.Details)

}
