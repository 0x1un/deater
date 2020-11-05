package agimpl

import "time"

type TimingTasker interface {
	Time() time.Duration
}
