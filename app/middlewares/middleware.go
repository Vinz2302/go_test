package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/app/config"
	"strconv"

	driver "rest-api/driver"
	res "rest-api/pkg/api-response"
	jwt "rest-api/pkg/jwt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	env, _              = config.Init()
	authorizationHeader = "Authorization"
	apiKeyHeader        = "X-API-key"
	cronExecutedHeader  = "X-Appengine-Cron"
	valName             = "FIREBASE_ID_TOKEN"
)

func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		authHeader := c.Request.Header.Get(authorizationHeader)
		token := strings.Replace(authHeader, "Bearer ", "", 1)

		validate_token, err := jwt.ValidateToken(token, env.App.Secret_key)
		if err != nil {
			errorMessage := fmt.Sprintf("%v", err)
			c.JSON(http.StatusUnauthorized, res.UnAuthorized(errorMessage))
			c.Abort()
			return
		}

		roleId, errExtract := jwt.ExtractTokenRoleID(validate_token)
		if errExtract != nil {
			errorMessage := fmt.Sprintf("%v", err)
			c.JSON(http.StatusUnauthorized, res.UnAuthorized(errorMessage))
			c.Abort()
			return
		}

		log.Println("auth time = ", startTime)
		c.Set("role_id", roleId)
		c.Set("token", token)
		c.Next()
	}
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authUser := c.Request.Header.Get("Authorization")
		fmt.Println("auth test = ", authUser)
		userRepository := driver.UserRepository
		intAuthUser, err := strconv.Atoi(authUser)
		if err != nil {
			errorMessage := fmt.Sprintf("%v", err)
			c.JSON(http.StatusUnauthorized, res.UnAuthorized(errorMessage))
			c.Abort()
			return
		}
		userLogin, err := userRepository.GetById(intAuthUser)
		if err != nil {
			errorMessage := fmt.Sprintf("%v", err)
			c.JSON(http.StatusUnauthorized, res.UnAuthorized(errorMessage))
			c.Abort()
			return
		}
		c.Set("user", userLogin)
		//c.Set("user_id", userLogin)
		//c.Set("user_name", *userLogin.Name)
		//c.Set("user_nip", *userLogin.Nip)
		//c.Set("user_role_id", *userLogin.RoleId)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Auth-User-Id")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, HEAD, PATCH, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
