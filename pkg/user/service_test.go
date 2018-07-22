package user_test

import (
	"fmt"
	"github.com/hmoniaga/go-friend/pkg/user"
)

func ExampleService_AddFriend() {
	s := user.NewService(user.NewInMemoryRepository())
	s.AddFriend("harry@nite.com", "jeremiah@gmail.com")
	fmt.Println(s.ListFriends("harry@nite.com"))
	fmt.Println(s.ListFriends("jeremiah@gmail.com"))

	// Output:
	// [jeremiah@gmail.com] <nil>
	// [harry@nite.com] <nil>
}

func ExampleService_FriendList() {
	s := user.NewService(user.NewInMemoryRepository())
	s.AddFriend("harry@nite.com", "jeremiah@gmail.com")
	fmt.Println(s.ListFriends("harry@nite.com"))
	fmt.Println(s.ListFriends("unknown.guy@gmail.com"))

	// Output:
	// [jeremiah@gmail.com] <nil>
	// [] user not found
}

func ExampleService_MutualFriends() {
	s := user.NewService(user.NewInMemoryRepository())
	s.AddFriend("harry@nite.com", "mutual@gmail.com")
	s.AddFriend("jeremiah@gmail.com", "mutual@gmail.com")
	fmt.Println(s.ListMutualFriends("harry@nite.com", "jeremiah@gmail.com"))

	// Output:
	// [mutual@gmail.com] <nil>
}

func ExampleService_AddBlockedUser() {
	s := user.NewService(user.NewInMemoryRepository())
	s.AddBlockedUser("harry@nite.com", "mutual@gmail.com")
	fmt.Println(s.AddFriend("harry@nite.com", "mutual@gmail.com"))

	// Output:
	// harry@nite.com is blocking: mutual@gmail.com
}

func ExampleService_Notify() {
	s := user.NewService(user.NewInMemoryRepository())
	s.AddFriend("harry@nite.com", "mutual@gmail.com")
	fmt.Println(s.Notify("harry@nite.com", "hello katie@example.com"))

	// Output:
	// [katie@example.com mutual@gmail.com] <nil>
}
