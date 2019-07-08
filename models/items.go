package models

import (
	"regexp"
	"strconv"
)


type Items []Item

func GetAllItems() Items {
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
	var items Items
	for i := range titles {
		id := i + 1
		title := titles[i]
		imageName := strconv.Itoa((i % 12) + 1) + ".jpg"
		imageUrl := "http://localhost:8000/" + imageName
		item := Item{Id:id, Title: title, ImageUrl: imageUrl}
		items = append(items, item)
	}

	return items
}

func GetPopularItems() Items {
	titles := []string{
		"DOPE",
		"LUCIFER",
		"バッド・ブラッド悲しみのマフィア",
		"サバイバー",
		"オレンジイズニューブラック",
		"ストレンジャーシングス",
		"DARK",
		"ブラック・ミラー",
		"シャフト",
		"ジェシカジョーンズ",
		"ブラック・ミラ バンダースナッチ",
		"WHATIF 選択の連鎖",
	}
	var items Items
	for i := range titles {
		id := i + 1
		title := titles[i]
		imageName := strconv.Itoa((i % 12) + 1) + ".jpg"
		imageUrl := "http://localhost:8000/popular/" + imageName
		item := Item{Id:id, Title: title, ImageUrl: imageUrl}
		items = append(items, item)
	}

	return items
}

func (i Items) Filter(title string) Items {
	items := Items{}

	r := regexp.MustCompile(title)

	for _, item := range i {
		if r.MatchString(item.Title) {
			items = append(items, item)
		}
	}

	return items
}

