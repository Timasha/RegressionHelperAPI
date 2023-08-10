package logic

import (
	"RegressionHelperAPI/internal/regression/logic/models"
)

type RegressionLogicProvider struct{}

func (r *RegressionLogicProvider) LinearRegression(xys []models.Point) [2]float64 {
	var (
		sumxs  float64
		sumx2s float64
		sumxy  float64
		sumys  float64
	)
	for i := 0; i < len(xys); i++ {
		sumxs += xys[i].X
		sumx2s += xys[i].X * xys[i].X
		sumxy += xys[i].X * xys[i].Y
		sumys += xys[i].Y
	}
	n := float64(len(xys))
	var (
		det  float64 = sumx2s*n - sumxs*sumxs
		det1 float64 = sumxy*n - sumxs*sumys
		det2 float64 = sumx2s*sumys - sumxy*sumxs
	)
	return [2]float64{det1 / det, det2 / det}
}

func (r *RegressionLogicProvider) NonLinear2Regression(xys []models.Point) [3]float64 {
	var (
		sumxs  float64
		sumx2s float64
		sumxy  float64
		sumys  float64
		sumx3s float64
		sumx4s float64
		sumxy2 float64
	)
	for i := 0; i < len(xys); i++ {
		sumxs += xys[i].X
		sumx2s += xys[i].X * xys[i].X
		sumxy += xys[i].X * xys[i].Y
		sumys += xys[i].Y
		sumx3s += xys[i].X * xys[i].X * xys[i].X
		sumx4s += xys[i].X * xys[i].X * xys[i].X * xys[i].X
		sumxy2 += xys[i].X * xys[i].Y * xys[i].Y
	}
	n := float64(len(xys))
	var (
		det  float64 = n*sumx2s*sumx4s + sumx2s*sumx3s*sumxs + sumx2s*sumxs*sumx3s - (sumx2s*sumx2s*sumx2s + sumxs*sumxs*sumx4s + n*sumx3s*sumx3s)
		det1 float64 = sumys*sumx2s*sumx4s + sumxy2*sumx3s*sumxs + sumx2s*sumxy*sumx3s - (sumxy2*sumx2s*sumx2s + sumx4s*sumxs*sumxy + sumys*sumx3s*sumx3s)
		det2 float64 = n*sumxy*sumx4s + sumx2s*sumx3s*sumys + sumx2s*sumxs*sumxy2 - (sumx2s*sumxy*sumx2s + sumx4s*sumys*sumxs + n*sumxy2*sumx3s)
		det3 float64 = n*sumx2s*sumxy2 + sumx2s*sumxy*sumxs + sumys*sumxs*sumx3s - (sumx2s*sumx2s*sumys + sumxy2*sumxs*sumxs + n*sumx3s*sumxy)
	)
	return [3]float64{det1 / det, det2 / det, det3 / det}
}
