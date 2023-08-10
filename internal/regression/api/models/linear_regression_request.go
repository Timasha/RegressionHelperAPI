package models

import "RegressionHelperAPI/internal/regression/logic/models"

type LinearRegressionRequest struct {
	Pairs []models.Point `json:"pairs"`
}
