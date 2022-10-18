package delivery

import (
	"net/http"
	"sosmed/features/post/domain"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type postHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := postHandler{srv: srv}
	e.GET("/posts", handler.ShowAllPost())
	e.GET("/posts/me", handler.ShowMyPost())
	e.GET("/posts/:id", handler.ShowSpesificPost())
	e.POST("/posts", handler.CreatePost())
	e.PUT("/posts/:id", handler.EditPost())
	e.DELETE("/posts/:id", handler.DeletePost())
}

func (ph *postHandler) ShowAllPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ph.srv.ShowAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success get all post", ToResponse(res, "all")))
	}
}

func (ps *postHandler) ShowMyPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		//ID, _ := strconv.Atoi(c.Param("id"))
		res, err := ps.srv.ShowMy(1)
		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success get my post", ToResponse(res, "all")))
	}
}

func (ps *postHandler) ShowSpesificPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))
		res, err := ps.srv.ShowSpesific(ID)
		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success get my post", ToResponse(res, "all")))
	}
}

func (ph *postHandler) CreatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input PostingFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := ph.srv.Create(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success create new post", ToResponse(res, "post")))
	}
}

func (ph *postHandler) EditPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))
		var input PostingFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := ph.srv.Edit(ID, cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success edit post", ToResponse(res, "post")))
	}
}

func (ph *postHandler) DeletePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))
		err := ph.srv.Delete(ID)
		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, "Success delete post")
	}
}
