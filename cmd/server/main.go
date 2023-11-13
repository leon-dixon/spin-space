package main

import (
	"context"
	"spin-space/internal/api/handlers"
	"spin-space/internal/model"
	"spin-space/internal/store"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	vinylStore := store.NewVinylRecordStore()
	vinylHandler := handlers.NewVinylHandler(vinylStore)

	dummyVinyls := []model.VinylRecord{
		{
			Artist:    "Artist 1",
			AlbumName: "Album 1",
			Year:      1980,
			// ... other fields
		},
		{
			Artist:    "Artist 2",
			AlbumName: "Album 2",
			Year:      1990,
			// ... other fields
		},
		// ... more records
	}

	for _, v := range dummyVinyls {
		vinyl := v
		vinylStore.Create(context.Background(), &vinyl)
	}

	router.GET("/vinyls", vinylHandler.GetVinyls)
	router.GET("/vinyls/:id", vinylHandler.GetVinylByID)

	router.Run(":8080")
}
