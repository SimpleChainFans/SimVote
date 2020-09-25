package cmd

import (
	"fmt"
	"github.com/didip/tollbooth"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"strings"
)

type Server interface {
	Run(port string, configPath string) error
	Close() error
}

type serverConfig struct {
	SipcApi         string    `required:"true" yaml:"sipc_Api"`
	LocalApi        string    `required:"true" yaml:"local_Api"`
	AppId           string    `required:"true" yaml:"appId"`
	AppKey          string    `required:"true" yaml:"appKey"`
	VoteAmount      string    `required:"true" yaml:"voteAmount"`
	VerifyURL       string    `required:"true" yaml:"verify_url"`
	ContractVersion string    `required:"true" yaml:"contract_version"`
	EthRpcHost      string    `required:"true" yaml:"eth_rpc"`
	CallBackURL     string    `required:"true" yaml:"callBackUrl"`
	IndexURL        string    `required:"true" yaml:"indexUrl"`
	VotePrivateKey  string    `required:"true" yaml:"vote_privateHexKey"`
	VoteAddress     string    `required:"true" yaml:"vote_address"`
	Secret          string    `required:"true" yaml:"secret"`
	DbConfig        *DBConfig `required:"true" yaml:"vote_app"`
}

type server struct {
	name string

	conf   *serverConfig
	db     *gorm.DB
	engine *echo.Echo
}

func NewServer(name string) Server {
	s := new(server)
	s.name = name
	return s
}

func (rs *server) Run(port string, configPath string) error {
	// config
	if err := rs.config(configPath); err != nil {
		return fmt.Errorf("rs.config(): %s", err.Error())
	}

	// mysql
	if err := rs.dbClient(); err != nil {
		return fmt.Errorf("rs.dbClient(): %s", err.Error())
	}

	// router
	if err := rs.router(); err != nil {
		return fmt.Errorf("rs.router(): %s", err.Error())
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	return rs.engine.Start(port)
}

func (rs *server) router() error {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.SecureWithConfig(middleware.DefaultSecureConfig))

	jwtConfig := &middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(rs.conf.Secret),
	}
	limiter := tollbooth.NewLimiter(20, nil)
	limiter.SetIPLookups([]string{"X-Forwarded-For", "X-Real-IP", "RemoteAddr"})

	e.Use(middleware.LoggerWithConfig(GetLoggerConfig()))
	e.Use(middleware.Recover())

	rs.engine = e

	NewRouters(rs, e, jwtConfig)
	return nil
}

func (rs *server) config(configPath string) error {
	rs.conf = new(serverConfig)
	err := configor.Load(rs.conf, configPath)
	if err != nil {
		return err
	}
	return nil
}

func (rs *server) dbClient() error {
	db, err := GetDBConnection(rs.conf.DbConfig)
	if err != nil {
		return fmt.Errorf("%+v", err)
	}
	rs.db = db
	return nil
}

func (rs *server) Close() error {

	if err := rs.engine.Close(); err != nil {
		return err
	}

	if rs.db != nil {
		if err := rs.db.Close(); err != nil {
			return err
		}
	}
	return nil
}
