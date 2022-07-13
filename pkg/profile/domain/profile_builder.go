package domain

import "fmt"

type ProfileBuilder interface {
	SetUUID(uuid string) ProfileBuilder
	SetImage(image string) ProfileBuilder
	SetTitle(title string) ProfileBuilder
	SetDescription(description string) ProfileBuilder
	Build() (Profile, error)
}

type profileBuilder struct {
	Profile
	error
}

func NewProfileBuilder() *profileBuilder {
	return new(profileBuilder)
}

func (pb *profileBuilder) SetUUID(uuid string) *profileBuilder {
	if uuid == "" {
		pb.error = fmt.Errorf("%w: %v", ErrorProfileInvalidImage, pb.error)
		return pb
	}
	pb.uuid = uuid
	return pb
}

func (pb *profileBuilder) SetImage(image string) *profileBuilder {
	if image == "" {
		pb.error = fmt.Errorf("%w: %v", ErrorProfileInvalidImage, pb.error)
		return pb
	}
	pb.image = image
	return pb
}

func (pb *profileBuilder) SetTitle(title string) *profileBuilder {
	if title == "" {
		pb.error = fmt.Errorf("%w: %v", ErrorProfileInvalidTitle, pb.error)
		return pb
	}
	pb.title = title
	return pb
}

func (pb *profileBuilder) SetDescription(description string) *profileBuilder {
	if len(description) < 10 || len(description) > 2500 {
		pb.error = fmt.Errorf("%w: %v", ErrorProfileInvalidDescription, pb.error)
		return pb
	}
	pb.description = description
	return pb
}

func (pb *profileBuilder) Build() (Profile, error) {
	return pb.Profile, pb.error
}
