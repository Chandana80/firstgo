package connection

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Agentconnection() string {
	fmt.Println("Starting the application...")

	requestBody := `{
  "providerId": "providerId",
  "reportingParams": {
    "clientApp": "clientapp",
    "clientSku": "clientsku",
    "clientVersion": "snapshot",
    "partnerId": "testpartnerid",
    "authId": "123456789",
    "alternateId": "alternateid",
    "eroId": "eroid"
  },
  "credentialSet": {
    "alternateIds": [],
    "credentials": [
      {
        "authenticationFieldId": "aad4f3a1-8f45-4d99-9657-389589a8bdb3",
        "authenticationFieldType": "OAuthProxyToken",
        "authenticationFieldValue": "oAuthProxyToken",
        "certVersion": "v01_credential_service_cryption_nonprod_key.corp.intuit.net",
        "encrypted": false
      }
    ],
    "participatingOfferingIds": []
  },
  "executionParams": {
    "dateRange": {
       "startDate": "2018-07-11T23:10:23Z",
      "endDate": "2018-08-15T23:45:59Z"
    },
    "aggregationMode": "REALTIME",
    "authId": "75011180085",
    "agentAuth": "native"
  },
  "entities": [
    {
      "entityName": "ACCOUNT",
      "dataLevel": [
        "summary",
        "details",
        "transactions"
      ]
    }
  ]
}`
	req, err := http.NewRequest("POST", "http://localhost:8080/wsi/v2/acquire", strings.NewReader(requestBody))

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("intuit_tid", "qa-DAS-automation-201808060012")
	req.Header.Set("intuit_providerid", "providerId")
	req.Header.Set("intuit_offeringid", "45")
	req.Header.Set("intuit_appid", "wsi")
	req.Header.Set("Authorization", "Intuit_IAM_Authentication intuit_token_type=IAM-Ticket,intuit_token=intuit_token,intuit_userid=intuit_useId,intuit_appid=wsi,intuit_app_secret=test")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("intuit_fds_agent_config_url", "https://financialprovider-e2e.platform.intuit.com/v1/providers/providerId/channels/webService/channelId?agentType=WebserviceAgent")
	req.Header.Set("intuit_app_secret", "test")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("RAW_HOST_DATA_ENABLED", "TRUE")
	req.Header.Set("intuit_fdp_usecase", "Oauth")
	req.Header.Set("intuit_fdp_flowname", "OauthMigration")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err.Error()
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		return string(data)
	}
}
