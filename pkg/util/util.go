package util

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// ToJson converts a struct to json
func ToJson(p interface{}) []byte {
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b
}

func StudyStart(year int) string {
	actyear := strconv.Itoa(year)

	front := actyear[:4]
	back := actyear[len(actyear)-1:]

	list_year := []int{20202, 20221, 20212}
	for _, x := range list_year {
		if x == year {
			back = "1"
			break
		}
	}

	f, _ := strconv.Atoi(front)
	b, _ := strconv.Atoi(back)

	new_year := f - 7
	fix_year := fmt.Sprintf("%v%v", new_year, b)

	return fix_year
}
