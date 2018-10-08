package service

import (
	"bytes"
	"das-go/entity"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
)

func GetIUSToken(username string , password string) *entity.IusResponse{
	var iusToken entity.IusResponse

	values := map[string]string{"username": username, "password": password}
	requestBody, _ := json.Marshal(values)

	uuid, err := uuid.NewV4()
	request, err := http.NewRequest("POST", "https://iusUrl", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("intuit_offeringid", "Intuit.platform.dc.automation")
	request.Header.Set("intuit_originatingip", "127.0.0.1")
	request.Header.Set("intuit_tid", uuid.String())
	request.Header.Set("Authorization", "Intuit_IAM_Authentication intuit_appid=Intuit.platform.dc.automation,intuit_app_secret=intuitAppSecret")
	request.Header.Set("Accept", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Printf("IUS Response: %s\n", string(contents))
		json.Unmarshal([]byte(string(contents)), &iusToken)

		return &iusToken
	}
	return &iusToken
}
