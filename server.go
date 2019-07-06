package main

import (
	"encoding/json"
	"github.com/c8112002/netflix_clone_server/models"
	"log"
	"net/http"
	"strconv"
)

const port = 8000

func getItems(w http.ResponseWriter, r *http.Request)  {

	titles := []string{
		"aDOPE",
		"aaLUCIFER",
		"aバッド・ブラッド悲しみのマフィア",
		"aaサバイバー",
		"aオレンジイズニューブラック",
		"aaストレンジャーシングス",
		"aDARK",
		"aaブラック・ミラー",
		"aシャフト",
		"aaジェシカジョーンズ",
		"aブラック・ミラ バンダースナッチ",
		"aaWHATIF 選択の連鎖",
		"aDOPE",
		"aLUCIFER",
		"aバッド・ブラッド悲しみのマフィア",
		"aサバイバー",
		"aオレンジイズニューブラック",
		"aストレンジャーシングス",
		"aDARK",
		"aブラック・ミラー",
		"aシャフト",
		"aジェシカジョーンズ",
		"aブラック・ミラ バンダースナッチ",
		"aWHATIF 選択の連鎖",
	}
	var items []models.Item
	for i := range titles {
		id := i + 1
		title := titles[i]
		imageName := strconv.Itoa((i % 12) + 1) + ".jpg"
		imageUrl := "http://localhost:" + strconv.Itoa(port) + "/" + imageName
		item := models.Item{Id:id, Title: title, ImageUrl: imageUrl}
		items = append(items, item)
	}

	res,err := json.Marshal(items)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	http.HandleFunc("/items", getItems)
	http.Handle("/", http.FileServer(http.Dir("assets/images")))
	err := http.ListenAndServe(":" + strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
