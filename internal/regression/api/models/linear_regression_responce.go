package models

type LinearRegressionResponce struct {
	A   float64 `json:"a"`
	B   float64 `json:"b"`
	Err string  `json:"error"`
}
