package crawler

import (
	"encoding/json"
	"io"
	"time"
	"net/http"
	"net/url"
	
)

const checkTor string = "https://check.torproject.org/api/ip"

var data struct {
    IsTor bool `json:"IsTor"`
}

func checkIfTorConnectionExists(torProxy string) (bool, error) {
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