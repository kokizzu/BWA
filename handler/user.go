package handler

import (
	"BWA/auth"
	"BWA/helper"
	"BWA/rpcp"
	"BWA/user"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.Service
	authService auth.Service	
	rpcp.UnimplementedUserServiceServer
}

func NewUserHandler(userService user.Service, authService auth.Service) *UserHandler {
	return &UserHandler{
		userService: userService, 
		authService: authService,
	}
}

func (h *UserHandler) RegisterUserGrpc(ctx context.Context, in *rpcp.RegisterUserInput) (out *rpcp.RegisterUserOutput, err error) {
	
	input := &user.RegisterUserInput{}
	response, formatter := h.registerUser(input.FromProto(in))
	
	response.ToMetaProto(out.Meta)
	if formatter == nil {
		err = errors.New(response.Meta.Message)
		return
	}
	formatter.ToDataProto(out.Data)
	return
}

func (h *UserHandler) registerUser(input *user.RegisterUserInput) (helper.Response, *user.UserFormater) {
	
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		// error disini waktu mau insert ke db
		response := helper.APIResponse("register account failed", http.StatusBadRequest, "failed", nil)
		return response, nil
	}

	authToken, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("register account failed", http.StatusBadRequest, "failed", nil)
		return response, nil

	}

	formatter := user.FormatUser(newUser, authToken)

	response := helper.APIResponse("account has been created", http.StatusOK, "success", formatter)
	return response, &formatter
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	//tangkap input dari usser
	//map dari user input ke registeruserinput

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response, _ := h.registerUser(&input)

	c.JSON(response.Meta.Code, response)
}

func (h *UserHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("loggedin failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	logedInUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}

		response := helper.APIResponse("loggedin failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	authToken, err := h.authService.GenerateToken(logedInUser.ID)
	if err != nil {
		response := helper.APIResponse("login account failed", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := user.FormatUser(logedInUser, authToken)
	response := helper.APIResponse("loggedin success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *UserHandler) CheckEmailAvailability(c *gin.Context) {
	//input emAIL DARI USER
	//input di mapping ke struct input
	// struct input dipassing ke service
	//dari service memanggil repository, apakah email sudah ada
	//repo querry ke db

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvalable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"error": "server error"}
		response := helper.APIResponse("email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_availabale": isEmailAvalable,
	}

	var metaMessage string
	if isEmailAvalable {
		metaMessage = "email is available"
	} else {
		metaMessage = "email has been used by another"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}

func (h *UserHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
