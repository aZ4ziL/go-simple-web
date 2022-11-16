package handlers

// Flasher
// untuk pesan flash untuk pemberitahuan dan alert
type Flasher struct {
	Type    string
	Message string
}

// Set alert
func (f *Flasher) Set(_type, msg string) {
	f.Type = _type
	f.Message = msg
}

func (f *Flasher) Del() {
	f.Type = ""
	f.Message = ""
}
