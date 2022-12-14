package delivery

import (
	"fmt"
	"net/http"
	"sosmed/config"
	"sosmed/features/user/domain"
	"sosmed/utils/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}
	e.GET("/users", handler.MyProfile(), middleware.JWT([]byte(config.JwtKey)))
	e.PUT("/users", handler.UpdateProfile(), middleware.JWT([]byte(config.JwtKey)))
	e.DELETE("/users", handler.Deactivate(), middleware.JWT([]byte(config.JwtKey)))
	e.GET("/users/:email", handler.ShowByEmail())
	e.POST("/register", handler.Register())
	e.POST("/login", handler.Login())
}

func (uh *userHandler) MyProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := common.ExtractToken(c)
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		} else {
			res, err := uh.srv.MyProfile(uint(userID))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
			return c.JSON(http.StatusOK, SuccessResponse("success get my profile", ToResponse(res, "user")))
		}
	}
}

func (uh *userHandler) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := common.ExtractToken(c)
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		} else {
			var input UserFormat
			if err := c.Bind(&input); err != nil {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot bind update data"))
			}

			cnv := ToDomain(input)
			res, err := uh.srv.UpdateProfile(cnv, uint(userID))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
			return c.JSON(http.StatusAccepted, SuccessResponse("Success update user", ToResponse(res, "user")))
		}
	}
}

func (uh *userHandler) Deactivate() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := common.ExtractToken(c)
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		} else {
			res, err := uh.srv.Deactivate(userID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
			return c.JSON(http.StatusAccepted, SuccessResponse("success delete user", ToResponse(res, "")))
		}
	}
}

func (uh *userHandler) ShowByEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.QueryParam("email")
		res, err := uh.srv.ShowByEmail(email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success get user", ToResponse(res, "user")))
	}
}

func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UserFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		cnv := ToDomain(input)
		res, err := uh.srv.Register(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("Success create new user", ToResponse(res, "user")))
	}
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		cnv := ToDomain(input)
		res, err := uh.srv.Login(cnv)
		fmt.Println(res.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		tkn := common.GenerateToken(uint(res.ID))

		return c.JSON(http.StatusAccepted, SuccessLogin("Success to login", ToResponse(res, "login"), tkn))
	}
}
