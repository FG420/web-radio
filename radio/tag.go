package radio

type Tag struct {
	Value string
}

func (t *Tag) GetValues() string {
	return t.Value
}

func NewTag(name string) *Tag {
	return &Tag{
		Value: name,
	}
}

// func GetTags() []Tag {
// 	stations := FetchAllStations()

// 	var tags []Tag
// 	tagExist := make(map[string]bool)

// 	for _, station := range stations {
// 		count := 0
// 		if station.Tags[count] == "" {
// 			continue
// 		}

// 		if _, exists := tagExist[station.Tags[count]]; exists {
// 			continue
// 		}

// 		count++

// 		// arrayTags := strings.Split(station.Tags, ",")
// 		// for _, tag := range arrayTags {
// 		// 	tt := NewTag(tag)
// 		// 	tags = append(tags, *tt)
// 		// }

// 	}

// 	log.Print(len(tags))
// 	return tags
// }
