package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	h "user-service/src/helper"
	"user-service/src/v1/model"
	"user-service/src/v1/usecase"
)

func SignUp(c *gin.Context) {
	var (
		user model.User
		err  error
	)

	v := validator.New()
	err = c.ShouldBindJSON(&user)

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Failed when parsing request"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	if err = v.Struct(user); err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Validation error"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	err = usecase.Registration(&user)

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Failed to register user"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}

func SignIn(c *gin.Context) {
	var (
		userLogin model.UserLogin
		err       error
	)

	v := validator.New()
	err = c.ShouldBindJSON(&userLogin)

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Failed when parsing request"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	if err = v.Struct(userLogin); err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Validation error"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	err = usecase.Login(userLogin)

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Failed to login"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}
