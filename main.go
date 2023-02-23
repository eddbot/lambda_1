package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Input struct {
	EmployeeID string `json:"employeeId"`
}

type Output struct {
	EmployeeId  string `json:"employeeId"`
	LeavingDate string `json:"leavingDate"`
	Email       string `json:"email"`
}

func handleRequest(ctx context.Context, i Input) (Output, error) {
	out := Output{
		EmployeeId:  i.EmployeeID,
		LeavingDate: "date",
		Email:       "dave@test.com",
	}
	return out, nil
}

func main() {
	lambda.Start(handleRequest)
}
