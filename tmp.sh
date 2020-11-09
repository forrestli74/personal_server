URL="https://localhost"
# URL="http://localhost:8080"
curl --insecure \
     --request "POST" \
     --location "$URL/twirp/tmp.RoomService/DeleteRoom" \
     --header "Content-Type:application/json" \
     --data '{"room_id": "room"}'

curl --insecure \
     --request "POST" \
     --location "$URL/twirp/tmp.RoomService/CreateRoom" \
     --header "Content-Type:application/json" \
     --data '{"room_id": "room", "room_setting": { "tick_setting": {"size": 1, "frequency_millis":400 }} }'

# curl --insecure \
#      --request "POST" \
#      --location "https://localhost:8080/twirp/tmp.RoomService/AddWriter" \
#      --header "Content-Type:application/json" \
#      --data '{"room_id": "room", "proposed_ids": ["a"]}' \
#      --verbose

# curl --insecure \
#      --request "POST" \
#      --location "http://localhost:8080/twirp/tmp.RoomService/DeleteRoom" \
#      --header "Content-Type:application/json" \
#      --data '{"room_id": "room"}' \
#      --verbose
