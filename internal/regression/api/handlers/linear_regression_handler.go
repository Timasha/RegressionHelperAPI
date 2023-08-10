package handlers

import (
	"RegressionHelperAPI/internal/regression/api/models"
	"RegressionHelperAPI/internal/regression/logic"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

var LinearRegressionHandler fiber.Handler = func(c *fiber.Ctx) error {
	var (
		req  models.LinearRegressionRequest
		resp models.LinearRegressionResponce
	)
	if string(c.Request().Header.ContentType()) != "application/json" {
		resp.Err = fiber.ErrUnsupportedMediaType.Message
		byteResp, _ := json.Marshal(resp)
		c.Write(byteResp)
		return nil
	}

	body := c.Body()

	unmarshErr := json.Unmarshal(body, &req)
	if unmarshErr != nil {
		resp.Err = unmarshErr.Error()
		byteResp, _ := json.Marshal(resp)
		c.Write(byteResp)
		return nil
	}
	var logic logic.RegressionLogicProvider

	consts := logic.LinearRegression(req.Pairs)
	resp.A = consts[0]
	resp.B = consts[1]
	byteResp, _ := json.Marshal(resp)

	c.Write(byteResp)
	return nil
}
