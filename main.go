package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	sparta "github.com/mweagle/Sparta"
	"github.com/wickett/wordy/ignorewords"
	"github.com/wickett/wordy/parsing"
)

////////////////////////////////////////////////////////////////////////////////
func paramVal(keyName string, defaultValue string) string {
	value := os.Getenv(keyName)
	if "" == value {
		value = defaultValue
	}
	return value
}

var s3Bucket = paramVal("S3_TEST_BUCKET", "arn:aws:s3:::wickett-sparta")

func wordyEvent(event *json.RawMessage,
	context *sparta.LambdaContext,
	w http.ResponseWriter,
	logger *logrus.Logger) {
	// gonna send back some JSON here
	w.Header().Set("Content-Type", "application/json")
	ignorewords.LoadWords()

	var lambdaEvent sparta.APIGatewayLambdaJSONEvent
	_ = json.Unmarshal([]byte(*event), &lambdaEvent)
	//	text = lambdaEvent.PathParams.Passage

	lookup := lambdaEvent.PathParams["passage"]
	text := parsing.RetrievePassage(lookup)
	//	text := parsing.RetrievePassage("Job1")
	logger.WithFields(logrus.Fields{
		"Event": string(*event),
	}).Info("Request received")

	w.Write(parsing.WordWeight(text))
}

func appendS3Lambda(api *sparta.API, lambdaFunctions []*sparta.LambdaAWSInfo) []*sparta.LambdaAWSInfo {
	options := new(sparta.LambdaFunctionOptions)
	options.Timeout = 30
	lambdaFn := sparta.NewLambda(sparta.IAMRoleDefinition{}, wordyEvent, options)
	apiGatewayResource, _ := api.NewResource("/wordy/lookup/{passage+}", lambdaFn)
	apiGatewayResource.NewMethod("GET", http.StatusOK)

	lambdaFn.Permissions = append(lambdaFn.Permissions, sparta.S3Permission{
		BasePermission: sparta.BasePermission{
			SourceArn: s3Bucket,
		},
		Events: []string{"s3:ObjectCreated:*", "s3:ObjectRemoved:*"},
	})
	return append(lambdaFunctions, lambdaFn)
}

////////////////////////////////////////////////////////////////////////////////
// Return the *[]sparta.LambdaAWSInfo slice
//
func spartaLambdaData(api *sparta.API) []*sparta.LambdaAWSInfo {

	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFunctions = appendS3Lambda(api, lambdaFunctions)
	return lambdaFunctions
}

func main() {
	stage := sparta.NewStage("prod")
	apiGateway := sparta.NewAPIGateway("wordyAPI", stage)
	apiGateway.CORSEnabled = true

	//lambda info
	os.Setenv("AWS_PROFILE", "sparta")
	os.Setenv("AWS_REGION", "us-east-1")

	// Provision a new S3 bucket with the resources in the supplied subdirectory
	s3Site, _ := sparta.NewS3Site("./html")

	stackName := "WordyApplication"
	sparta.Main(stackName,
		"Wordy application",
		spartaLambdaData(apiGateway),
		apiGateway,
		s3Site)

}
