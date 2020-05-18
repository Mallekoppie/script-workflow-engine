package models

type Workflow struct {
	Id          string         `bson:"_id" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Steps       []WorkflowStep `json:"steps"`
}
