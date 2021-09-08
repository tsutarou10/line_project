package dynamo

type utnaFoodRegisterStatus struct {
	Status    string `dynamo:"status"`
	Number    int64  `dynamo:"number"`
	UpdatedAt int64  `dynamo:"updatedAt"`
}
