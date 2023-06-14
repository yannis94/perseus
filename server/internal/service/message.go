package service

type Message struct {
    Id int
    Room_id int
    User_id int
    Username string
    Content string
}

func NewMessageService() *Message {
    return &Message{}
}

func (m *Message) New(user_id int, username, content string) error {
    //add message to db
    return nil
}
