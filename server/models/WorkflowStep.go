package models

type WorkflowStep struct {
	StepId          int               `json:"step_id"`
	ScriptToExecute Script            `json:"script_to_execute"`
	Completed       bool              `json:"completed"`
	InProcess       bool              `json:"in_process"`
	Success         bool              `json:"success"`
	Message         string            `json:"message"`
	Parameters      map[string]string `json:"parameters"`
	Results         map[string]string `json:"results"`
}
