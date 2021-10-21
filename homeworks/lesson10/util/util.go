package util

import (
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang/homeworks/lesson10/entities"
	"math/rand"
)

/* Sinh mã unique cho primary key
 */
func NewID() (id string) {
	id, _ = gonanoid.New(8)
	return
}

/* Kiểm tra xem biến a có nằm trong mảng int_arr
 */
func Int_in_array(a int, int_arr []int) bool {
	for _, b := range int_arr {
		if b == a {
			return true
		}
	}
	return false
}

/*
Kiểm tra err khác nil thì rollback transaction
*/
func Check_err(err error, trans *pg.Tx) bool {
	if err != nil {
		_ = trans.Rollback()
		return false
	}
	return true
}

var (
	random *rand.Rand // Đối tượng dùng để tạo random number
)

/*
Sinh ngẫu nhiên các roles trả về mảng int và string.
Mảng string dùng cho enum
*/
func Gen_random_roles(numberOfRoles int) ([]int, []string) {
	int_roles := []int{}
	enum_roles := []string{}
	for i := 0; i < numberOfRoles; i++ {
		var role int
		for { //Loop cho đến khi tạo ra phần tử mới
			role = 1 + random.Intn(8)
			if !Int_in_array(role, int_roles) {
				break
			}
		}

		int_roles = append(int_roles, role)
		enum_roles = append(enum_roles, entities.ROLES[role])
	}
	return int_roles, enum_roles
}

/*
Chuyển đổi int[] thành string[]
*/
func Convert_introles_to_enumroles(int_roles []int) (enum_roles []string) {
	for _, role := range int_roles {
		enum_roles = append(enum_roles, entities.ROLES[role])
	}
	return enum_roles
}


func ResponseErr(c *fiber.Ctx, statusCode int, message string) (error){
	return c.JSON(fiber.Map{
		"status":  statusCode,
		"message": message,
	})
}