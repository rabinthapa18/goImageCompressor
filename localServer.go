package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func RunLocally(Handler func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) {

	// start the local server
	fmt.Println("Starting the server on localhost:3000")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		reqBody := &ReqBody{}

		err := json.NewDecoder(r.Body).Decode(reqBody)
		if err != nil {
			fmt.Println("Error in decoding request body")
		}

		bucket := reqBody.Bucket
		objectKey := reqBody.ObjectKey
		saveName := reqBody.SaveName

		body := fmt.Sprintf(`{"bucket": "%s", "objectKey": "%s", "saveName": "%s"}`, bucket, objectKey, saveName)

		request := events.APIGatewayProxyRequest{
			Body: body,
		}
		resp, err := Handler(request)
		if err != nil {
			fmt.Println("Error in handling request")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte(resp.Body))

	})
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error in starting server", err)
	}

}
