package main

import (
	"gee"
	"log"
	"time"
)

func Logger() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s : %v \n", c.StatusCode, c.Path, time.Since(t))
	}
}

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(500, "server error\n")
		log.Printf("[%d] %s : %v for v2\n", c.StatusCode, c.Path, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Use(Logger())
	r.Static("/midea", "./static")

	r.Run(":9999")
}
