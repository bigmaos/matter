package task

import (
	"fmt"

	"github.com/mlee-msl/taskgroup"
)

type TaskPacker struct {
	tg    *taskgroup.TaskGroup
	tasks []*taskgroup.Task
}

type TaskUnit struct {
	TaskF    taskgroup.TaskFunc
	TaskNO   int
	MustSucc bool
}

// RegisterTasksMustSucc 注册必须成功的tasks
func (p *TaskPacker) RegisterTasksMustSucc(tasks []*TaskUnit) error {
	for _, t := range tasks {
		t.MustSucc = true
	}
	return p.RegisterTasks(tasks)
}

func (p *TaskPacker) RegisterTasks(tasks []*TaskUnit) error {
	if checkTaskNOConflict(tasks) {
		return fmt.Errorf("Register Task: multi TaskNO same")
	}
	for _, t := range tasks {
		newTask := taskgroup.NewTask(uint32(t.TaskNO), t.TaskF, t.MustSucc)
		p.tasks = append(p.tasks, newTask)
	}
	return nil
}

// 检查是否有冲突的taskNO
func checkTaskNOConflict(tasks []*TaskUnit) bool {
	var taskNOMap = make(map[int]bool)
	for _, t := range tasks {
		if _, ok := taskNOMap[t.TaskNO]; ok {
			return true
		}
		taskNOMap[t.TaskNO] = true
	}
	return false
}

func (p *TaskPacker) InitTaskGroup(tasks []*TaskUnit, allSucc bool) error {
	if allSucc {
		err := p.RegisterTasksMustSucc(tasks)
		if err != nil {
			return err
		}
	} else {
		err := p.RegisterTasks(tasks)
		if err != nil {
			return err
		}
	}

	// 清空上次的taskPacker
	p.tg = nil
	p.tg = taskgroup.NewTaskGroup(
		taskgroup.WithWorkerNums(uint32(len(p.tasks))),
	)
	p.tg.AddTask(p.tasks...)
	return nil
}

func (p *TaskPacker) RunTaskGroupOnce() (map[uint32]*taskgroup.TaskResult, error) {
	return p.tg.RunExactlyOnce()
}
