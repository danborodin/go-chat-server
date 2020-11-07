# gochatserver

User struct ->

User { ID primitive.ObjectID, Username string, Salt string, Password string}

Message struct ->

Message {ID int, Sender User, Text string, Date string }

Channel struct ->

Channel {ID int, Name string, Users []User, Messages []Message}

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

http://172.104.202.219:8080/channels
GET

http://172.104.202.219:8080/channel
POST
{"name":""}

http://172.104.202.219:8080/channel/{id}
GET

ws://172.104.202.219:8080/channel/{id}
