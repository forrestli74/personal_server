# curl --insecure \
#      --request "POST" \
#      --location "https://localhost:8080/twirp/tmp.RoomService/CreateRoom" \
#      --header "Content-Type:application/json" \
#      --data '{"room_id": "room"}' \
#      --verbose
# curl --insecure \
#      --request "POST" \
#      --location "https://localhost:8080/twirp/tmp.RoomService/AddWriter" \
#      --header "Content-Type:application/json" \
#      --data '{"room_id": "room", "proposed_ids": ["a"]}' \
#      --verbose
curl --insecure \
     --request "POST" \
     --location "https://localhost:8080/twirp/tmp.RoomService/DeleteRoom" \
     --header "Content-Type:application/json" \
     --data '{"room_id": "room"}' \
     --verbose
