package execution

import (
	"os"
	"os/exec"

	"collab-code-platform/internal/models"
	"context"
	"time"
)

func ExecutePython(
	code string,
) (*models.ExecutionResult, error) {

	file, err := os.CreateTemp(
		"",
		"code-*.py",
	)

	if err != nil {
		return nil, err
	}

	defer os.Remove(
		file.Name(),
	)

	_, err = file.WriteString(
		code,
	)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"docker",
		"run",
		"--rm",
		"-v",
		file.Name()+":/code/main.py",
		"python:3.11",
		"python",
		"/code/main.py",
	)

	output, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return &models.ExecutionResult{
			Stdout: "",
			Stderr: "Time Limit Exceeded",
		}, nil
	}

	if err != nil {
		return &models.ExecutionResult{
			Stdout: "",
			Stderr: string(output),
		}, nil
	}

	return &models.ExecutionResult{
		Stdout: string(output),
		Stderr: "",
	}, nil

}
