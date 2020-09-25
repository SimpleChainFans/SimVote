package cmd

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
)

const Version = "0.0.1"
const VoteName = "vote_app"

type DBConfig struct {
	User     string `default:"root" yaml:"user"`
	Password string `default:"" yaml:"password"`
	Name     string `yaml:"ip"`
	Port     uint   `default:"3306" yaml:"port"`
	DbName   string `required:"true" yaml:"db_name"`
	Charset  string `default:"utf8" yaml:"charset"`
	MaxIdle  int    `default:"10" yaml:"max_idle"`
	MaxOpen  int    `default:"50" yaml:"max_open"`
	LogMode  bool   `yaml:"log_mode"`
	Loc      string `required:"true" yaml:"loc"`
}

func GetDBConnection(conf *DBConfig) (*gorm.DB, error) {
	format := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s"
	dsn := fmt.Sprintf(format, conf.User, conf.Password, conf.Name, conf.Port, conf.DbName, conf.Charset, url.QueryEscape(conf.Loc))
	logrus.Infof("dsn=%s", dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.LogMode(conf.LogMode)
	db.DB().SetMaxIdleConns(conf.MaxIdle)
	db.DB().SetMaxOpenConns(conf.MaxOpen)
	return db, nil
}

type JwtCustomClaims struct {
	UserId      uint
	LoginVerify int
	jwt.StandardClaims
}

func GetLoggerConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return false
		},
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           os.Stdout,
	}
}

func (this *server) GetUserId(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	userId := claims.UserId
	return userId
}

func LoginVerify() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*JwtCustomClaims)
			if claims.LoginVerify == 0 {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "invalid or expired jwt",
					Internal: nil,
				}
			}
			return next(c)
		}
	}
}

type ApiResponse struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 状态短语
	Result  interface{} `json:"result"`  // 数据结果集
}

func ResponseOutput(c echo.Context, code int, message string, result interface{}) error {
	return c.JSON(http.StatusOK, ApiResponse{
		Code:    code,
		Message: message,
		Result:  result,
	})
}

func ResponseError(c echo.Context, code int, message string) error {
	return ResponseOutput(c, code, message, nil)
}

func (this *server) ResponseSuccess(c echo.Context, result interface{}) error {
	return ResponseOutput(c, 0, "操作成功", result)
}

func (this *server) ResponseError(c echo.Context, code int, message string) error {
	return ResponseOutput(c, code, message, nil)
}

func (this *server) InvalidParametersError(c echo.Context) error {
	return ResponseError(c, 1, "参数不合法")
}

func (this *server) InternalServiceError(c echo.Context) error {
	return ResponseError(c, 2, "内部错误")
}

func (this *server) NoPermissionError(c echo.Context) error {
	return ResponseError(c, 3, "无权限访问")
}
