package request

// ProfileRequest GET /profile validation
type ProfileRequest struct {
	Sample    string `form:"sample" validate:"required"`
}
