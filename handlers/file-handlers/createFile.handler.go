package fileHandlers

import (
	filecontrollers "ism/controllers/file-controllers"
	"ism/models"
	"ism/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateHandler(context *gin.Context) {
	file, header, _ := context.Request.FormFile("file")

	var user models.UserEntity

	jwtData, _ := context.Get("user")

	// convert header to user enitity
	errors := utils.StringToEntity(jwtData, &user)

	if errors != nil {
		utils.APIResponse(context, "User does not exist", http.StatusNotFound, http.MethodPost, nil)
		return
	}
	result, err := utils.UploadFile(file, header.Header.Get("Content-Type"))

	if err != nil {
		utils.APIResponse(context, "Unable to upload file to the server", http.StatusFailedDependency, http.MethodPost, nil)
	}

	fileInput := filecontrollers.FileInput{
		ID:     result.PublicID,
		Type:   result.Format,
		Name:   header.Filename,
		UserId: user.ID,
	}

	fileResponse, statusCode := h.service.CreateFile(&fileInput)

	if statusCode != http.StatusCreated {
		//  delete the file in cloudinary if it is not created in the DB
		utils.DeleteFile(result.PublicID)
	}

	switch statusCode {
	case http.StatusCreated:
		utils.APIResponse(context, "Uploaded the file successfully.", http.StatusCreated, http.MethodPost, fileResponse)
		return

	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occured", http.StatusExpectationFailed, http.MethodPost, nil)
		return

	case http.StatusConflict:
		utils.APIResponse(context, "File already exists. Please try with another file", http.StatusCreated, http.MethodPost, nil)
		return
	}

}
