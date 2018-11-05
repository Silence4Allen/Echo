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
Type:maybe the name or the type of this Item
Id:the id of this Item
Payload:to save other data
 */
type Item struct {
	Url     string
	Type    string
	Id      string
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
