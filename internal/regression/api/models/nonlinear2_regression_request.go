package models

import "RegressionHelperAPI/internal/regression/logic/models"

type NonLinear2RegressionRequest struct {
	Pairs []models.Point `json:"pairs"`
}
