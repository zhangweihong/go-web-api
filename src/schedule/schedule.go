package schedule

import "github.com/jasonlvhit/gocron"

var s *gocron.Scheduler

//开启定时
func Start() {
	s = gocron.NewScheduler()
	s.Start()
}

//开始一个循环间隔定时任务
func AddEvery(seconds uint64, exec func()) *gocron.Job {
	jb := s.Every(seconds).Seconds()
	jb.Do(func() {
		if exec != nil {
			exec()
		}
	})
	return jb
}

//开始一个执行一次的定时任务
func Once(seconds uint64, exec func()) *gocron.Job {
	jb := s.Every(seconds).Seconds()
	jb.Do(func() {
		if exec != nil {
			exec()
		}
		s.RemoveByRef(jb)
	})
	return jb
}

//停止一个任务
func Stop(job *gocron.Job) {
	s.RemoveByRef(job)
}

//停止所有的任务
func StopAll() {
	s.Clear()
}
