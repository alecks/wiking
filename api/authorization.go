package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.etcd.io/bbolt"
	"golang.org/x/crypto/bcrypt"
)

type signupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`

	Superuser      byte   `json:"superuser"`
	PermitUsername string `json:"permitUsername"`
	PermitPassword string `json:"permitPassword"`
}

type authorizeRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	routes = append(routes, route{
		path:   apiPath + "auth/user",
		method: "POST",
		handlers: []gin.HandlerFunc{
			func(c *gin.Context) {
				body := &signupRequest{}
				err := c.BindJSON(body)
				if err != nil {
					c.AbortWithStatus(400)
				}

				// TODO: Comment this
				err = db.Update(func(tx *bbolt.Tx) error {
					var err error
					bkt := tx.Bucket([]byte("Users"))
					sBkt := tx.Bucket([]byte("Superusers"))

					if bkt.Get([]byte(body.Username)) != nil {
						c.AbortWithStatus(409)
					}

					hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
					if err != nil {
						return err
					}

					if body.Superuser != superuserAllow["normalUser"] {
						firstUser := true
						sBkt.ForEach(func(k, v []byte) error {
							firstUser = false
							return nil
						})

						if firstUser {
							createUser(bkt, body, hash)
							sBkt.Put([]byte(body.Username), []byte{superuserAllow["ownerUser"]})
							log.Println("Registered a new ownerUser!")
							c.AbortWithStatus(201)
						}
						pHash := bkt.Get([]byte(body.PermitUsername))
						if pHash == nil {
							c.AbortWithStatus(401)
						}

						ok := false
						userType := ""
						for i, v := range superuserAllow {
							if v == body.Superuser {
								ok = true
								userType = i
								break
							}
						}
						if sudo := sBkt.Get([]byte(body.PermitUsername)); !ok || sudo == nil || sudo[0] < superuserAllow["adminUser"] {
							c.AbortWithStatus(403)
						}

						if err := bcrypt.CompareHashAndPassword(pHash, []byte(body.PermitPassword)); err == nil {
							createUser(bkt, body, hash)
							sBkt.Put([]byte(body.Username), []byte{body.Superuser})
							log.Println("Registered a new " + userType + ".")
							c.AbortWithStatus(201)
						}
					} else {
						createUser(bkt, body, hash)
						sBkt.Put([]byte(body.Username), []byte{superuserAllow["normalUser"]})
						log.Println("Registered a new normalUser.")
					}

					return err
				})

				if err != nil {
					c.AbortWithError(500, err)
				}
			},
		},
	})

	// This averages 15ms locally, which is quite slow in comparison to other endpoints.
	routes = append(routes, route{
		path:   apiPath + "auth/authorize",
		method: "POST",
		handlers: []gin.HandlerFunc{
			func(c *gin.Context) {
				body := &authorizeRequest{}
				err := c.BindJSON(body)
				if err != nil {
					// Couldn't parse JSON; bad request
					c.AbortWithStatus(400)
				}

				err = db.Update(func(tx *bbolt.Tx) error {
					var err error

					// All users
					bkt := tx.Bucket([]byte("Users"))
					// All logged in users
					lBkt := tx.Bucket([]byte("Authorized"))

					if hash := bkt.Get([]byte(body.Username)); hash != nil {
						// The user exists
						if err := bcrypt.CompareHashAndPassword(hash, []byte(body.Password)); err == nil {
							// The password was correct
							// Expiration: 168h = 7d = 1w
							// TODO: Improve the conversions; not exactly sure whether or not they're efficient.
							lBkt.Put([]byte(body.Username), []byte{byte(time.Now().Add(time.Hour * 168).Unix())})

							// User logged in; see other
							c.JSON(303, map[string]interface{}{
								"error":   false,
								"message": apiPath + "pages",
							})
						} else {
							// The password was incorrect; forbidden
							c.JSON(403, map[string]interface{}{
								"error":   true,
								"message": "Incorrect password.",
							})
						}
					} else {
						// Invalid username; unauthorized
						c.JSON(401, map[string]interface{}{
							"error":   true,
							"message": "Invalid username.",
						})
					}

					return err
				})
				if err != nil {
					// Error occurred while updating database; internal server error
					c.AbortWithError(500, err)
				}
			},
		},
	})
}

func createUser(bkt *bbolt.Bucket, body *signupRequest, hash []byte) error {
	return bkt.Put([]byte(body.Username), hash)
}
