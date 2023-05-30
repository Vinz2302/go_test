package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/app/config"

	//driver "rest-api/driver"
	res "rest-api/pkg/api-response"
	jwt "rest-api/pkg/jwt"

	//"github.com/dgrijalva/jwt-go"
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
		fmt.Println("auth = ", authHeader)
		fmt.Println("token", token)

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
		//c.Set("token", token)
		c.Next()
	}
}

func RoleAuth(role int) gin.HandlerFunc {
	return func(c *gin.Context) {

		roleId, exists := c.Get("role_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, res.UnAuthorized("Role id not found"))
			c.Abort()
			return
		}

		storedRoleId, ok := roleId.(*int)
		if !ok {
			c.JSON(http.StatusUnauthorized, res.UnAuthorized("Invalid role id format"))
			c.Abort()
			return
		}

		if *storedRoleId != role {
			c.JSON(http.StatusUnauthorized, res.UnAuthorized("Insufficient privileges"))
			c.Abort()
			return
		}

		c.Next()
	}
}

/* func DeserializeUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var access_token string

		cookie, err := c.Cookie("access_token")

		authorizationHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, res.UnAuthorized("You're not logged in"))
		}

		sub, err := jwt.ValidateToken(access_token, env.App.Secret_key)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, res.UnAuthorized(err.Error()))
		}

		userRepository := driver.UserRepository
	}
} */

/* func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Println("auth test = ", authHeader)

		token := extractTokenFromHeader(authHeader)

		fmt.Println("token =", token)

		userRepository := driver.UserRepository
		intAuthUser, err := strconv.Atoi(authHeader)
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
		c.Set("user_id", userLogin)
		c.Set("user_name", *userLogin.Name)
		c.Set("user_nip", *userLogin.Nip)
		c.Set("user_role_id", *userLogin.RoleId)
		c.Next()
	}
} */

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
