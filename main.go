package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"postgres/internal/config"
	"postgres/internal/entity"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Get the config from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load db config
	db, err := config.OpenDB(os.Getenv("POSTGRES_URL"), true)
	if err != nil {
		log.Fatalln(err)
	}
	defer config.CloseDB(db)

	// Init clean arch
	repository := config.InitRepository(db)
	usecase := config.InitUsecase(repository.AlbumRepository)
	handler := config.InitHandler(usecase.AlbumUsecase)

	// Using the handler

	// Get Handler
	// You can use/change the code below if you want to test the handler
	// albums, err := handler.AlbumHandler.GetAllAlbum()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(albums)

	// =============================================================
	//Insert Handler
	// album, err := handler.AlbumHandler.Create(&entity.Album{
	// 	Title: "Title-Insert via VSCode-4",
	// 	Artist: "Artist-Insert via VSCode-4",
	// 	Price: 2500,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// b, _ := json.Marshal(album)
	// fmt.Println((string(b)))
	// fmt.Println("Berhasil")

	// =============================================================
	//Batch Create Handler
	albums, err := handler.AlbumHandler.BatchCreate([]entity.Album{
		{
			Title: "Title-Insert via VSCode-5",
			Artist: "Artist-Insert via VSCode-5",
			Price: 2500,
		},
		{
			Title: "Title-Insert via VSCode-6",
			Artist: "Artist-Insert via VSCode-6",
			Price: 2500,
		},
	})
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(albums)
	fmt.Println((string(b)))
	fmt.Println("Berhasil")
}
