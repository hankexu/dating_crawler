package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span>([\d]+)KG</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业：</span>([^<]+)</td>`)
var censusRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var addressRe = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)

var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var constellationRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

var guessRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

var photoRe = regexp.MustCompile(`<img class="hidden" src="(http://[^"]*)" alt="[^>]*>`)

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		if string(match[1]) == "--" {
			return ""
		}
		return string(match[1])
	}
	return ""
}

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Census = extractString(contents, censusRe)
	profile.Car = extractString(contents, carRe)
	profile.Constellation = extractString(contents, constellationRe)
	profile.Education = extractString(contents, educationRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.House = extractString(contents, houseRe)
	profile.Address = extractString(contents, addressRe)
	profile.Photo = extractString(contents, photoRe)
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Id:      extractString([]byte(url), idUrlRe),
				Type:    "zhenai",
				Payload: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				URL:        string(m[1]),
				ParserFunc: ProfileParser(string(m[2])),
			})
	}

	return result
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, name)
	}
}
