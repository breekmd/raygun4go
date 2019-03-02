package raygun4go

import (
	"io"
	"net/http"
)

//Poster can be used to inject into Client custom POST handler
//This is very useful when using Google App engine and http.Post can not be used
type Poster interface {
	Post(url string, contentType string, body io.Reader, headers *map[string]string) (resp *http.Response, err error)
}

type DefaultPoster struct{
}

func (defaultPoster DefaultPoster) Post(url string, contentType string, body io.Reader, headers *map[string]string) (resp *http.Response, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url , body)
	req.Header.Add("Content-Type", contentType)

	if headers != nil{
		for k, v := range *headers{
			req.Header.Add(k, v)
		}
	}

	resp, err = client.Do(req)

	defer resp.Body.Close()

	return
}