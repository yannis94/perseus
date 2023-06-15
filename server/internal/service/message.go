package service

func NewMessageService() *MessageService {
    return &MessageService{}
}

func (m *MessageService) New(user_id int, username, content string) error {
    //add message to db
    return nil
}
