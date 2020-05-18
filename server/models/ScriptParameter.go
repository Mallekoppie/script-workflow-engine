package models

type ScriptParameter struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	DataType    string `json:"data_type"`
}
