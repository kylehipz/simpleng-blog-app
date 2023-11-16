package types

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}

type Follow struct {
	Follower string `json:"follower"`
	Followee string `json:"followee"`
}
