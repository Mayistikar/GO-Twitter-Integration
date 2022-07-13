package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portfolio/pkg/profile/application"
)

type ProfileHandler struct {
	profileApplication *application.ProfileApplication
}

func NewProfileHandler(profileApplication *application.ProfileApplication) *ProfileHandler {
	return &ProfileHandler{profileApplication}
}

func (p *ProfileHandler) GetProfile(ctx *gin.Context) {
	var profileRequestDTO RequestDTO
	if err := ctx.ShouldBindUri(&profileRequestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	profile, err := p.profileApplication.GetProfile(profileRequestDTO.UUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var profileResponseDTO ResponseDTO
	err = profileResponseDTO.mapProfile(profile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, profileResponseDTO)
}

func (p *ProfileHandler) CreateProfile(ctx *gin.Context) {
	var profileRequestDTO RequestDTO
	if err := ctx.ShouldBind(&profileRequestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	profile, err := profileRequestDTO.toProfile()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newProfile, err := p.profileApplication.CreateProfile(&profile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var profileResponseDTO ResponseDTO
	if err := profileResponseDTO.mapProfile(*newProfile); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, profileResponseDTO)
}

func (p *ProfileHandler) UpdateProfile(ctx *gin.Context) {
	var profileRequestDTO RequestDTO
	if err := ctx.ShouldBindUri(&profileRequestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&profileRequestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	profile, err := profileRequestDTO.toProfile()
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedProfile, err := p.profileApplication.UpdateProfile(&profile)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"error": err.Error(),
		})
		return
	}

	var profileResponseDTO ResponseDTO
	if err := profileResponseDTO.mapProfile(*updatedProfile); err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, profileResponseDTO)
}
