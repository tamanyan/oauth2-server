package request

// UserRequest GET /user validation
type UserRequest struct {
	Sample    string `form:"sample" validate:"required"`
}
