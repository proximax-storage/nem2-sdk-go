package utils

import "net/url"

type Url struct {
	path   string
	values url.Values
}

func NewUrl(path string) *Url {
	return &Url{
		path:   path,
		values: url.Values{},
	}
}

func (u *Url) SetParam(key, val string) {
	u.values.Set(key, val)
}

func (u *Url) AddParam(key, val string) {
	u.values.Add(key, val)
}

func (u *Url) GetParam(key string) string {
	return u.values.Get(key)
}

func (u *Url) DeleteParam(key string) {
	u.values.Del(key)
}

func (u *Url) Encode() string {
	if len(u.values) == 0 {
		return u.path
	}

	return u.path + "?" + u.values.Encode()
}
