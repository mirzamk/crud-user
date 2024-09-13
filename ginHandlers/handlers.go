package ginHandlers

import (
	"net/http"
	"tes-rssa/database"
	"tes-rssa/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllUser returs list of all users from the database
// @Summary return list of all
// @Description returs list of all users from the database
// @Tags Users
// @Success 200 {object} models.User
// @Router /api/v1/user [get]
func GetAllUser(c *gin.Context) {
	var users []models.User
	result := database.DB.Debug().Raw(`SELECT * FROM "user"`).Find(&users)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, nil)
		} else {
			c.JSON(http.StatusInternalServerError, nil)
		}
		return
	}
	c.JSON(http.StatusOK, users)
}
