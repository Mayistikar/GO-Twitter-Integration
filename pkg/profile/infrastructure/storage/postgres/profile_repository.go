package postgres

import (
	"gorm.io/gorm"
	"portfolio/pkg/profile/domain"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db}
}

func (pr *ProfileRepository) Get(profileUUID string) (domain.Profile, error) {
	var profileDTO ProfileDTO
	result := pr.db.Table(TableNameProfile).First(&profileDTO, "uuid = ?", profileUUID)

	if result.Error != nil {
		return domain.Profile{}, result.Error
	}

	profile, err := profileDTO.toProfile()
	if err != nil {
		return domain.Profile{}, err
	}

	return profile, nil
}

func (pr *ProfileRepository) Create(profile *domain.Profile) (*domain.Profile, error) {
	var profileDTO ProfileDTO
	if err := profileDTO.mapProfile(profile); err != nil {
		return &domain.Profile{}, err
	}

	result := pr.db.Table(TableNameProfile).Create(&profileDTO)
	if result.Error != nil {
		return &domain.Profile{}, result.Error
	}
	return profile, nil
}

func (pr *ProfileRepository) Update(profile *domain.Profile) (*domain.Profile, error) {
	var profileDTO ProfileDTO
	if err := profileDTO.mapProfile(profile); err != nil {
		return &domain.Profile{}, err
	}

	result := pr.db.Table(TableNameProfile).Save(&profileDTO)
	if result.Error != nil {
		return &domain.Profile{}, result.Error
	}
	return profile, nil
}
