package url

import (
	"encoding/json"
	"net/url"
)

type URL struct {
	Scheme string            `json:"schema"`
	Host   string            `json:"host"`
	Path   string            `json:"path"`
	Query  map[string]string `json:"query"`
}

func (u *URL) ToJSON() []byte {
	res, _ := json.Marshal(u)
	return res
}

func Parse(rawURL string) (*URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	q := make(map[string]string)
	if u.RawQuery != "" {
		for param, val := range u.Query() {
			q[param] = val[0]
		}
	}

	return &URL{
		Scheme: u.Scheme,
		Host:   u.Host,
		Path:   u.Path,
		Query:  q,
	}, nil
}
