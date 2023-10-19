package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hitoshi-w/go-lang/bootstrap"
	"github.com/hitoshi-w/go-lang/model"
	"github.com/hitoshi-w/go-lang/usecase"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase usecase.SignupUsecase
	Env *bootstrap.Env
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request model.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	_, err = sc.SignupUsecase.GetByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, "")
		return
	}

	ecrypedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	request.Password = string(ecrypedPassword)

	user := model.User{
		Email: request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	signupResponse := model.SignupResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
