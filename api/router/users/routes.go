package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/CGSG-2021-AE4/blog/api"
	"github.com/CGSG-2021-AE4/blog/api/router"
	"github.com/CGSG-2021-AE4/blog/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	Token string `json:"token"`
	Msg   string `json:"msg"`
}

func loginHandler(us api.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("TRY LOGIN")
		var info UserLoginReq
		if err := json.NewDecoder(c.Request.Body).Decode(&info); err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: fmt.Errorf("failed to parse json: %w", err).Error()})
			return
		}
		token, err := us.Login(c, info.Username, info.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: fmt.Errorf("failed to login: %w", err).Error()})
			return
		}
		c.JSON(http.StatusOK, UserLoginResp{Token: string(token), Msg: "Authorization complete"})
		log.Println("LOGIN COMPLETE")
	}
}

type UserRegReq struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegResp struct {
	Msg string `json:"msg"`
}

func registerHandler(us api.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Required header values
		log.Println("TRY REGISTER")
		var info UserRegReq
		if err := json.NewDecoder(c.Request.Body).Decode(&info); err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: fmt.Errorf("failed to parse json: %w", err).Error()})
			return
		}
		log.Println("Info", info)
		user := types.User{
			Id:       uuid.New(),
			Email:    info.Email,
			Username: info.Username,
			Password: info.Password,
		}
		if err := us.Register(c, &user); err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: fmt.Errorf("registration error: %w", err).Error()})
			return
		}
		c.JSON(http.StatusOK, UserRegResp{Msg: "Registration complete"})
	}
}

type getAccountInfoReq struct {
	Username string `json:"username"`
}

func getUserInfoHandler(us api.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("GET account info")
		if c.Keys["authorized"] != "true" {
			log.Println("Not authorized error", c.Keys["authErr"])
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: "not authorized"})
			return
		}
		var info getAccountInfoReq
		if err := json.NewDecoder(c.Request.Body).Decode(&info); err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		u, err := us.GetUserByName(c, info.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		c.JSON(http.StatusOK, u)
	}
}

type deleteUserReq struct {
	Id uuid.UUID `json:"id"`
}

func deleteHandler(us api.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("delete user")
		if c.Keys["authorized"] != "true" {
			log.Println("Not authorized error", c.Keys["authErr"])
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: "not authorized"})
			return
		}
		var info deleteUserReq
		if err := json.NewDecoder(c.Request.Body).Decode(&info); err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		err := us.Delete(c, info.Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		c.JSON(http.StatusOK, "User successfully deleted")
	}
}
