package handlers

import (
	"RegressionHelperAPI/internal/regression/api/models"
	"RegressionHelperAPI/internal/regression/logic"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

var NonLinear2RegressionHandler fiber.Handler = func(c *fiber.Ctx) error {
	var (
		req  models.NonLinear2RegressionRequest
		resp models.NonLinear2RegressionResponce
	)
	if string(c.Request().Header.ContentType()) != "application/json" {
		resp.Err = fiber.ErrUnsupportedMediaType.Message
		byteResp, _ := json.Marshal(resp)
		c.Write(byteResp)
		return nil
	}

	unmarshErr := json.Unmarshal(c.Body(), &req)

	if unmarshErr != nil {
		resp.Err = unmarshErr.Error()
		byteResp, _ := json.Marshal(resp)
		c.Write(byteResp)
		return nil
	}

	var logic logic.RegressionLogicProvider

	consts := logic.NonLinear2Regression(req.Pairs)

	resp.A = consts[0]
	resp.B = consts[1]
	resp.C = consts[2]

	byteResp, _ := json.Marshal(resp)
	c.Write(byteResp)
	return nil
}
