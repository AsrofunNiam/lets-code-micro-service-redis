package helper

import "github.com/AsrofunNiam/lets-code-micro-service-redis/model/domain"

func ListKey() domain.QueueKeys {
	return domain.QueueKeys{
		Keys: []string{"task_first", "task_second", "task_third"},
	}
}
