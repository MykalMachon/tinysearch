package seeds

import (
	"log/slog"

	"github.com/mykalmachon/tinysearch/indexer/models"
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(dbClient *gorm.DB) error
}

func CreateSources(dbClient *gorm.DB) error {
	return dbClient.Create(&[]models.Source{
		{
			Name:        "Mykal Codes",
			Description: "Mykal Machon's personal website",
			Url:         "https://mykal.codes/feeds/posts.xml",
		},
		{
			Name:        "Dave Rupert",
			Description: "Dave Rupert's personal website",
			Url:         "https://daverupert.com/atom.xml",
		},
		{
			Name:        "Chris Coyier",
			Description: "Chris Coyier's personal website",
			Url:         "https://chriscoyier.net/feed/",
		},
	}).Error
}

var Seeds = []Seed{
	{
		Name: "CreateSources",
		Run:  CreateSources,
	},
}

func All(dbClient *gorm.DB, log slog.Logger) {
	// check if seeds need to be run
	var count int64
	dbClient.Model(&models.Source{}).Count(&count)
	if count > 0 {
		log.Info("seeds have already been run... skipping!")
		return
	}

	for _, seed := range Seeds {
		err := seed.Run(dbClient)
		if err != nil {
			log.Error("failed to run seed", "seed", seed.Name, "error", err.Error())
		}
	}
}
