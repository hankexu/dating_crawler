package model

import "encoding/json"

type Profile struct {
	Name          string
	Gender        string
	Age           int
	Height        int
	Weight        int
	Income        string
	Marriage      string // 婚况
	Education     string
	Occupation    string // 职业
	Census        string // 户籍
	Constellation string // 星座
	House         string
	Car           string
	Address       string
	Photo         string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
