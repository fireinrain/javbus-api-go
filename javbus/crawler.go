package javbus

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"sync"
)

// Helper functions

func formatImageUrl(url string) string {
	// Replace with your implementation if needed

	return url
}

func bytesToNumberSize(size string) int64 {
	if size == "" {
		return 0
	}

	bytes, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		return 0
	}

	return bytes
}

//func parseMoviesPage(pageHTML string, filter func(movie Movie) bool) MoviesPage {
//	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageHTML))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var movies []Movie
//
//	doc.Find("#waterfall #waterfall .item").Each(func(i int, item *goquery.Selection) {
//		imgSrc, _ := item.Find(".photo-frame img").Attr("src")
//		img := formatImageUrl(imgSrc)
//		title := item.Find(".photo-frame img").AttrOr("title", "")
//		info := item.Find(".photo-info date")
//		id := info.Eq(0).Text()
//		date := info.Eq(1).Text()
//		tags := item.Find(".item-tag button").Map(func(_ int, tag *goquery.Selection) string {
//			return tag.Text()
//		})
//
//		movies = append(movies, Movie{
//			Date:  date,
//			ID:    id,
//			Img:   img,
//			Title: title,
//			Tags:  tags,
//		})
//	})
//
//	currentPage, _ := strconv.Atoi(doc.Find(".pagination .active a").Text())
//	pages := doc.Find(".pagination li a").Map(func(_ int, page *goquery.Selection) int {
//		pageNumber, _ := strconv.Atoi(page.Text())
//		return pageNumber
//	})
//	pages = filterPages(pages)
//
//	hasNextPage := doc.Find(".pagination li #next").Length() > 0
//	nextPage := currentPage + 1
//
//	return MoviesPage{
//		Movies:     filterMovies(movies, filter),
//		Pagination: Pagination{currentPage, hasNextPage, nextPage, pages},
//	}
//}

//func filterPages(pages []int) []int {
//	var result []int
//	pageReg := regexp.MustCompile(PAGE_REG)
//	for _, page := range pages {
//		if pageReg.MatchString(strconv.Itoa(page)) {
//			result = append(result, page)
//		}
//	}
//	return result
//}

//func filterMovies(movies []Movie, filter func(movie Movie) bool) []Movie {
//	var result []Movie
//	for _, movie := range movies {
//		if filter == nil || filter(movie) {
//			result = append(result, movie)
//		}
//	}
//	return result
//}

//func parseStarInfo(pageHTML string, starId string) StarInfo {
//	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageHTML))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	avatarSrc, _ := doc.Find("#waterfall .item .avatar-box .photo-frame img").Attr("src")
//	avatar := formatImageUrl(avatarSrc)
//	name := doc.Find("#waterfall .item .avatar-box .photo-info .pb10").Text()
//
//	infos := doc.Find("#waterfall .item .avatar-box .photo-info p")
//
//	var rest = make(map[string]string)
//	mapKeys := []string{"birthday", "age", "height", "bust", "waistline", "hipline", "birthplace", "hobby"}
//
//	infos.Each(func(i int, s *goquery.Selection) {
//		mapValue := starInfoMap[mapKeys[i]]
//		value := s.Text()
//		rest[mapKeys[i]] = strings.TrimPrefix(value, mapValue)
//	})
//
//	return StarInfo{
//		Avatar:     avatar,
//		ID:         starId,
//		Name:       name,
//		Birthday:   rest["birthday"],
//		Age:        rest["age"],
//		Height:     rest["height"],
//		Bust:       rest["bust"],
//		Waistline:  rest["waistline"],
//		Hipline:    rest["hipline"],
//		Birthplace: rest["birthplace"],
//		Hobby:      rest["hobby"],
//	}
//}

// ParseStarInfo
//
//	@Description: Parse star information
//	@param pageHTML
//	@param starID
//	@return StarInfo
//	@return error
func ParseStarInfo(pageHTML string, starID string) (StarInfo, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageHTML))
	if err != nil {
		return StarInfo{}, err
	}

	avatar := formatImageUrl(doc.Find("#waterfall .item .avatar-box .photo-frame img").AttrOr("src", ""))
	name := doc.Find("#waterfall .item .avatar-box .photo-info .pb10").Text()

	infos := doc.Find("#waterfall .item .avatar-box .photo-info p")

	rest := make(map[string]string)
	for key, mapValue := range starInfoMap {
		value := infos.FilterFunction(func(i int, s *goquery.Selection) bool {
			return strings.Contains(s.Text(), mapValue)
		}).Text()
		value = strings.ReplaceAll(value, mapValue, "")
		rest[key] = value
	}
	ageStr := strings.TrimSpace(rest["age"])
	if strings.HasSuffix(ageStr, "cm") {
		ageStr = strings.ReplaceAll(ageStr, "cm", "")
	}
	ageInt, err := strconv.Atoi(ageStr)

	if err != nil {
		fmt.Println("Parse age str error:", err)
		return StarInfo{}, err
	}

	heightStr := strings.TrimSpace(rest["height"])
	if strings.HasSuffix(heightStr, "cm") {
		heightStr = strings.ReplaceAll(heightStr, "cm", "")
	}
	heightInt, err := strconv.Atoi(heightStr)

	if err != nil {
		fmt.Println("Parse height str error:", err)
		return StarInfo{}, err
	}

	return StarInfo{
		Avatar:     avatar,
		ID:         starID,
		Name:       name,
		Birthday:   rest["birthday"],
		Age:        ageInt,
		Height:     heightInt,
		Bust:       rest["bust"],
		Waistline:  rest["waistline"],
		Hipline:    rest["hipline"],
		Birthplace: rest["birthplace"],
		Hobby:      rest["hobby"],
	}, nil
}

// GetStarInfo
//
//	@Description: Get star information
//	@param starID
//	@param movieType
//	@return StarInfo
//	@return error
func GetStarInfo(starID string, movieType MovieType) (StarInfo, error) {
	var prefix string

	JAVBUS := JavbusBaseUrl

	if movieType == "" || movieType == "normal" {
		prefix = JAVBUS
	} else {
		prefix = fmt.Sprintf("%s/%s", JAVBUS, movieType)
	}
	url := fmt.Sprintf("%s/star/%s", prefix, starID)

	// Create a channel to receive the result
	resultCh := make(chan StarInfo, 1)
	errorCh := make(chan error, 1)

	// Use a wait group to wait for the goroutine to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Launch a goroutine to fetch the star information
	go func() {
		defer wg.Done()
		var respText string
		err := HttpClient.BaseURL(url).ToString(&respText).Fetch(context.Background())
		if err != nil {
			errorCh <- err
			return
		}

		starInfo, err := ParseStarInfo(respText, starID)
		if err != nil {
			errorCh <- err
			return
		}

		resultCh <- starInfo
	}()

	// Use another goroutine to wait for either result or an error
	go func() {
		wg.Wait()
		close(resultCh)
		close(errorCh)
	}()

	// Wait for either result or an error
	select {
	case result := <-resultCh:
		return result, nil
	case err := <-errorCh:
		return StarInfo{}, err
	default:
		//do nothing for not dead lock
	}
	return StarInfo{}, errors.New("get star info failed")
}
