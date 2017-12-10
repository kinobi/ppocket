package pocket

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	apiRequestToken = "https://getpocket.com/v3/oauth/request"
	apiUserGrants   = "https://getpocket.com/auth/authorize"
	apiAuthorize    = "https://getpocket.com/v3/oauth/authorize"
)

// OAuthProcess proceed to the OAuth signup
func OAuthProcess(ppocketConsumerKey string, urlCallback string) (ppocketUserAccessToken, ppocketUsername string) {
	data := url.Values{}
	data.Set("consumer_key", ppocketConsumerKey)
	data.Set("redirect_uri", urlCallback)

	m, err := callOauthAPI(apiRequestToken, data, "request token")
	if err != nil {
		log.Fatal(err)
	}
	code := m.Get("code")

	fmt.Printf("Please visit %s?request_token=%s&redirect_uri=%s\n", apiUserGrants, code, urlCallback)
	fmt.Println("When it is done, press Y then Enter:")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		validation := strings.ToUpper(s.Text())
		if validation == "Y" {
			break
		}
		fmt.Println("When it is done, press Y then Enter:")
	}
	if err := s.Err(); err != nil {
		log.Fatalf("Failed to read standard input: %s", err)
	}

	data = url.Values{}
	data.Set("consumer_key", ppocketConsumerKey)
	data.Set("code", code)
	m, err = callOauthAPI(apiAuthorize, data, "access token")
	if err != nil {
		log.Fatal(err)
	}
	ppocketUserAccessToken = m.Get("access_token")
	ppocketUsername = m.Get("username")

	log.Printf("User: %s", ppocketUsername)
	log.Printf("Access Token: %s", ppocketUserAccessToken)

	return ppocketUserAccessToken, ppocketUsername
}

func callOauthAPI(endpoint string, query url.Values, step string) (data url.Values, err error) {
	res, err := http.PostForm(endpoint, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to get %s: %v", step, err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to get %s: %s [%s]", step, res.Header.Get("X-Error"), res.Header.Get("X-Error-code"))
	}

	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read %s: %v", step, err)
	}
	m, err := url.ParseQuery(string(buf))
	if err != nil {
		return nil, fmt.Errorf("Failed to parse %s: %v", step, err)
	}

	return m, nil
}
