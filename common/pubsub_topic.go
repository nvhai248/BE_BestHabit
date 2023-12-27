package common

import "bestHabit/pubsub"

const (
	TopicUserCreateNewTask   pubsub.Topic = "TopicUserCreateNewTask"
	TopicUserCreateNewHabit  pubsub.Topic = "TopicUserCreateNewHabit"
	TopicUserJoinChallenge   pubsub.Topic = "TopicUserJoinChallenge"
	TopicUserDeleteHabit     pubsub.Topic = "TopicUserDeleteHabit"
	TopicUserDeleteTask      pubsub.Topic = "TopicUserDeleteTask"
	TopicUserCancelChallenge pubsub.Topic = "TopicUserCancelChallenge"
	TopicUserUpdateTask      pubsub.Topic = "TopicUserUpdateTask"
	TopicUserUpdateHabit     pubsub.Topic = "TopicUserUpdateHabit"
)
