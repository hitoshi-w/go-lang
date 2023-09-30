package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"golang.org/x/sync/errgroup"
)


func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed: %v", err)
	}
}

func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":3000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "hello", r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
			log.Printf("failed to cloade %v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown %v", err)
	}

	return eg.Wait()
}