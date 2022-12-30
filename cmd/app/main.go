package main

import (
	"log"
	"sync"

	"github.com/aerosystems/nix-beginner-6/internal/resource"
)

func main() {
	userId := 7
	posts, err := resource.GetPosts(userId)
	if err != nil {
		log.Println(err)
	}

	chanComment := make(chan resource.Comment)
	wg := new(sync.WaitGroup)
	defer wg.Wait()

	for _, post := range posts {
		wg.Add(1)
		go func(post resource.Post) {
			defer wg.Done()
			comments, err := resource.GetComments(post.Id)
			if err != nil {
				log.Println(err)
			}

			for _, comment := range comments {
				wg.Add(1)
				go func(comment resource.Comment) {
					defer wg.Done()
					chanComment <- comment
				}(comment)

			}
		}(post)
	}

	go func() {
		for {
			resource.WriteComment(<-chanComment)
		}
	}()
}
