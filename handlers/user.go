package handlers

import (
	"errors"
	"mentalartsapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateUser godoc
// @Summary Create User Object
// @Description Create User
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 400
// @Failure 500
// @Router /user/ [post]
func (h Handler) CreateUser(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user.ID = 0

	if err := h.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = h.db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUsers godoc
// @Summary Get all Users array
// @Description Get User
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /user/ [get]
func (h Handler) GetUsers(c *gin.Context) {
	var users []models.User

	if err := h.db.Preload("Books").Preload("Books.User").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary Get User Object
// @Description Get User
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /user/ [get]
func (h Handler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	var user models.User

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Update a user by ID
// @Description Update a user by ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "Updated user object"
// @Success 200 {object} models.User
// @Router /user/{id} [put]
func (h Handler) UpdateUser(c *gin.Context) {
	var user models.User
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if id != user.ID {
		c.JSON(http.StatusBadRequest, "Inconsintency between Id param and User.ID")
	}

	if err := h.db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}

// DeleteUser godoc
// @Summary      DeleteUser an user
// @Description  delete by id user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param id path int true "User ID"
// @Success      200  {object}  models.User
// @Failure 400
// @Failure 500
// @Router       /user/{id} [delete]
func (h Handler) DeleteUserById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID must be Integer"})
		return
	}

	if err := h.db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)

	}
}
