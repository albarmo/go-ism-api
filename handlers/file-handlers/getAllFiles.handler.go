package fileHandlers

import (
	"ism/models"
	"ism/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAllFilesHandler(context *gin.Context) {
	jwtData, _ := context.Get("user")

	var user models.UserEntity
	// convert header to user enitity
	errors := utils.StringToEntity(jwtData, &user)

	if errors != nil {
		utils.APIResponse(context, "User does not exist", http.StatusNotFound, http.MethodGet, nil)
		return
	}
	userId := user.ID

	fileResponse, statusCode := h.service.GetAllFiles(userId)

	switch statusCode {
	case http.StatusOK:

		//  populate the url field
		for index := range fileResponse {
			file := &fileResponse[index]
			fileUrl := utils.GetFileUrl(file.AccessKey)
			file.Url = fileUrl
		}
		utils.APIResponse(context, "Received files", http.StatusOK, http.MethodGet, &fileResponse)
		return

	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occured", http.StatusExpectationFailed, http.MethodGet, nil)
		return

	case http.StatusConflict:
		utils.APIResponse(context, "File already exists. Please try with another file", http.StatusConflict, http.MethodGet, nil)
		return
	}
}
