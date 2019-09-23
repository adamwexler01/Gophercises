package jsonparser

import(
  "encoding/json"
)

func JsonParser(jsonbytes []byte) (map[string]topic, error){
  jsonMap := map[string]topic{}
  error := json.Unmarshal(jsonbytes, &jsonMap)
  if error != nil {
    return nil, error
  }

  return jsonMap, nil
}

type option struct {
  Text string
  Arc string
}

type topic struct {
  Title string
  Story []string
  Options []option
}
