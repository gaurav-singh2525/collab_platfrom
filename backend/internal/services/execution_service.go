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
	input string,
) (*models.ExecutionResult, error) {

	switch language {

	case "python":
		return execution.ExecutePython(code, input)

	case "javascript":
		return execution.ExecuteJavaScript(code, input)

	case "cpp":
		return execution.ExecuteCpp(code, input)

	default:
		return nil,
			fmt.Errorf(
				"unsupported language",
			)
	}
}
