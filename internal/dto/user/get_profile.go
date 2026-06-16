package user

type GetUserProfileResponseDto struct {
	Id                string  `json:"id"`
	Username          string  `json:"username"`
	Description       *string `json:"description"`
	ProfilePictureUrl *string `json:"profile_picture_url"`
	FollowersCount    int     `json:"followers_count"`
	FollowingCount    int     `json:"following_count"`
}
