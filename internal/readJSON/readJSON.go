package readJSON

import(
	"net/http"
	"io"
)

func Read(url string) ([]byte, error) {
	
	res, err := http.Get(url)
	
	if err != nil{
		return nil, err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil{
		return nil, err
	}
	
	return body, nil

}
