// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type JSONStats struct {
	State      string                 `json:"state"`
	SQL        string                 `json:"sql"`
	QueryStats map[string]interface{} `json:"queryStats"`
	Session    map[string]interface{} `json:"session"`
	JSON       map[string]interface{} `json:"json"`
}

type NewBatch struct {
	Text        string `json:"text"`
	ProjectName string `json:"projectName"`
}

type NewProject struct {
	ID string `json:"id"`
}
