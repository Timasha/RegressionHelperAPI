package logic_test

import (
	"RegressionHelperAPI/internal/regression/logic"
	"RegressionHelperAPI/internal/regression/logic/models"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearRegression(t *testing.T) {
	testTable := []struct {
		Pairs          []models.Point
		ExceptedConsts [2]float64
	}{
		{Pairs: []models.Point{models.Point{1, 2}, models.Point{2, 4.1}, models.Point{3, 7}, models.Point{4, 17}, models.Point{5, 22}}, ExceptedConsts: [2]float64{5.29, -5.45}},
		{Pairs: []models.Point{models.Point{45.1, 68.8}, models.Point{59, 61.2}, models.Point{57.2, 59.9}, models.Point{61.8, 56.7}, models.Point{58.8, 55}, models.Point{47.2, 54.3}, models.Point{55.2, 49.3}}, ExceptedConsts: [2]float64{-0.35, 76.88}},
	}

	var logic logic.RegressionLogicProvider
	for i := 0; i < len(testTable); i++ {
		result := logic.LinearRegression(testTable[i].Pairs)
		assert.LessOrEqual(t, math.Abs(result[0]-testTable[i].ExceptedConsts[0]), 0.01)
		assert.LessOrEqual(t, math.Abs(result[1]-testTable[i].ExceptedConsts[1]), 0.01)
	}
}

func TestNonLinear2Regression(t *testing.T) {
	testTable := []struct {
		Pairs          []models.Point
		ExceptedConsts [3]float64
	}{
		{Pairs: []models.Point{models.Point{1, 2}, models.Point{2, 4.1}, models.Point{3, 7}, models.Point{4, 17}, models.Point{5, 22}}, ExceptedConsts: [3]float64{1429.61, -1224.76142857143, 205.008571428572}},
	}

	var logic logic.RegressionLogicProvider

	for i := 0; i < len(testTable); i++ {
		result := logic.NonLinear2Regression(testTable[i].Pairs)
		assert.LessOrEqual(t, math.Abs(result[0]-testTable[i].ExceptedConsts[0]), 0.01)
		assert.LessOrEqual(t, math.Abs(result[1]-testTable[i].ExceptedConsts[1]), 0.01)
		assert.LessOrEqual(t, math.Abs(result[2]-testTable[i].ExceptedConsts[2]), 0.01)
	}
}
