package request

// AdminRequest GET /admin validation
type AdminRequest struct {
	Sample    string `form:"sample" validate:"required"`
}
