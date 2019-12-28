package main

import (
	"log"

	"gopl.io/ch8/thumbnail"
)

// func main() {
// 	images := []string{"a.png", "b.png", "main.go"}
// 	err := makeThumbnails(images)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

func main() {
	images := []string{"a.png", "b.png", "main.go"}
	thumbnails, err := makeThumbnails2(images)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(thumbnails)
}

func makeThumbnails(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}

	return nil
}

func makeThumbnails2(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	// buffered channel
	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}

		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}
