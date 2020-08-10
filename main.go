package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jainSamkit/youtubeDownloader/models/ytdownloader"
	"github.com/jainSamkit/youtubeDownloader/types"
)

//YoutubeURL struct
type YoutubeURL struct {
	URL string `json:"url"`
}

//ResponseData ...
type ResponseData struct {
	Links []types.VideoLink `json:"links"`
	ERROR string            `json:"error"`
}

//Handler called on start
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	videoURL := YoutubeURL{}
	body := []byte(req.Body)
	err := json.Unmarshal(body, &videoURL)

<<<<<<< HEAD
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: err.Error()}, nil
	}
=======
	// //json error
	s := "https://www.youtube.com/watch?v=k9zTr2MAFRg"
>>>>>>> master

	d := ytdownloader.New(videoURL.URL)
	videoLinks, responsePipe := d.GetVideoLinks()

<<<<<<< HEAD
	res := ResponseData{}
	res.Links = videoLinks
=======
	// s := "https://www.youtube.com/watch?v=1I-3vJSC-Vo"
>>>>>>> master

	if responsePipe.Err != nil {
		res.ERROR = responsePipe.Err.Error()
	} else {
		res.ERROR = ""
	}

	fmt.Println("Video links: ", videoLinks)
	fmt.Println("Response Pipe:", responsePipe)

	resJSON, _ := json.Marshal(res)

	fmt.Println("printing final result json: ", string(resJSON))

	if res.ERROR != "" {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: string(resJSON)}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(resJSON)}, nil

}

func main() {
	lambda.Start(Handler)
}
