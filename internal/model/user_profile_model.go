package model

type UserProfileModel struct {
	ID                string
	Username          string
	Description       *string
	ProfilePictureUrl *string
	FollowersCount    int
	FollowingCount    int
}
