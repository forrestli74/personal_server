# Personal Server
[![Build Status](https://travis-ci.org/lijiaqigreat/personal_server.svg?branch=master)](https://travis-ci.org/lijiaqigreat/personal_server)
[![](https://dockerbuildbadges.quelltext.eu/status.svg?organization=lijiaqigreat&repository=personal_server)](https://cloud.docker.com/repository/docker/lijiaqigreat/personal_server/builds)
##

## Notes for myself

1. 1-10 player wants to play a private game
2. 1-100 player wants to play a public game

curl --request "POST" --location "http://35.203.149.76:8080/twirp/tmp.RoomService/CreateRoom" --header "Content-Type:application/json" --data '{"room_id": "room", "room_setting": {tick_setting": {"size": 1, "frequency_millis": 2000}}}'


