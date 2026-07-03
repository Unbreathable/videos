package mutexbench

type CounterCommand int

const (
	CounterCommandIncrement CounterCommand = iota
)

type CounterWorker struct {
	counter     int64
	commandChan chan CounterCommand
}

func (c *CounterWorker) Init() {
	c.counter = 0
	c.commandChan = make(chan CounterCommand)

	go func() {
		for cmd := range c.commandChan {
			switch cmd {
			case CounterCommandIncrement:
				c.counter++
			}
		}
	}()
}

func (c *CounterWorker) Increment() {
	c.commandChan <- CounterCommandIncrement
}

func (c *CounterWorker) Get() int64 {
	return c.counter
}
