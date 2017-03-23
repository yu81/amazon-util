package url

import (
	"encoding/json"
	"os"
	"regexp"

	"github.com/jphines/bitly-api-go"
)

const amazonDpBase = "https://www.amazon.co.jp/dp/"

var ASINRegex = regexp.MustCompile("(B[A-Z0-9]{9})|([0-9]{10})")

func IsASIN(id string) bool {
	r := ASINRegex.Copy()
	return r.MatchString(id) && len(id) == 10
}

func ExtractASIN(u string) string {
	r := ASINRegex.Copy()
	matched := r.FindAllStringSubmatch(u, 1)
	if len(matched) > 0 && len(matched[0]) > 0 && IsASIN(matched[0][0]) {
		return matched[0][0]
	}
	return ""
}

func CreateAmazonJpURLFromASIN(id string) string {
	return amazonDpBase + id
}

func CreateAmazonJpURLFromASINWithAffiliate(id, tag string) string {
	return CreateAmazonJpURLFromASIN(id) + "/?tag=" + tag
}

func ToSimpleAmazonLink(link string) string {
	return CreateAmazonJpURLFromASIN(ExtractASIN(link))
}

func ToSimpleAmazonLinkWithAffiliate(link, tag string) string {
	return CreateAmazonJpURLFromASINWithAffiliate(ExtractASIN(link), tag)
}

func ShortenURLWithBitly(link, apiKey, token, secret, login string) (string, error) {
	connection := bitly_api.NewConnectionOauth(token, "", apiKey, login, secret)
	shorten, err := connection.Shorten(link)
	if err != nil {
		return "", err
	}

	return NewBitlyShortenResponseFromMap(shorten).URL, nil
}

func GetBitlyCredentials() BitlyCredential {
	return BitlyCredential{
		APIKey:       os.Getenv("BITLY_API_KEY"),
		ClientSecret: os.Getenv("BITLY_CLIENT_SECRET"),
		Login:        os.Getenv("BITLY_LOGIN_USER"),
		ClientID:     os.Getenv("BITLY_API_CLIENT_ID"),
	}
}

type BitlyCredential struct {
	APIKey       string
	ClientSecret string
	ClientID     string
	Login        string // bitly username.
}

type BitlyShortenResponse struct {
	GlobalHash string `json:"global_hash"`
	Hash       string `json:"hash"`
	LongURL    string `json:"long_url"`
	NewHash    int    `json:"new_hash"`
	URL        string `json:"url"`
}

func NewBitlyShortenResponseFromMap(m map[string]interface{}) BitlyShortenResponse {
	response := BitlyShortenResponse{}
	j, err := json.Marshal(m)
	if err != nil {
		return response
	}

	json.Unmarshal(j, &response)
	return response
}
