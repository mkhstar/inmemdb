package parser

type requestBody struct {
	OutputFormat string   `json:"outputFormat"`
	Command      []string `json:"command"`
}
