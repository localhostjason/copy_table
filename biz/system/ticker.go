package system

import "sync"

// Task 任务接口，1. 拍卖行 2. 金币寄售 工厂模式
type Task interface {
	Run()
	Stop()
	SetWaitGroup(wg *sync.WaitGroup)
}

type TaskManage struct {
	Tasks []Task
	Wg    *sync.WaitGroup
}

func NewTaskManage() *TaskManage {
	return &TaskManage{
		Wg: &sync.WaitGroup{},
	}
}

func (m *TaskManage) Add(task Task) {
	m.Wg.Add(1)
	m.Tasks = append(m.Tasks, task)
	task.SetWaitGroup(m.Wg)
}

func (m *TaskManage) Run() {
	for _, task := range m.Tasks {
		task.Run()
	}
}

func (m *TaskManage) Stop() {
	for _, task := range m.Tasks {
		task.Stop()
	}
}
