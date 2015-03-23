package EventManager;

type Event struct {
	name string
	namespace string
	data *interface{}
}

func (event Event) Name () string {
	return event.name
}

func (event Event) Namespace () string {
	return event.namespace
}

func (event Event) Data () *interface{} {
	return event.data
}
