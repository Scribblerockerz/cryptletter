package scheduler

import (
	"github.com/Scribblerockerz/cryptletter/pkg/attachment"
	"github.com/Scribblerockerz/cryptletter/pkg/logger"
	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
	"time"
)

type Scheduler interface {
	ScheduleAttachmentCleanup()
	StartAsync()
}
type scheduler struct{
	instance *gocron.Scheduler
}

//ScheduleAttachmentCleanup will setup the task to cleanup all timed out attachments
func (s scheduler) ScheduleAttachmentCleanup() {
	cronExpression := viper.GetString("app.attachments.cleanup_schedule")
	if cronExpression == "" {
		cronExpression = "* * * * *"
	}

	s.instance.Cron(cronExpression).Do(func() {
		logger.LogInfo("Perform attachment cleanup")
		attachmentHandler := attachment.NewAttachmentHandler(viper.GetString("app.attachments.driver"))
		attachmentHandler.Cleanup()
	})
}

//StartAsync will activate the scheduler
func (s scheduler) StartAsync() {
	s.instance.StartAsync()
}

//NewScheduler will create a new scheduler instance
func NewScheduler() Scheduler {
	return &scheduler{
		instance: gocron.NewScheduler(time.UTC),
	}
}
