package application

import (
	"portfolio/pkg/profile/domain"
)

type ProfileApplication struct {
	profileRepository domain.ProfileRepository
}

func NewProfileApplication(profileRepository domain.ProfileRepository) *ProfileApplication {
	return &ProfileApplication{profileRepository}
}

func (p *ProfileApplication) GetProfile(profileID string) (domain.Profile, error) {
	return p.profileRepository.Get(profileID)
}

func (p *ProfileApplication) CreateProfile(profile *domain.Profile) (*domain.Profile, error) {
	return p.profileRepository.Create(profile)
}

func (p *ProfileApplication) UpdateProfile(profile *domain.Profile) (*domain.Profile, error) {
	return p.profileRepository.Update(profile)
}
