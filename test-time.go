package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// type ChaincodeData struct {
// 	Id        string `json:"id"`
// 	MSPId     string `json:"mspID"`
// 	Cert      string `json:"cert"`
// 	Issuer    string `json:"issuer"`
// 	Subject   string `json:"subject"`
// 	Signature string `json:"signature"`
// 	// Attributes *attrmgr.Attributes `json:"attributes"`
// 	CN string `json:"cn"`
// 	OU string `json:"ou"`
// }

type Data struct {
	Id                 int       `json:"id"`
	StartTime          time.Time `json:"startTime"`
	StartTimeTimestamp int64     `json:"timestamp"`
}

func main() {
	date := time.Now()

	d := Data{
		Id:                 1,
		StartTime:          date,
		StartTimeTimestamp: date.Unix(),
	}

	str, _ := json.Marshal(d)

	fmt.Println(string(str))

}
