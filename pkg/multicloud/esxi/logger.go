
package esxi

import (
	"github.com/vmware/govmomi/vim25/progress"

	"yunion.io/x/log"
)

type logger struct {
	name string
	chid chan progress.Report
	over chan struct{}
}

func (l *logger) Sink() chan<- progress.Report {
	return l.chid
}

func (l *logger) End() {
	close(l.over)
}

func newLeaseLogger(name string, cap int) *logger {
	return &logger{
		name: name,
		chid: make(chan progress.Report, cap),
		over: make(chan struct{}),
	}
}

func (l *logger) Log() {
	go func() {
		var pre float32 = 0

		log.Debugf("logger.Log...")
	Loop:
		for {
			select {
			case r, ok := <-l.chid:
				if !ok {
					break Loop
				}
				if r.Error() != nil {
					log.Errorf("%s report, error: %s", l.name, r.Error())
					break Loop
				}
				if r.Percentage() >= pre {
					log.Debugf("%s report: speed: %s, percentage: %f%%", l.name, r.Detail(), pre)
					pre += 10
				}
				if r.Percentage() >= 110 {
					break Loop
				}
			case <-l.over:
				break Loop
			}

		}
	}()
}
