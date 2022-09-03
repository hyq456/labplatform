package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"labplatform/model"
	"labplatform/utils/errmsg"
	"strconv"
	"time"
)

//token的过期时长
const UserDuration = time.Minute * 5

//cache的名字
func getUserCacheName(userId int) string {
	return "user_" + strconv.Itoa(userId)
}

func GetOneUserCache(userId int) (model.User, int) {
	key := getUserCacheName(userId)
	val, err := model.DbRedis.Get(model.Ctx, key).Result()
	if err == redis.Nil || err != nil {
		return model.User{}, errmsg.ERROR
	} else {
		model.DbRedis.Expire(model.Ctx, key, UserDuration)
		user := model.User{}
		if err := json.Unmarshal([]byte(val), &user); err != nil {
			//t.Error(target)
			return model.User{}, errmsg.ERROR
		}
		return user, errmsg.SUCCSE
	}
}

//向cache保存一篇文章
func SetOneUserCache(userId int, user model.User) int {
	key := getUserCacheName(userId)
	content, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	errSet := model.DbRedis.Set(model.Ctx, key, content, UserDuration).Err()
	if errSet != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
