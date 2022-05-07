package controllers

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	imageNames []string
	imageQueue []string
)

func RandomImage(c *gin.Context) {
	imagePath := viper.GetString("Image.Path")

	if len(imageNames) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no images found"})
		return
	}

	if len(imageQueue) == 0 {
		imageQueue = make([]string, len(imageNames))
		copy(imageQueue, imageNames)
	}

	img, idx := randImg(imageQueue)
	imageQueue = remove(imageQueue, idx)

	c.File(imagePath + "/" + img)
}

func ImageScanner() {
	fmt.Println("Scanning for images...")

	imagePath := viper.GetString("Image.Path")
	files, err := ioutil.ReadDir(imagePath)

	if err != nil {
		panic(err)
	}

	if len(files) == len(imageNames) {
		return
	}

	for _, file := range files {
		// if the image is not in the map, add it
		if !contains(imageNames, file.Name()) {
			imageNames = append(imageNames, file.Name())
		}
	}

	fmt.Println("Found " + strconv.Itoa(len(imageNames)) + " images")

	time.AfterFunc(time.Duration(viper.GetInt("Image.Scan.Interval"))*time.Minute, ImageScanner)
}

func randImg(arr []string) (string, int) {
	idx := rand.Intn(len(arr))
	return arr[idx], idx
}

func minIdx(arr []int) int {
	idx := 0
	m := arr[idx]

	for k, v := range arr {
		if m > v {
			m = v
			idx = k
		}
	}
	return idx
}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
