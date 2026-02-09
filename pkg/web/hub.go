package web

type Hub struct {
	Browsers map[Browser]bool

	Broadcast chan Event

	Commands chan []byte

	Register chan *Browser

	Deregister chan *Browser

	VideoServer *VideoServer
}

func NewHub() Hub {
	return Hub {
		Browsers: make(map[Browser]bool),
		Broadcast: make(chan Event),
		Commands: make(chan []byte),
		Register: make(chan *Browser),
		Deregister: make(chan *Browser),
	}
}

func (h *Hub) RemoveBrowser(browser *Browser) {
	delete(h.Browsers, *browser)
	close(browser.Queue)
}

func (h *Hub) Listen() {
	for {
		select {
		case browser := <-h.Register:
			h.Browsers[*browser] = true

		case browser := <-h.Deregister:
			if _, ok := h.Browsers[*browser]; ok {
				h.RemoveBrowser(browser)
			}

		case msg := <-h.Broadcast:
			for browser := range h.Browsers {
				select {
				case browser.Queue <- msg:
				default:
					h.RemoveBrowser(&browser)
				}
			}
		}
	}
}
