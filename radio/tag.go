package radio

import (
	json2 "encoding/json"
	"strconv"
)

const (
	TagsURL = "https://de1.api.radio-browser.info/json/tags"
)

type (
	Tag struct {
		Name string
	}
	Tags []*Tag
)

func (ts Tags) Len() int               { return len(ts) }
func (ts Tags) Less(i int, j int) bool { return ts[i].Name < ts[j].Name }
func (ts Tags) Swap(i int, j int)      { ts[i], ts[j] = ts[j], ts[i] }

func (t *Tag) GetValues() string {
	return t.Name
}

func NewTag(name string) *Tag {
	return &Tag{
		Name: name,
	}
}

func FetchTags() *Tags {
	res := Post(TagsURL, "", nil)
	return UnmarshalTags(res)
}

func FetchTagsDetailed(order Order, reverse bool, hideBroken bool) *Tags {
	q := make(map[string]string)
	q["order"] = string(order)
	q["reverse"] = strconv.FormatBool(reverse)
	q["hidebroken"] = strconv.FormatBool(hideBroken)
	res := Post(TagsURL, "", q)
	return UnmarshalTags(res)
}

func UnmarshalTags(json string) *Tags {
	var languages *Tags
	json2.Unmarshal([]byte(json), &languages)
	return languages
}

func GetTags() *Tags {
	var tags Tags
	fetch := FetchTagsDetailed(OrderName, false, true)
	tagExist := make(map[string]bool)
	for _, t := range *fetch {

		if _, exists := tagExist[t.Name]; exists {
			continue
		}
		tagExist[t.Name] = true
		tags = append(tags, t)
	}

	return &tags
}
