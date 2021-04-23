package main

import (
	"alexa-rank-viewer/pkg/crawler"
	"alexa-rank-viewer/pkg/flags"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	if len(flags.Args) < 2 {
		log.Println("執行檔輸入參數不足至少需要2個！ 如：./cmd action 10")
		os.Exit(127) //	"command not found"
	}

	var err error
	topResult := make(map[int]string) //爬蟲結果輸出變數
	var topNum int                    //控制顯示結果排名數
	action := flags.Args[0]           //$ ./clawer <action> <arg1> [<arg2>...]
	switch action {
	case "top": // show top <number> sites URL on www.alexa.com/topsites/
		topNum, err = strconv.Atoi(flags.Args[1])
		if err != nil { //此狀況為排名數字異常
			log.Printf("執行檔輸入參數模式不正確： %v > %v模式下參數必須要為正整數.\n", flags.Args, action)
			os.Exit(127) //	"command not found"
		}

		err = crawler.Run("https://www.alexa.com/topsites/",
			".tr.site-listing div.td",
			".right",
			true,
			&TopMode{TopNumber: topNum, TopResult: &topResult},
		)
	case "country": //show top 20 sites URL on www.alexa.com/topsites/ by country
		topNum = 20
		err = crawler.Run("https://www.alexa.com/topsites/countries/"+flags.Args[1],
			".tr.site-listing div.td",
			".right",
			true,
			&TopMode{TopNumber: topNum, TopResult: &topResult},
		)

	default:
		log.Printf("執行檔輸入參數模式不正確： %v\n", flags.Args)
		os.Exit(127) //"command not found"
	}

	if err != nil {
		log.Printf("執行期間發生錯誤 > %v\n", err)
		os.Exit(1) //Catchall for general errors
	}

	fmt.Println("<< 以下為爬蟲結果 >>")
	for index := 0; index < topNum+1; index++ {
		if name, isExsit := topResult[index]; isExsit {
			fmt.Printf("排名 [%v] : %v\n", index, name)
		}
	} //排名有可能因爬蟲方式發生不存在狀況,故需要針對字串做出判斷

	fmt.Println("<< 以上為爬蟲結果 >>")
	os.Exit(0)

}

//以物件導向實踐排名方法，讓其他開發者可以任意使用該pkg自訂爬蟲檢索結果處理
//not require any modifications to the existing code base.
type TopMode struct {
	TopNumber int             //檢索前排名數
	TopResult *map[int]string //爬蟲檢索結果輸出變數
}

//對應crawler.Structure接口
func (strut *TopMode) ValueResult(result []string) (err error) { //使用者可依需求調整功能ㄋ

	topResult := make(map[int]string, strut.TopNumber) //map[排名]=網站名稱
	rank := -1                                         //給予初始值不可以和strconv.Atoi同為0
	for index, value := range result {
		if index >= (strut.TopNumber * 2) { //爬蟲結果僅需索取排名內
			break
		}

		if rank == 0 { //爬蟲排名數字異常故跳過！
			continue
		}

		if index%2 == 0 { //爬蟲結果偶數為排名
			var err error
			rank, err = strconv.Atoi(value)
			if err != nil { //此狀況為排名數字異常
				continue
			}
		} else {
			topResult[rank] = value
		}
	}

	*strut.TopResult = topResult //將處理結果送回struct
	return nil
}
