package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Замените путь к файлу на свой
	filePath := "test-im.jpeg"

	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Читаем Exif-метаданные из файла
	exifData, err := exif.Read(file)
	if err != nil {
		log.Fatal(err)
	}

	// Получаем все теги метаданных
	tags, err := exifData.AllTags()
	if err != nil {
		log.Fatal(err)
	}

	// Выводим метаданные
	for _, tag := range tags {
		fmt.Printf("%v: %v\n", tag.TagName, tag.Value)
	}
}
