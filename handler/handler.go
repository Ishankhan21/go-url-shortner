package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	handler := func(response http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		newPath, ok := pathsToUrls[path]
		if !ok {
			fallback.ServeHTTP(response, request)
		} else {
			http.Redirect(response, request, newPath, http.StatusSeeOther)
			return
		}
	}
	return handler
}

func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var urlsMap map[string]string

	err := json.Unmarshal(jsonBytes, &urlsMap)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return MapHandler(urlsMap, fallback), nil
}
