package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, "https://www.amazon.co.jp/dp/"+s, CreateAmazonJpURLFromASIN(s))
	}
	for _, s := range ngASINCases {
		assert.Equal(t, "https://www.amazon.co.jp/dp/"+s, CreateAmazonJpURLFromASIN(s))
	}
}

func TestCreateAmazonJpURLFromASINWithAffiliate(t *testing.T) {
	for _, s := range okASINCases {
		assert.Equal(t, "https://www.amazon.co.jp/dp/"+s+"/?tag="+testTag, CreateAmazonJpURLFromASINWithAffiliate(s, testTag))
	}
	for _, s := range ngASINCases {
		assert.Equal(t, "https://www.amazon.co.jp/dp/"+s+"/?tag="+testTag, CreateAmazonJpURLFromASINWithAffiliate(s, testTag))
	}
}

func TestExtractASIN(t *testing.T) {
	for _, s := range okASINCases {
		assert.Equal(t, "https://www.amazon.co.jp/dp/"+s, CreateAmazonJpURLFromASIN(s))
	}
	for _, s := range ngASINCases {
		assert.Equal(t, "https://www.amazon.co.jp/dp/"+s, CreateAmazonJpURLFromASIN(s))
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
		assert.Equal(t, "https://www.amazon.co.jp/dp/"+ExtractASIN(s), ToSimpleAmazonLink(s))
	}
}

func TestToSimpleAmazonLinkWithAffiletate(t *testing.T) {
	for _, s := range okURLCases {
		assert.Equal(t, "https://www.amazon.co.jp/dp/"+ExtractASIN(s)+"/?tag="+testTag, ToSimpleAmazonLinkWithAffiliate(s, testTag))
	}
}
