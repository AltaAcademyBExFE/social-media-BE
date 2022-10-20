package delivery

import (
	"context"
	"net/http"
	sc "sosmed/config"
	"sosmed/features/post/domain"
	"sosmed/utils/common"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type postHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := postHandler{srv: srv}
	e.GET("/posts", handler.ShowAllPost())
	e.GET("/posts/me", handler.ShowMyPost(), middleware.JWT([]byte(sc.JwtKey)))
	e.GET("/posts/:id", handler.ShowSpesificPost(), middleware.JWT([]byte(sc.JwtKey)))
	e.POST("/posts", handler.CreatePost(), middleware.JWT([]byte(sc.JwtKey)))
	e.PUT("/posts/:id", handler.EditPost(), middleware.JWT([]byte(sc.JwtKey)))
	e.DELETE("/posts/:id", handler.DeletePost(), middleware.JWT([]byte(sc.JwtKey)))
}

func (ph *postHandler) ShowAllPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, rel, err := ph.srv.ShowAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success get all post", ToResponse(res, rel, "all")))
	}
}

func (ps *postHandler) ShowMyPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		//ID, _ := strconv.Atoi(c.Param("id"))
		userID := common.ExtractToken(c)
		res, rel, err := ps.srv.ShowMy(userID)
		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success get my post", ToResponse(res, rel, "all")))
	}
}

func (ps *postHandler) ShowSpesificPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))
		res, rel, err := ps.srv.ShowSpesific(ID)
		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success get my post", ToResponse(res, rel, "ally")))
	}
}

func (ph *postHandler) CreatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input PostingFormat
		userID := common.ExtractToken(c)
		input.UserID = userID

		errEnv := godotenv.Load("config.env")
		if errEnv != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("Error loading .env file"))
		}

		cfg, errDef := config.LoadDefaultConfig(context.TODO())
		if errDef != nil {
			var erroDef string = "Error: "
			erroDef += erroDef
			return c.JSON(http.StatusBadRequest, FailResponse(erroDef))
		}

		client := s3.NewFromConfig(cfg)
		uploader := manager.NewUploader(client)

		isSuccess := true
		file, er := c.FormFile("img")
		if er != nil {
			isSuccess = false
		} else {
			src, err := file.Open()
			if err != nil {
				isSuccess = false
			} else {
				result, errImg := uploader.Upload(context.TODO(), &s3.PutObjectInput{
					Bucket: aws.String("project-sosmed"),
					Key:    aws.String(file.Filename),
					Body:   src,
					ACL:    "public-read",
				})

				if errImg != nil {
					return c.JSON(http.StatusBadRequest, FailResponse("Berhasil Upload Images"))
				}

				input.Images = result.Location
			}
			defer src.Close()
		}

		if isSuccess {
			if err := c.Bind(&input); err != nil {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
			}
			cnv := ToDomain(input)
			res, err := ph.srv.Create(cnv)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}

			return c.JSON(http.StatusCreated, SuccessResponse("Success create new post", ToResponse(res, nil, "post")))
		}
		return c.JSON(http.StatusBadRequest, FailResponse("fail upload file"))
	}
}

func (ph *postHandler) EditPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))
		var input PostingFormat
		userID := common.ExtractToken(c)
		input.UserID = userID
		errEnv := godotenv.Load("config.env")
		if errEnv != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("Error loading .env file"))
		}

		cfg, errDef := config.LoadDefaultConfig(context.TODO())
		if errDef != nil {
			var erroDef string = "Error: "
			erroDef += erroDef
			return c.JSON(http.StatusBadRequest, FailResponse(erroDef))
		}

		client := s3.NewFromConfig(cfg)
		uploader := manager.NewUploader(client)

		isSuccess := true
		file, er := c.FormFile("img")
		if er != nil {
			isSuccess = false
		} else {
			src, err := file.Open()
			if err != nil {
				isSuccess = false
			} else {
				result, errImg := uploader.Upload(context.TODO(), &s3.PutObjectInput{
					Bucket: aws.String("project-sosmed"),
					Key:    aws.String(file.Filename),
					Body:   src,
					ACL:    "public-read",
				})

				if errImg != nil {
					return c.JSON(http.StatusBadRequest, FailResponse("Berhasil Upload Images"))
				}

				input.Images = result.Location
			}
			defer src.Close()
		}

		if isSuccess {
			if err := c.Bind(&input); err != nil {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
			}
			cnv := ToDomain(input)
			res, err := ph.srv.Edit(ID, cnv)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}

			return c.JSON(http.StatusCreated, SuccessResponse("Success edit post", ToResponse(res, nil, "post")))
		}
		return c.JSON(http.StatusBadRequest, FailResponse("fail upload file"))
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
		return c.JSON(http.StatusOK, FailResponse("Success delete post"))
	}
}
