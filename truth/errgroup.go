package truth

import (
	"fmt"
	"step/misc/randtools"

	"golang.org/x/sync/errgroup"
)

func justErrors() {

	urls := []string{
		"1",
		"2",
		"3",
	}
	var g errgroup.Group
	for i, url := range urls {
		i := i
		url := url
		g.Go(func() error {
			id := randtools.Range(1, 10)
			fmt.Printf("%v %v\n", i, id)
			if id != i {
				return fmt.Errorf("%v!=%v %v", i, id, url)
			}

			return nil
		})
	}

	er := g.Wait()
	if er != nil {
		fmt.Println(er)
	}
}
