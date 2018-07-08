package domain

type AssigmentList struct {
	List []Assigment
}

func (this *AssigmentList) Run() error {
	for _, assigenment := range this.List {
		assigenment.Run()
	}

	return nil
}
