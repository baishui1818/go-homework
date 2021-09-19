package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
)

//启动一个 http server
func StartHttpServer(srv *http.Server) error {
	http.HandleFunc("/test", TestServer2)
	fmt.Println("http server start")
	err := srv.ListenAndServe()
	return err
}

//增加一个 http hanlder
func TestServer2(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, " test server2!\n")
}

func main() {
	//返回一个空的Context
	ctx := context.Background()
	//创建一个可取消的子Context 用于取消下游的 context
	ctx, cancel := context.WithCancel(ctx)
	//通过 WithContext 可以创建一个带取消的 Group 使用 errgroup 取消 goroutine
	group, errCtx := errgroup.WithContext(ctx)

	srv := &http.Server{Addr: ":9090"}

	group.Go(func() error {
		return StartHttpServer(srv)
	})

	group.Go(func() error {
		<-errCtx.Done()
		fmt.Println("http server stop")
		return srv.Shutdown(errCtx)
	})

	chanel := make(chan os.Signal, 1)
	signal.Notify(chanel)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-chanel:
				cancel()
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")

}
