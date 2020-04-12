package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/labstack/echo"
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
	routes = append(routes, func() {
		e.POST(apiPath+"auth/user", func(c echo.Context) error {
			body := &signupRequest{}
			err := json.NewDecoder(c.Request().Body).Decode(body)
			if err != nil {
				// Invalid JSON; bad request
				return echo.NewHTTPError(400, err)
			}

			// TODO: Comment this
			err = db.Update(func(tx *bbolt.Tx) error {
				var err error
				bkt := tx.Bucket([]byte("Users"))
				sBkt := tx.Bucket([]byte("Superusers"))

				if bkt.Get([]byte(body.Username)) != nil {
					// User already exists; conflict
					return echo.NewHTTPError(409, "User already exists.")
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
						// User created as owner; created
						return c.String(201, "User created as owner.")
					}
					pHash := bkt.Get([]byte(body.PermitUsername))
					if pHash == nil {
						// Permitting user doesn't exist; unauthorized
						return echo.NewHTTPError(401, "Permitting user doesn't exist.")
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
						// Permitting user's level isn't sufficient; forbidden
						return echo.NewHTTPError(403, "Permitting user's level isn't sufficient.")
					}

					if err := bcrypt.CompareHashAndPassword(pHash, []byte(body.PermitPassword)); err == nil {
						createUser(bkt, body, hash)
						sBkt.Put([]byte(body.Username), []byte{body.Superuser})
						log.Println("Registered a new " + userType + ".")
						// Registered a new #{userType}; created
						return c.String(201, "Registered a new "+userType+".")
					}
				} else {
					// TODO: Integrate into the stuff above ^
					createUser(bkt, body, hash)
					sBkt.Put([]byte(body.Username), []byte{superuserAllow["normalUser"]})
					log.Println("Registered a new normalUser.")
					// Registered a new normalUser; created
					return c.String(201, "Registered a new normalUser.")
				}

				return err
			})

			return err
		})

		// This averages 15ms locally, which is quite slow in comparison to other endpoints.
		e.POST(apiPath+"auth/authorize", func(c echo.Context) error {
			body := &authorizeRequest{}
			err := json.NewDecoder(c.Request().Body).Decode(body)
			if err != nil {
				// Couldn't parse JSON; bad request
				return echo.NewHTTPError(400, "Couldn't parse JSON.")
			}

			err = db.Update(func(tx *bbolt.Tx) error {
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
						return c.String(303, apiPath+"pages")
					}

					// The password was incorrect; forbidden
					return c.String(403, "Incorrect password.")
				}
				// Invalid username; unauthorized
				return c.JSON(401, "Invalid username.")
			})
			return err
		})
	})
}

func createUser(bkt *bbolt.Bucket, body *signupRequest, hash []byte) error {
	return bkt.Put([]byte(body.Username), hash)
}
