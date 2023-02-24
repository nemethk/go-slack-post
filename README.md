# go-slack-post
## USAGE
### Varriables
| Tag   | Details |
|:--------:|:-------------:|
| `-slack-token` |  Bearer token. |
| `-channel-id` | The selected channel ID. |
| `-message` | Message in Markdown format. |
### Run
```
go run main.go \
    -slack-token $STOKEN \
    -channel-id $SCHANNEL \
    -message "This is a mrkdwn section block :ghost: *this is bold*, and ~this is crossed out~, <https://google.com|this is a link>"
```
### JSON
```
{
    "channel": "ASDF1234",
    "blocks": [
        {
        "type": "section",
        "text": {
            "type": "mrkdwn",
            "text": "This is a mrkdwn section block :ghost: *this is bold*, and ~this is crossed out~, <https://google.com|this is a link>"
            }
        }
    ]
}
```