package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanGrzegorczyk/challenge_meli/database"
	"github.com/ivanGrzegorczyk/challenge_meli/model"
)

func AddRuleHandler(c *gin.Context) {
	var rule model.Rule
	err := c.BindJSON(&rule)
	if err != nil {
		c.String(400, "Error binding rule: %s", err)
		return
	}

	database.InsertRule(rule)

	c.JSON(200, rule)
}
