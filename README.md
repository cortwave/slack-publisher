# slack-publisher
lightweight dockerized daemon for publishing messages to slack

## Build (optionally, needed to update ca-certificates)

``` bash
./build.sh 0.1.1
```

## Run

``` bash
docker run -e "TOKEN=your-slack-api-token" -e "CHANNEL=slack-channel" -d -p 8000:8000 cortwave/slack-publisher:0.1.1
```

## Use

``` bash
curl -d '{ "text" :"Hello!" }' localhost:8000/publish
```
