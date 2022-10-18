package delivery

import (
	"net/http"
	"sosmed/features/user/common"
	"sosmed/features/user/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}
	e.GET("/users", handler.MyProfile())
	e.PUT("/users", handler.UpdateProfile())
	e.DELETE("/users", handler.Deactivate())
	e.GET("/users/{email}", handler.ShowByEmail())
	e.POST("/register", handler.Register())
	e.POST("/login", handler.Login())
}

func (uh *userHandler) MyProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := common.ExtractToken(c)
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		} else {
			userID, err := strconv.Atoi(c.Param("id"))
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
			userID, err := strconv.Atoi(c.Param("id"))
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
		}
		var data DeleteFormat
		if err := c.Bind(&data); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind delete user"))
		}

		cnv := ToDomain(data)
		res, err := uh.srv.Deactivate(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusAccepted, SuccessResponse("success delete user", ToResponse(res, "")))
	}
}

func (uh *userHandler) ShowByEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Param("email")
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
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusAccepted, SuccessResponse("Success to login", ToResponse(res, "login")))
	}
}
