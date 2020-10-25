# gochatserver

User struct ->

User { ID primitive.ObjectID, Username string, Salt string, Password string}

Message struct ->

Message {ID int, Sender User, Text string, Date string }

Channel struct ->

Channel {ID int, Name string, Users []User, Messages []Message}
