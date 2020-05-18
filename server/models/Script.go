package models

type Script struct {
	Id               string            `bson:"_id" json:"id"`
	Location         string            `json:"location"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	InputParameters  []ScriptParameter `json:"input_parameters"`
	OutputParameters []ScriptParameter `json:"output_parameters"`
}
