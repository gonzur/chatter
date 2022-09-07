package chatrooms

type Room struct {
	name string
	// should eventually hold a set of permissions not bool
	members map[*Member]bool
	join    chan *Member
	leave   chan *Member
	cast    chan Message
}

func (r *Room) Init(roomName string) {
	r.name = roomName
	r.members = make(map[*Member]bool)
	r.join = make(chan *Member)
	r.leave = make(chan *Member)
	r.cast = make(chan Message)
	go r.Run()
}

func (r *Room) Run() {
	for {
		select {
		case member := <-r.join:
			r.members[member] = true

		case member := <-r.leave:
			goodbye := makeMessage("", "Goodbye")
			r.cast <- goodbye
			delete(r.members, member)

		case message := <-r.cast:
			for m := range r.members {
				if m.ID != message.Sender {
					m.send <- message
				}

			}
		}
	}
}
