package common

import "bestHabit/pubsub"

const (
	TopicUserCreateNewTask  pubsub.Topic = "TopicUserCreateNewTask"
	TopicUserCreateNewHabit pubsub.Topic = "TopicUserCreateNewHabit"
	TopicUserDeleteHabit    pubsub.Topic = "TopicUserDeleteHabit"
	TopicUserDeleteTask     pubsub.Topic = "TopicUserDeleteTask"
)
