package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

var (
    BEARER_TOKEN string = ""
    CHANNEL_ID string = ""
) 

func main() {

    // required
    ptrBearerToken := flag.String("slack-token", "", "Slack Bearer Token")
    ptrChannelId := flag.String("channel-id", "", "Slack Channel ID")

    flag.Parse()
    BEARER_TOKEN = *ptrBearerToken
    CHANNEL_ID = *ptrChannelId

    type Data struct {
        Channel string `json:"channel"`
        Blocks []struct {
            Type string `json:"type"`
            Text struct {
                Type string `json:"type"`
                Text string `json:"text"`
            } `json:"text"`
        } `json:"blocks"`
    }

    data := Data{
        Channel: CHANNEL_ID,
        Blocks: []struct{ 
            Type: "section", 
            Text: struct{ 
                Type: "mrkdwn", 
                Text: "New IP is: " + "\"" + NEW_IP + "\"",
            }, 
        }, 
    }

    jsonData, _ := json.Marshal(data)

    url := "https://slack.com/api/chat.postMessage"
    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    req.Header.Add("Content-Type", "application/json; charset=utf-8")
    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", BEARER_TOKEN))
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
    println(string(body))
}
