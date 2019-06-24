package models

//ExportData is used to support values to export JSON
type ExportData struct {
	Message string `json:"message"`
	Data [] interface{} `json:"data"`
	StatusCode int `json:"statusCode"`
}

//Exports just a ExportData slice
type Exports [] ExportData