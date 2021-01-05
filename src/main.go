package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)
const dbPath = "../db.json"

func main() {
	ensureDataDirExists()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.File("../client/public/index.html")
	})

	router.GET("/api/blocks", getBlocksHandler)
	router.POST("/api/blocks", updateBlocksHandler)


	router.Static("/build", "../client/public/build")

	router.StaticFile("/favicon.png", "../client/public/favicon.ico")
	router.StaticFile("/index.css", "../client/public/index.css")
	router.StaticFile("/tailwind.css", "../client/public/tailwind.css")

	router.Run()
}

func ensureDataDirExists() {
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		dataFile, err := os.OpenFile(dbPath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer dataFile.Close()

		// empty JSON array
		_, err = dataFile.Write([]byte("[]"))
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
		
	}
}

func save(data []byte) error {
	file, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

func writeErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, err.Error())
}


func updateBlocksHandler(context *gin.Context) {

}

func getBlocksHandler(context *gin.Context) {

}
