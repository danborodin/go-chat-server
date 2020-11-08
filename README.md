# gochatserver

User struct ->

User { ID primitive.ObjectID, Username string, Salt string, Password string}

Message struct ->

Message {ID int, Sender User, Text string, Date string }

Room object ->

Room {ID int, Name string, Owner string, Messages []Message}

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
