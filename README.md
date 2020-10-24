# gochatserver

User struct ->

User { ID int, Username string, Password string, Nickname string }

Message struct ->

Message {ID int, Sender User, Text string, Date string }

Channel struct ->

Channel {ID int, Name string, Users []User, Messages []Message}
