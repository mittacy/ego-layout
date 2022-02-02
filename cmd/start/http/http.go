package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/bootstrap"
	"github.com/mittacy/ego-layout/config"
	"github.com/mittacy/ego-layout/middleware"
	"github.com/mittacy/ego-layout/router"
	"github.com/mittacy/ego/hook"
	"github.com/mittacy/ego/library/async"
	"github.com/mittacy/ego/library/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var Cmd = &cobra.Command{
	Use:   "http",
	Short: "start the http api server. Example: server start http",
	Long:  "start the http api server. Example: server start http -c=.env.development -e=development -p=8080",
	Run:   run,
}

var (
	port int
	conf string
	env  string
	l    *log.Logger
)

func init() {
	Cmd.Flags().StringVarP(&conf, "conf", "c", ".env.development", "配置文件路径")
	Cmd.Flags().StringVarP(&env, "env", "e", "development", "运行环境")
	Cmd.Flags().IntVarP(&port, "port", "p", 8080, "监听端口")
}

func run(cmd *cobra.Command, args []string) {
	bootstrap.InitHTTP(conf, env, port)
	l = log.New("http")
	l.Infow("conf message", "conf", conf, "env", env, "port", port)

	// 初始化异步服务连接，后续才能插入队列
	async.Init(config.AsyncRedisClientOpt())

	// 启动http服务
	r := gin.New()
	r.Use(middleware.RecoveryWithZap(log.New("request"), true))
	r.Use(middleware.RequestTrace())
	r.Use(middleware.Ginzap(log.New("request"), time.RFC3339, true))

	router.InitRouter(r)
	router.InitAdminRouter(r)

	// 启动API服务
	if err := serve(r); err != nil {
		l.Panicw("API服务异常退出", "err", err)
	}
}

// serve 启动服务
func serve(r *gin.Engine) error {
	s := &http.Server{
		Addr:           ":" + viper.GetString("APP_PORT"),
		Handler:        r,
		ReadTimeout:    time.Second * viper.GetDuration("APP_READ_TIMEOUT"),
		WriteTimeout:   time.Second * viper.GetDuration("APP_WRITE_TIMEOUT"),
		MaxHeaderBytes: 1 << 20,
	}

	l.Infow("服务启动", "端口号", s.Addr)
	hook.Run(hook.BeforeStart)
	return s.ListenAndServe()
}
