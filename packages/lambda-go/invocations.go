package remotionlambda

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/go-playground/validator/v10"
)

func invokeRenderLambda(options RemotionOptions) (*RemotionRenderResponse, error) {

	// Create a new AWS session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(options.Region)},
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a new Lambda client
	svc := lambda.New(sess)

	internalParams, validateError := constructRenderInternals(&options)

	if validateError != nil {
		return nil, validateError
	}

	internalParamJsonObject, marshallingError := json.Marshal(internalParams)
	if marshallingError != nil {

		return nil, marshallingError
	}

	invocationPayload := &lambda.InvokeInput{
		FunctionName: aws.String(options.FunctionName),
		Payload:      internalParamJsonObject,
	}

	// Invoke Lambda function
	invocationResult, invocationError := svc.Invoke(invocationPayload)

	if invocationError != nil {
		return nil, invocationError
	}

	// Unmarshal response from Lambda function
	var renderResponseOutput RemotionRenderResponse
	responseMarshallingError := json.Unmarshal(invocationResult.Payload, &renderResponseOutput)
	if responseMarshallingError != nil {
		return nil, responseMarshallingError
	}

	return &renderResponseOutput, nil
}

func invokeRenderProgressLambda(config RenderConfig) (*RenderProgressResponse, error) {

	// Create a new AWS session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(config.Region)},
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a new Lambda client
	svc := lambda.New(sess)

	validate := validator.New()
	validationErrors := validate.Struct(config)

	if validationErrors != nil {
		return nil, validationErrors
	}

	internalParamsJSON, marshallingError := json.Marshal(config)
	if marshallingError != nil {

		return nil, marshallingError
	}

	invocationParams := &lambda.InvokeInput{
		FunctionName: aws.String(config.FunctionName),
		Payload:      internalParamsJSON,
	}

	// Invoke Lambda function
	invokeResult, invokeError := svc.Invoke(invocationParams)

	if invokeError != nil {
		return nil, invokeError
	}

	// Unmarshal response from Lambda function
	var renderProgressOutput RenderProgressResponse
	resultUnmarshallError := json.Unmarshal(invokeResult.Payload, &renderProgressOutput)
	if resultUnmarshallError != nil {
		return nil, resultUnmarshallError
	}

	return &renderProgressOutput, nil
}
