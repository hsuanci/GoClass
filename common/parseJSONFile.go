package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"praticeone/types"
)

var Role []types.Role

func ParseJSONFile() {
	jsonFile, err := os.Open("Sample/布袋戲資料.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &Role)
}
