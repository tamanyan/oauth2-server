package request

// AdminSignupRequest GET /admin validation
type AdminSignupRequest struct {
	Sample string `form:"sample" validate:"required"`
}
