package models

type Execution struct {
	Id            string            `bson:"_id" json:"id"`
	Flow          Workflow          `json:"flow"`
	CurrentStepId int               `json:"current_step_id"`
	Context       map[string]string `json:"context"`
}
