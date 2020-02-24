package models

//ExpressionRequest is the model used for the expression endpoint
type ExpressionRequest struct {
	Value string `json: "Value"`
}

//ExpressionResponse is the model used for the expression endpoint
type ExpressionResponse struct {
	Statement string `json: "Statement"`
}
