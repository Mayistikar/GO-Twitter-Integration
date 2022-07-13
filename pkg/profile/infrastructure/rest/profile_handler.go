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
	var profileDTO ProfileDTO
	if err := ctx.ShouldBindUri(&profileDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	profile, err := p.profileApplication.GetProfile(profileDTO.UUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = profileDTO.mapProfile(profile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, profileDTO)
}

func (p *ProfileHandler) CreateProfile(ctx *gin.Context) {
	var profileDTO ProfileDTO
	if err := ctx.ShouldBind(&profileDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	profile, err := profileDTO.toProfile()
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

	var newProfileDTO ProfileDTO
	if err := newProfileDTO.mapProfile(*newProfile); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, newProfileDTO)
}

func (p *ProfileHandler) UpdateProfile(ctx *gin.Context) {
	var profileDTO ProfileDTO
	if err := ctx.ShouldBindUri(&profileDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&profileDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	profile, err := profileDTO.toProfile()
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

	var updatedProfileDTO ProfileDTO
	if err := updatedProfileDTO.mapProfile(*updatedProfile); err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedProfileDTO)

}
