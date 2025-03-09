package main

import (
	"context"
	"time"
)

func locale(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
	}

	return "EN/US", nil
}

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "Hello", nil
	}
}

func printGreeting(ctx context.Context) error {
	greeting, err := gre
}

func main() {

}
