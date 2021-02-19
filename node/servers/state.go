package servers

type State struct {
	Pending map[string][]string
}

func NewState() *State {
	return &State{Pending: make(map[string][]string)}

}

func (ns *State) AddPendingRequest(objectid string, value string) {
	ns.Pending[objectid] = append(ns.Pending[objectid], value)
}

func (ns *State) DeletePendingRequestByIdx(objectid string, idx int) {
	ns.Pending[objectid] = append(ns.Pending[objectid][:idx], ns.Pending[objectid][idx+1:]...)
}

func (ns *State) DeletePendingRequestByValue(objectid string, value string) {
	foundIdx := -1
	for i, v := range ns.Pending[objectid] {
		if v == value {
			foundIdx = i
			break
		}
	}
	if foundIdx != -1 {
		ns.Pending[objectid] = append(ns.Pending[objectid][:foundIdx], ns.Pending[objectid][foundIdx+1:]...)
	}

}
