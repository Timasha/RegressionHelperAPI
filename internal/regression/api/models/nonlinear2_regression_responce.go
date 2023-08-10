package models

type NonLinear2RegressionResponce struct {
	A   float64 `json:"a"`
	B   float64 `json:"b"`
	C   float64 `json:"c"`
	Err string  `json:"error"`
}
