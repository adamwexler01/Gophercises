package urlshort

import (
	"net/http"
	"gopkg.in/yaml.v2"
	"encoding/json"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//Don't need to cast the function with .handlerfunc because
	//it is already implied through the method signature
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request){
		//When checking a map, an ok value will be apart of the response
		//that value will be true if the value exists and false if it doesn't
		if dest, ok := pathsToUrls[request.URL.Path]; ok {
				http.Redirect(writer, request, dest, http.StatusFound)
				return
		} else {
				fallback.ServeHTTP(writer, request)
		}
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	// var yamlOutput[2] map[string]string
	var yamlOutput []uri
	output := map[string]string{}
	error := yaml.Unmarshal(yml, &yamlOutput)
	if error != nil {
		return nil, error
	}
	for _, value := range yamlOutput {
		output[value.Path] = value.Url
	}

	return MapHandler(output, fallback), nil
}

func JSONHandler(jsonByte []byte, fallback http.Handler) (http.HandlerFunc, error){
	var jsonToStruct []uri
	error := json.Unmarshal(jsonByte, &jsonToStruct)
	if error != nil {
		return nil, error
	}
	jsonMap := map[string]string{}
	for _, value := range jsonToStruct {
		jsonMap[value.Path] = value.Url
	}

	return MapHandler(jsonMap, fallback), nil
}

type uri struct {
	Path string `yaml:"path"`
	Url string `yaml:"url"`
}
