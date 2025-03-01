package remotionlambda

import (
	"encoding/json"
	"log"
	"testing"
)

func TestPrintVersion(t *testing.T) {
	payload, err := constructRenderInternals(&RemotionOptions{
		Region:       "us-east-1",
		Composition:  "react-svg",
		FunctionName: "remotion-render",
		ServeUrl:     "testbed",
		Codec:        "h264",
	})
	if err != nil {
		log.Fatalf("Error marshaling struct to JSON: %v", err)
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshaling struct to JSON: %v", err)
	}

	println(string(jsonData))
}
