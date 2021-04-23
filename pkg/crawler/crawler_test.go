package crawler

import "testing"

//crawler unit test
func TestCrawler(testT *testing.T) {

	topResult := []string{}
	err := Run("https://www.alexa.com/topsites/",
		".tr.site-listing div.td",
		".right",
		true,
		&ExampleStruct{TopNumber: 10, TopResult: &topResult},
	)
	if err != nil {
		testT.Errorf(`項目測試錯誤 %v`, err)
	}

	if len(topResult) < 1 {
		testT.Errorf("爬蟲測試結果異常:\n%v\n需要檢查網站是否遭受異動！", topResult)
	}

}
