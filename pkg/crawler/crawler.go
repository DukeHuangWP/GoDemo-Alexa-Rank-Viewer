package crawler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

/*
執行爬蟲並將結果輸出到自訂變數
@param url: 爬蟲目標網址
@param query: 爬蟲html檢索元素語法（煩參考github.com/PuerkitoBio/goquery）
@param notQuery:  爬蟲html檢索排除元素語法（煩參考github.com/PuerkitoBio/goquery）
@param trimSpace:  爬蟲結果是否要去除左右贅字(\n \t 空白)
@param value:  負責對爬蟲結果輸出到變數，為自訂型別，可OOP繼承與方法重載
@return error : 錯誤提示
*/
func Run(url string, query string, notQuery string, trimSpace bool, value Structure) (err error) {
	result, err := QueryUrl(url, query, notQuery, trimSpace)
	if err != nil {
		return err
	}

	err = value.ValueResult(result)
	if err != nil {
		return err
	}

	return nil
}

/*
執行爬蟲並將結果輸出成切片格式
@param url: 爬蟲目標網址
@param query: 爬蟲html檢索元素語法（煩參考github.com/PuerkitoBio/goquery）
@param notQuery:  爬蟲html檢索排除元素語法（煩參考github.com/PuerkitoBio/goquery）
@param trimSpace:  爬蟲結果是否要去除贅字(\n \t 空白)
@return result : 爬蟲結果
@return error : 錯誤提示
*/
func QueryUrl(url string, query string, notQuery string, trimSpace bool) (result []string, err error) {
	response, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return result, fmt.Errorf("status code error: %d %s", response.StatusCode, response.Status)
	}

	btyes, err := ioutil.ReadAll(response.Body)
	reader := bytes.NewReader(btyes)
	dom, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return result, err
	}
	dom.Find(query).Not(notQuery).Each(func(_ int, selection *goquery.Selection) {
		if trimSpace {
			result = append(result, strings.TrimSpace(selection.Text()))
		} else {
			result = append(result, selection.Text())
		}

	})
	return result, nil
}
