package main

import "fmt"

type Ball struct {
	Radius   int
	Material string
}

type Bouncer interface {
	Bounce()
}

// football inherits Ball

type Football struct {
	Ball
}

// we can also use pointers
// note that if the instance in empty

type BasketBall struct {
	*Ball
}

func (b Ball) Bounce() {
	fmt.Println("Boucing ball => ", b)
}

func BounceBouncer(b Bouncer) {
	fmt.Println("BounceBouncer")
	b.Bounce()
}

func main() {
	fb1 := Football{
		Ball{
			Radius:   5,
			Material: "cheap plastic",
		},
	}
	fmt.Printf("fb = %v\n", fb1)
	// we access this method directly on a Football instance eventhough it is defined on the parent struct Ball
	fb1.Bounce()
	// synonym
	fb1.Ball.Bounce()
	bb1 := BasketBall{
		&Ball{
			Radius:   10,
			Material: "Leather",
		},
	}

	// this panics (nil pointer error) because we instantiated an "empty" ball:
	//  the bounce method only handles actual pointers
	/* bb2 := BasketBall{}
	bb2.Bounce() */

	bb1.Bounce()
}
