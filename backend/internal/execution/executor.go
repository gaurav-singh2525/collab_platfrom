package execution

import (
	"os"
	"os/exec"

	"collab-code-platform/internal/models"
	"context"
	"strings"
	"time"
)



func ExecutePython(
	code string, input string,
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

	if err := file.Close(); err != nil {
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
		"-i",
		"--rm",
		"--network", "none",
		"--read-only",
		"--tmpfs", "/tmp",
		"--cpus", "1",
		"--memory", "128m",
		"--pids-limit", "32",
		"-v", file.Name()+":/code/main.py:ro",
		"python:3.11",
		"python",
		"/code/main.py",
	)
	cmd.Stdin =
		strings.NewReader(input)

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

func ExecuteJavaScript(
	code string, input string,
) (
	*models.ExecutionResult,
	error,
) {
	file, err := os.CreateTemp(
		"",
		"code-*.js",
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

	if err := file.Close(); err != nil {
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
		"-i",
		"--rm",
		"--network", "none",
		"--read-only",
		"--tmpfs", "/tmp",
		"--cpus", "1",
		"--memory", "128m",
		"--pids-limit", "32",
		"-v", file.Name()+":/code/main.js:ro",
		"node:22-alpine",
		"node",
		"/code/main.js",
	)

	cmd.Stdin =
		strings.NewReader(input)

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

func ExecuteCpp(
	code string, input string,
) (
	*models.ExecutionResult,
	error,
) {
	file, err := os.CreateTemp(
		"",
		"code-*.cpp",
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

	if err := file.Close(); err != nil {
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
		"-i",
		"--rm",
		"--network", "none",
		"--read-only",
		"--tmpfs", "/tmp",
		"--cpus", "1",
		"--memory", "256m",
		"--pids-limit", "32",
		"-v", file.Name()+":/code/main.cpp:ro",
		"gcc:15-bookworm",
		"sh",
		"-c",
		"g++ /code/main.cpp -o /tmp/main && /tmp/main",
	)

	cmd.Stdin =
		strings.NewReader(input)

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
