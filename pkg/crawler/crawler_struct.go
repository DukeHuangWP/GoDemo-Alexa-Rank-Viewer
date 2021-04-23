package crawler

//爬蟲檢索結果處理接口
//使用者可以自訂型別,讓使用者可藉由OOP繼承與方法重載擴增功能
type Structure interface {
	ValueResult(result []string) error //func設為全域，提供使用者可任意對介面繼承、方法重載
}

//排名方法
type ExampleStruct struct {
	TopNumber int       //檢索前排名數
	TopResult *[]string //爬蟲檢索結果輸出變數
}

//對應Structure接口
func (strut *ExampleStruct) ValueResult(result []string) (err error) { //使用者可依需求調整功能
	number := strut.TopNumber * 2
	if len(result) < number { //top不可以超過爬蟲排名上限
		number = len(result)
	}
	*strut.TopResult = result[:number] //將處理結果送回struct
	return nil
}
