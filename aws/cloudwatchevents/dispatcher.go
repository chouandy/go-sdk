package cloudwatchevents

import "errors"

type Context struct {
	Name   string                 `json:"name"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type HandlerFunc func(ctx *Context) error

type Event struct {
	Handler     HandlerFunc
	Middlewares []HandlerFunc
}

type Dispatcher struct {
	Events map[string]*Event
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		Events: make(map[string]*Event),
	}
}

func (d *Dispatcher) Add(name string, event *Event) {
	d.Events[name] = event
}

func (d *Dispatcher) Get(name string) *Event {
	event, ok := d.Events[name]
	if !ok {
		return &Event{
			Handler: NotFoundHandler,
		}
	}
	return event
}

func (d *Dispatcher) Dispatch(ctx *Context) error {
	event := d.Get(ctx.Name)

	for _, middleware := range event.Middlewares {
		if err := middleware(ctx); err != nil {
			return err
		}
	}

	return event.Handler(ctx)
}

func NotFoundHandler(ctx *Context) error {
	return errors.New("event not found")
}
