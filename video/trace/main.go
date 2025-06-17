package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// 使用 context 实现链路追踪

type ctxKeyTraceId struct {
}

// 生成随机数

func CreateTraceId() string {
	return fmt.Sprintf("%d", rand.Int63())
}

//中间件，为每个context 注入 traceId

func TraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceId := CreateTraceId()
		ctx := context.WithValue(r.Context(), ctxKeyTraceId{}, traceId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func logWithTraceId(ctx context.Context, info string) {
	traceId, _ := ctx.Value(ctxKeyTraceId{}).(string)
	fmt.Printf("【traceId: %s】%s \n", traceId, info)
}

// 请求处理

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logWithTraceId(ctx, "request start")
	doSomething(ctx)
	doAnother(ctx)
	logWithTraceId(ctx, "request end")
	fmt.Fprintf(w, "ok\n")
}

func doSomething(ctx context.Context) {
	logWithTraceId(ctx, "doSomething...")
	go func(ctx context.Context) {
		time.Sleep(time.Millisecond * 500)
		logWithTraceId(ctx, "doSomething child request")
	}(ctx)
}
func doAnother(ctx context.Context) {
	logWithTraceId(ctx, "doAnother...")
}

func main() {
	//随机数因子
	rand.Seed(time.Now().UnixNano())
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	svr := &http.Server{Addr: ":9091", Handler: TraceMiddleware(mux)}
	fmt.Println("listen start :9091")
	if err := svr.ListenAndServe(); err != nil {
		fmt.Println("err:", err)
	}
}
