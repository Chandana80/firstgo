package service

import (
	"das-go/entity"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strings"
)


func PrepareAuthHeader(iusToken *entity.IusResponse, realmId string) string{
	auth := `Intuit_IAM_Authentication intuit_userid=IamUserId, intuit_token=IamToken, intuit_realmid=IamRealmId, intuit_appid=Intuit.platform.dc.automation,intuit_app_secret=intuitAppSecret`
	rep := strings.NewReplacer("IamUserId", iusToken.IamTicket.UserId,
		"IamToken", iusToken.IamTicket.Ticket,
		"IamRealmId", realmId,)

	authHeader := rep.Replace(auth)
	return authHeader
}

func GetCredentials(authHeader string, realmId string, credentialId string) *entity.CredentialSet{
	var credentialSet entity.CredentialSet

	uuid, err := uuid.NewV4()
	request, err := http.NewRequest("GET", "https://credsurl/v2/realms/"+realmId+"/credentials/"+credentialId, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("intuit_assetid", "4901059886677171085")
	request.Header.Set("intuit_country", "US")
	request.Header.Set("intuit_iddomain", "global")
	request.Header.Set("intuit_locale", "EN")
	request.Header.Set("intuit_originating_assetid", "4901059886677171085")
	request.Header.Set("intuit_offeringid", "8")
	request.Header.Set("intuit_originatingip", "127.0.0.1")
	request.Header.Set("intuit_tid", uuid.String())
	request.Header.Set("Authorization", authHeader)
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
		fmt.Printf("Credential Service Response: %s\n", string(contents))
		json.Unmarshal([]byte(string(contents)), &credentialSet)

		return &credentialSet

	}
	return &credentialSet
}