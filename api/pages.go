package main

import (
	"github.com/labstack/echo"
	"go.etcd.io/bbolt"
)

func init() {
	routes = append(routes, func() {
		e.GET(apiPath+"pages/:page", func(c echo.Context) error {
			err := db.View(func(tx *bbolt.Tx) error {
				var err error

				bkt := tx.Bucket([]byte("Pages"))
				page := c.Param("page")
				if page == "" {
					var arr []string
					bkt.ForEach(func(k, v []byte) error {
						arr = append(arr, "k:v")
					})

					return c.JSON(200, map[string]interface{}{

					})
				}

				return err
			})

			return err
		})
	})
}
