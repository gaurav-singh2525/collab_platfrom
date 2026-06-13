package services

import (
	"fmt"

	"collab-code-platform/internal/execution"
	"collab-code-platform/internal/models"
)

type ExecutionService struct {
}

func NewExecutionService() *ExecutionService {
	return &ExecutionService{}
}

func (s *ExecutionService) Execute(
	language string,
	code string,
) (*models.ExecutionResult, error) {

	switch language {

	case "python":
		return execution.ExecutePython(code)

	case "javascript":
		return execution.ExecuteJavaScript(code)

	case "cpp":
		return execution.ExecuteCpp(code)

	default:
		return nil,
			fmt.Errorf(
				"unsupported language",
			)
	}
}
