package cmd

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 * @Classname cron
 * @Author Johnathan
 * @Date 2020/8/13 9:19
 * @Created by Goalnd 2020
 */
func CronTask() {
	c := cron.New()
	err := c.AddFunc("@every 5s", func() {
		fmt.Println("current time is ", time.Now())
		resp, err := http.Get("http://localhost:7688/api/v1/vote/cron/status")
		if err != nil {
			logrus.Error(err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("result is " + string(body))
	})
	err = c.AddFunc("@every 5s", func() {
		fmt.Println("current time is ", time.Now())
		resp, err := http.Get("http://localhost:7688/api/v1/vote/cron/send")
		if err != nil {
			logrus.Error(err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("result is " + string(body))
	})
	err = c.AddFunc("@every 5s", func() {
		fmt.Println("current time is ", time.Now())
		resp, err := http.Get("http://localhost:7688/api/v1/vote/cron/deploy")
		if err != nil {
			logrus.Error(err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("result is " + string(body))
	})
	err = c.AddFunc("@every 5s", func() {
		fmt.Println("current time is ", time.Now())
		resp, err := http.Get("http://localhost:7688/api/v1/vote/cron/repeat")
		if err != nil {
			logrus.Error(err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("result is " + string(body))
	})
	err = c.AddFunc("@every 1m", func() {
		fmt.Println("current time is ", time.Now())
		resp, err := http.Get("http://localhost:7688/api/v1/vote/cron/expired")
		if err != nil {
			logrus.Error(err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("result is " + string(body))
	})
	err = c.AddFunc("@every 1m", func() {
		fmt.Println("current time is ", time.Now())
		resp, err := http.Get("http://localhost:7688/api/v1/vote/cron/timeout")
		if err != nil {
			logrus.Error(err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("result is " + string(body))
	})
	err = c.AddFunc("@every 5m", func() {
		fmt.Println("current time is ", time.Now())
		resp, err := http.Get("http://localhost:7688/api/v1/vote/cron/verify")
		if err != nil {
			logrus.Error(err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("result is " + string(body))
	})
	if err != nil {
		panic(err)
	}
	c.Run()
}
