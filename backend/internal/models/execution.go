package models

type ExecutionResult struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}