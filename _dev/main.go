package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/s-vvardenfell/datetimeparser"
)

type TextWithDate []struct {
	DateCreatedRaw string `json:"date_created_raw"`
}

/*
- Доработки:
	* "возможные" доработки записать в readme
	* Mon Jan 02 15:04:05 -0700 2006  - в таком не парсится год
	* "сдвиг" строки при парсинге слова-месяца (марта-арта-рта и тд)
	* возможен рефакторинг функций поиска/парсинга цифр в строке - можно упростить и сделать более ре-юзабельную под-функцию
	* parseSeparatedDigits с 1月12日 23:20 не работает - видимо, из-за длинны руны
	* запоминание корректной позиции даты сейчас происходит только после второй итерации - нужно сделать корректный сдвиг; сдвиг сделал, но начало аффектить другие форматы
*/

var inputDates = []string{
	// "3/24/2016",
	// "31 minutes ago",
	// "four hours ago",
	// "31 minutes ago",
	// "four hours ago",
	// "21 minutes ago",
	// "Мар 12, 2012",
	"2006-01-02T15:04-0700",
	"2006-01-02T15:04-0700",
	"2006-01-02T15:04-0700",
	"2006-01-02T15:04-0700",

	"2021年06月29日09時20分",
	"2021年06月29日09時20分",
	"2021年06月29日09時20分",
	"2021年06月29日09時20分",

	"15:04:05 02-01-2006",
	"15:04:05 02-01-2006",
	"15:04:05 02-01-2006",
	"15:04:05 02-01-2006",
	"15:04:05 02-01-2006",

	"2007-01-02 15:04:05",
	"2007-01-02 15:04:05",
	"2007-01-02 15:04:05",
	"2007-01-02 15:04:05",
	"2007-01-02 15:04:05",

	// "1675929598",
	// "15:04:05",
}

func main() {
	var prs datetimeparser.TimeParser

	for i := range inputDates {
		// fmt.Println(inputDates[i])
		prs.Parse(inputDates[i])
		// fmt.Println(prs)
		fmt.Println(prs.ToUTCTime())
	}

	return

	//------------------

	data, err := getSimplifiedKibanaJsonFromFile(
		// "/home/user/Documents/tasksfiles/test_date_parsing/test_data/results_1M_since_01-01-2021.json") // 1004066/556
		// "/home/user/Documents/tasksfiles/test_date_parsing/test_data/results_1M_since_01-01-2022.json") // 1009292/546
		"/home/user/Documents/tasksfiles/test_date_parsing/test_data/results_1M_since_01-05-2022.json") // 1003848/78
	// "/home/user/Documents/tasksfiles/test_date_parsing/test_data/results_1M_since_01-08-2022.json") // 986198/869
	if err != nil {
		log.Fatal(err)
	}

	success := 0
	fail := 0
	start := time.Now()

	file, err := os.OpenFile(
		"/home/user/Documents/tasksfiles/test_date_parsing/results/results_1M_failed.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// file2, err := os.OpenFile(
	// 	"/home/user/Documents/tasksfiles/test_date_parsing/results/results_1M_successed.txt",
	// 	os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer file2.Close()

	for i := range data {
		if data[i].DateCreatedRaw == "" {
			continue
		}

		// fmt.Printf("Processing: <%s>\n", data[i].DateCreatedRaw)

		_, err := prs.Parse(data[i].DateCreatedRaw)
		if err != nil {
			_, err = file.WriteString(data[i].DateCreatedRaw + "\n")
			if err != nil {
				log.Fatal(err)
			}

			fail++
		} else {
			// fmt.Println(ret)

			// _, err = file2.WriteString(ret.String() + "\t" + data[i].DateCreatedRaw + "\n")
			// if err != nil {
			// 	log.Fatal(err)
			// }

			success++
		}

		fmt.Printf("successed: %d\tfailed: %d\n", success, fail)
	}

	fmt.Printf("Total processed: %d\n", success+fail)

	elapsed := time.Since(start)
	fmt.Printf("Converting took %s\n", elapsed)
}

func getSimplifiedKibanaJsonFromFile(infile string) (TextWithDate, error) {
	data, err := os.ReadFile(infile)
	if err != nil {
		return nil, err
	}

	var res TextWithDate

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
