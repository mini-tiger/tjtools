package control

type Semaphore struct {
	bufSize int
	channel chan struct{}
}

func NewSemaphore(concurrencyNum int) *Semaphore {
	return &Semaphore{channel: make(chan struct{}, concurrencyNum), bufSize: concurrencyNum}
}

func (s *Semaphore) TryAcquire() bool {
	select {
	case s.channel <- struct{}{}:
		return true
	default:
		return false
	}
}

func (s *Semaphore) Acquire() {
	s.channel <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.channel
}

func (s *Semaphore) AvailablePermits() int {
	return s.bufSize - len(s.channel)
}
