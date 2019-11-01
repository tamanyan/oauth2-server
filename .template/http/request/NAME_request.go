package request

// {{camelcase .NAME}}Request GET /{{lowercaseletters .NAME}} validation
type {{camelcase .NAME}}Request struct {
	Sample    string `form:"sample" validate:"required"`
}
