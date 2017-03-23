package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"os"
)

const testTag = "yu81-22"

var (
	okASINCases = []string{"B00BXVR8FU", "4873116309", "B014US3FQI"}
	ngASINCases = []string{"999988888887", "C00BXVR8FU"}
	okURLCases  = []string{
		"https://www.amazon.co.jp/%E3%82%BD%E3%83%8B%E3%83%BC-SONY-DSC-RX100-F1-8%E3%83%AC%E3%83%B3%E3%82%BA%E6%90%AD%E8%BC%89-Cyber-shot/dp/B00898JY8E/ref=s9_ri_gw_g421_i8_r?pf_rd_m=AN1VRQENFRJN5&pf_rd_s=&pf_rd_r=G577NXCT8G8FS3E10VDD&pf_rd_t=36701&pf_rd_p=005e9dd3-c918-464e-81c6-8eb6da846ee5&pf_rd_i=desktop",
		"https://www.amazon.co.jp/gp/product/B00KR9ML5G/ref=br_asw_pdt-4?pf_rd_m=AN1VRQENFRJN5&pf_rd_s=&pf_rd_r=G577NXCT8G8FS3E10VDD&pf_rd_t=36701&pf_rd_p=3f546ffb-f3ae-45ce-a712-15f1d8607693&pf_rd_i=desktop",
	}
)

func TestCreateAmazonJpURLFromASIN(t *testing.T) {
	for _, s := range okASINCases {
		assert.Equal(t, amazonDpBase+s, CreateAmazonJpURLFromASIN(s))
	}
	for _, s := range ngASINCases {
		assert.Equal(t, amazonDpBase+s, CreateAmazonJpURLFromASIN(s))
	}
}

func TestCreateAmazonJpURLFromASINWithAffiliate(t *testing.T) {
	for _, s := range okASINCases {
		assert.Equal(t, amazonDpBase+s+"/?tag="+testTag, CreateAmazonJpURLFromASINWithAffiliate(s, testTag))
	}
	for _, s := range ngASINCases {
		assert.Equal(t, amazonDpBase+s+"/?tag="+testTag, CreateAmazonJpURLFromASINWithAffiliate(s, testTag))
	}
}

func TestExtractASIN(t *testing.T) {
	for _, s := range okASINCases {
		assert.Equal(t, amazonDpBase+s, CreateAmazonJpURLFromASIN(s))
	}
	for _, s := range ngASINCases {
		assert.Equal(t, amazonDpBase+s, CreateAmazonJpURLFromASIN(s))
	}
}

func TestIsASIN(t *testing.T) {
	for _, s := range okASINCases {
		assert.True(t, IsASIN(s), s+" is not ASIN.")
	}
	for _, s := range ngASINCases {
		assert.False(t, IsASIN(s), s+" is ASIN.")

	}

}

func TestToSimpleAmazonLink(t *testing.T) {
	for _, s := range okURLCases {
		assert.Equal(t, amazonDpBase+ExtractASIN(s), ToSimpleAmazonLink(s))
	}
}

func TestToSimpleAmazonLinkWithAffiletate(t *testing.T) {
	for _, s := range okURLCases {
		assert.Equal(t, amazonDpBase+ExtractASIN(s)+"/?tag="+testTag, ToSimpleAmazonLinkWithAffiliate(s, testTag))
	}
}

func TestShortenURLWithBitly(t *testing.T) {
	apiKey := os.Getenv("BITLY_API_KEY")
	clientSecret := os.Getenv("BITLY_CLIENT_SECRET")
	shorten, err := ShortenURLWithBitly("https://www.google.co.jp", apiKey, "" , clientSecret)
	assert.NoError(t, err)
	assert.NotEmpty(t, shorten)
}
