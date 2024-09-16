package radio

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func Post(url string, payload string, query map[string]string) string {

	payloadReader := strings.NewReader(payload)
	req, _ := http.NewRequest("GET", url, payloadReader)

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return string(body)
}
