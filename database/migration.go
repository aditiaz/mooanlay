package database

import (
	"fmt"
	"moonlay/models"
	"moonlay/pkg/postgresql"
)

func RunMigration() {
	err := postgresql.DB.AutoMigrate(
		&models.List{},
		&models.SubList{},
		&models.PostImage{},
		&models.SubList{},
		&models.PostImageSub{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
