package radio

import "sort"

type (
	Tag struct {
		Name string
	}
	Tags []*Tag
)

func (ts Tags) Len() int               { return len(ts) }
func (ts Tags) Less(i int, j int) bool { return ts[i].Name < ts[j].Name }
func (ts Tags) Swap(i int, j int)      { ts[i].Name, ts[j].Name = ts[j].Name, ts[i].Name }

func (t *Tag) GetValues() string {
	return t.Name
}

func NewTag(name string) *Tag {
	return &Tag{
		Name: name,
	}
}

func GetTags() Tags {
	var tags Tags
	stations := FetchAllStations()
	tagExist := make(map[string]bool)

	for _, station := range stations {
		for _, tag := range station.Tags {
			if tag.Name == "No tags available!" {
				continue
			}

			if _, exists := tagExist[tag.Name]; exists {
				continue
			}

			tagExist[tag.Name] = true
			newT := NewTag(tag.Name)
			tags = append(tags, newT)
			sort.Sort(tags)
		}

	}

	return tags
}
