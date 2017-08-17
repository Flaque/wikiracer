package wikimedia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testJSON = `
{
	"continue": {
	  "plcontinue": "13926|0|Siberia",
	  "continue": "||"
	},
	"query": {
	  "pages": {
		"13926": {
		  "pageid": 13926,
		  "ns": 0,
		  "title": "Hyena",
		  "links": [
			{
			  "ns": 0,
			  "title": "Sechuran fox"
			},
			{
			  "ns": 0,
			  "title": "Second Sudanese Civil War"
			},
			{
			  "ns": 0,
			  "title": "Selous' mongoose"
			},
			{
			  "ns": 0,
			  "title": "Senegal"
			},
			{
			  "ns": 0,
			  "title": "Serakhs"
			},
			{
			  "ns": 0,
			  "title": "Serval"
			},
			{
			  "ns": 0,
			  "title": "Servaline genet"
			},
			{
			  "ns": 0,
			  "title": "Shakespeare"
			},
			{
			  "ns": 0,
			  "title": "Short-eared dog"
			},
			{
			  "ns": 0,
			  "title": "Short-tailed mongoose"
			}
		  ]
		},
		"4269567": {
		  "pageid": 4269567,
		  "ns": 0,
		  "title": "Dog"
		}
	  }
	}
  }`

func TestGetContinue(t *testing.T) {
	continueString := getPlcontinueFromJSONBytes([]byte(testJSON))
	assert.Equal(t, continueString, `13926|0|Siberia`)
}

func TestGetLinksFromJSONBytes(t *testing.T) {
	pages, err := getLinksFromJSONBytes([]byte(testJSON))
	assert.Nil(t, err)
	assert.NotEmpty(t, pages["Hyena"])
}
