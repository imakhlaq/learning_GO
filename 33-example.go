package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type User struct {
	id       string
	name     string
	comments []string
	friends  []string
}

type Response struct {
	data interface{}
	err  error //if in case of error
}

func getUserDetails(ctx context.Context) (*User, error) {
	ctx = context.WithValue(ctx, "id", 33)
	var wait sync.WaitGroup
	resChan := make(chan Response, 2) //because we will write 2 times before ready

	//u can do one request at a time and block the goroutine
	// but instead u can use goroutine and channels to make request concurrently

	wait.Add(2) //because starting 2 goroutine
	// fetch the comments
	go getComments(ctx, resChan, &wait)
	//fetch friends
	go getFriends(ctx, resChan, &wait)

	wait.Wait() //waiting for goroutines to finish
	close(resChan)

	userProfile := &User{id: "99i29029i", name: "Akhlaq"}

	for data := range resChan {

		switch mgs := data.data.(type) { // stitching base on the data type
		case []string:
			userProfile.comments = mgs //if data is []string
		case int:
			fmt.Println("type is int")
		}
	}
	return userProfile, nil
}

func getComments(ctx context.Context, resChan chan<- Response, wait *sync.WaitGroup) {
	id := ctx.Value("id").(int) //explicit type casting
	fmt.Println(id)

	time.Sleep(200 * time.Millisecond) //simulating network call delay

	comments := []string{
		"How u get their",
		"How u get their",
		"Hey buddy",
	}

	resChan <- Response{
		data: comments,
		err:  nil, // incase of error
	}
	wait.Done() //signal that this goroutine is finished
}

func getFriends(ctx context.Context, resChan chan<- Response, wait *sync.WaitGroup) {
	id := ctx.Value("id").(int) //casting the value to int
	fmt.Println(id)

	resChan <- Response{data: []string{"19danfjnh", "sfjqjdj112939", "wfjwifjwijf122"}, err: nil}
	wait.Done()
}

func main() {

	ctx := context.Background()
	res, err := getUserDetails(ctx)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}
