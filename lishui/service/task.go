package service

import "github.com/gogf/gf/frame/g"

var Task = new(taskService)
type taskService struct{}
func (this *taskService) UpdateTask(taskId int) error {
	_,err:=g.DB().Table("s_task").Update(g.Map{"status":1},"id=?",taskId)
	return err
}
