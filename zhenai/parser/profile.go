package parser

import (
	"regexp"
	"strconv"

	"../../engine"
	"../../model"
)

var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)

var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(` <td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)

var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var householdRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var constellationRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		profile.Weight = weight
	}

	profile.Marriage = extractString(contents, marriageRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hokou = extractString(contents, householdRe)
	profile.Xinzuo = extractString(contents, constellationRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result

}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

}
