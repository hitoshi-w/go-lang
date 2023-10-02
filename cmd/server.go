package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

// http.Serverはカスタマイズ性がある
type Server struct {
	srv *http.Server
	l net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l: l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// os.Interrupt(Ctrl＋C)とsyscall.SIGTERM(プロセス終了)のシグナルを受信したときにコンテキストをキャンセル
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)

	// エラーグループに別ゴルーチンを登録し、サーバーを起動する。
	/// サーバーの実行中にエラーが発生したときはエラーグループに報告。ErrServerClosedはサーバーが正常に終了したことを示すので異常ではない
	eg.Go(func() error {
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close %v", err)
			return err
		}
		return nil
	})

	// チャンネルからの通知（終了通知）を待機する（通知が来たらこれ以降の処理が実行される）
	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown %v", err)
	}

	// Goメソッドで起動した別ごルーチンの終了を待つ
	return eg.Wait()
}

