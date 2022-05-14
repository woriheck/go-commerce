package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/api/idtoken"
)

func main() {
	http.HandleFunc("/", HelloProduct)
	http.ListenAndServe(":8080", nil)
}

type Response struct {
	Message string `json:"message"`
}

func HelloProduct(w http.ResponseWriter, r *http.Request) {
	targetURL := "https://dev-backend-bprvtpcz2a-as.a.run.app/pricing"
	audience := "https://dev-backend-bprvtpcz2a-as.a.run.app/url"
	makeGetRequest(w, targetURL, audience)

	jsonOut, _ := json.Marshal(Response{Message: "Hello Product"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsonOut)
}

// `makeGetRequest` makes a request to the provided `targetURL`
// with an authenticated client using audience `audience`.
func makeGetRequest(w io.Writer, targetURL string, audience string) error {
	// Example `audience` value (Cloud Run): https://my-cloud-run-service.run.app/
	// (`targetURL` and `audience` will differ for non-root URLs and GET parameters)
	ctx := context.Background()

	// client is a http.Client that automatically adds an "Authorization" header
	// to any requests made.
	client, err := idtoken.NewClient(ctx, audience)
	if err != nil {
		return fmt.Errorf("idtoken.NewClient: %v", err)
	}

	resp, err := client.Get(targetURL)
	if err != nil {
		return fmt.Errorf("client.Get: %v", err)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	return nil
}
