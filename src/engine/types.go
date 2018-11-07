package engine

/**
Url:the target of this request
ParserFunc:deal with this Url
*/
type Request struct {
	Url        string
	ParserFunc ParserFunc
}

/**
Url:the url of this Item
Name:maybe the name or the type of this Item
Id:the id of this Item
Payload:to save other data
 */
type Item struct {
	//Name           string
	Id          string
	Url         string
	Type        string
	PayloadType string
	//ItemSourcePath string
	Payload interface{}
}

/**
Requests: the new request
Items: the target data
 */
type ParseResult struct {
	Requests []Request
	Items    []Item
}

/**
deal with the contents and then return the result
 */
type ParserFunc func(contents []byte, url string) ParseResult
