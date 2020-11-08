# gochatserver

User struct ->

User { _id primitive.ObjectID, username string, salt string, password string}

Message struct ->

Message {_id int, sender string, text string, date string }

Room object ->

Room {_id int, name string, owner string, messages []Message}

http://172.104.202.219:8080/register
POST
{
  "username": "",
  "password": ""
}

http://172.104.202.219:8080/login
POST
{
  "username": "",
  "password": ""
}

http://172.104.202.219:8080/rooms
GET

http://172.104.202.219:8080/room
POST
{"name":""}

http://172.104.202.219:8080/room/{id}
GET

ws://172.104.202.219:8080/room/{id}
