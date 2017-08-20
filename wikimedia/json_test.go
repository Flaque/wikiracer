package wikimedia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testWithContinue = `
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

const testWithEmptyFirstAndNoContinue = `
{
  "batchcomplete": "",
  "query": {
    "pages": {
      "6678": {
        "pageid": 6678,
        "ns": 0,
        "title": "Cat"
      },
      "4269567": {
        "pageid": 4269567,
        "ns": 0,
        "title": "Dog",
        "links": [
          {
            "ns": 0,
            "title": "List of dog sports"
          },
          {
            "ns": 0,
            "title": "List of extinct dog breeds"
          }
        ]
      }
    }
  },
  "limits": {
    "links": 500
  }
}`

func TestGetContinue(t *testing.T) {
	continueString := getPlcontinueFromJSONBytes([]byte(testWithContinue))
	assert.Equal(t, continueString, `13926|0|Siberia`)
}

func TestGetContinueWhenThereIsNoContinue(t *testing.T) {
	continueString := getPlcontinueFromJSONBytes([]byte(testWithEmptyFirstAndNoContinue))
	assert.Equal(t, continueString, "")
}

func TestGetLinksFromJSONBytes(t *testing.T) {
	pages, err := getLinksFromJSONBytes([]byte(testWithContinue))
	assert.Nil(t, err)
	assert.NotEmpty(t, pages["Hyena"])
	assert.Equal(t, pages["Hyena"][0], "Sechuran fox")
}

func TestGetLinksFromJSONBytesWhenTheFirstIsEmpty(t *testing.T) {
	pages, err := getLinksFromJSONBytes([]byte(testWithEmptyFirstAndNoContinue))
	assert.Nil(t, err)
	assert.Empty(t, pages["Cat"])
	assert.NotEmpty(t, pages["Dog"])
	assert.Equal(t, pages["Dog"][0], "List of dog sports")
}
