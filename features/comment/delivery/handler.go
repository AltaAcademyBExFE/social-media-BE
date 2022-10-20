package delivery

import (
	"net/http"
	"sosmed/config"
	"sosmed/features/comment/domain"
	"sosmed/utils/common"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type commentHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := commentHandler{srv: srv}
	e.POST("/comments", handler.CreateComment(), middleware.JWT([]byte(config.JwtKey)))
	e.DELETE("/comments/:id", handler.DeleteComment(), middleware.JWT([]byte(config.JwtKey)))
}

func (ch *commentHandler) CreateComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input CommentFormat
		userID := common.ExtractToken(c)
		input.UserID = userID
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := ch.srv.Create(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success create new comment", ToResponse(res, "comment")))
	}
}

func (ch *commentHandler) DeleteComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))
		err := ch.srv.Delete(ID)
		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, FailResponse("Success delete comment"))
	}
}
