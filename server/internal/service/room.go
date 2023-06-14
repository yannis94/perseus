package service

type Room struct {
    Id int
    Name string
    Capacity int
}

func NewRoomService() *Room {
    return &Room{}
}

func (r *Room) Create(name string, capacity int) (Room, error) {
    //fetch room from db
    return Room{}, nil
}

func (r *Room) Join(room_id int) (Room, error) {
    //fetch room by id from db
    return Room{}, nil
}

func (r *Room) Leave(room_id int) error {
    return nil
}
