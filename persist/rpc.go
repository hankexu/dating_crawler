package persist

import (
	"crawler/engine"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := Save(s.Client, s.Index, item)
	if err == nil {
		*result = "OK"
		log.Printf("Item %v saved.", item)
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}
