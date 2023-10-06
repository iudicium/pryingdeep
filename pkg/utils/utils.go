package utils

import (
	"encoding/json"
	"io"
	"time"
	"net/http"
	"net/url"
    

    "github.com/gocolly/colly/v2"
    "github.com/r00tk3y/prying-deep/models"
	
)

const checkTor string = "https://check.torproject.org/api/ip"

var data struct {
    IsTor bool `json:"IsTor"`
}


func CheckIfTorConnectionExists(torProxy string) (bool, error) {
    torProxyUrl, err := url.Parse(torProxy)
    if err != nil {
        return false, err
    }

    // Set up a custom HTTP transport to use the proxy and create the client
    torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
    client := &http.Client{Transport: torTransport, Timeout: time.Second * 5}

    // Make request
    resp, err := client.Get(checkTor)
    if err != nil {
        return false, err
    }
    defer resp.Body.Close()

    // Read response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return false, err
    }

    // Unmarshal the JSON response into the data struct
    if err := json.Unmarshal(body, &data); err != nil {
        return false, err
    }

    return data.IsTor, nil
}

func CreateMapFromValues(data map[string][]string) models.PropertyMap {
    resultMap := make(models.PropertyMap)

    for key, values := range data {
        if len(values) == 1 {
            resultMap[key] = values[0]
        } else {
            resultMap[key] = values
        }
    }

    return resultMap
}
func ConvertContextToPropertyMap(ctx colly.Context) models.PropertyMap {
    resultMap := make(models.PropertyMap)

    // Use ctx.ForEach to process the context values
    ctx.ForEach(func(k string, v interface{}) interface{} {
        resultMap[k] = v
        return nil
    })

    return resultMap
}
