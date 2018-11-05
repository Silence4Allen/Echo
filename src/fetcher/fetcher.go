package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html/charset"
	"bufio"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code : %d", resp.StatusCode)
	}

	reader := bufio.NewReader(resp.Body)
	encodingReader := transform.NewReader(reader, determineEncoding(reader).NewDecoder())
	return ioutil.ReadAll(encodingReader)
}

func determineEncoding(reader *bufio.Reader) encoding.Encoding {
	bytes, err := reader.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	return encoding

}
