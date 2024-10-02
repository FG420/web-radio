package radio

import (
	"sort"
	"strings"
	"sync"
)

var (
	tagCache Tags
	once     sync.Once
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

func GetTags() Tags {
	once.Do(func() {
		var tags Tags
		stations := FetchAllStations()
		tagExist := make(map[string]bool)

		for _, station := range stations {
			for _, tag := range station.Tags {
				if strings.Contains(tag.Name, "http") {
					continue
				}
				if _, exists := tagExist[tag.Name]; exists {
					continue
				}

				tagExist[tag.Name] = true
				tags = append(tags, &tag)
			}
		}

		sort.Sort(tags)
		tagCache = tags
	})

	return tagCache
}
