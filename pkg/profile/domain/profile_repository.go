package domain

type ProfileRepository interface {
	Get(UUID string) (Profile, error)
	Create(*Profile) (*Profile, error)
	Update(*Profile) (*Profile, error)
}
