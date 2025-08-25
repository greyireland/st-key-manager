package pricing

import (
	"fmt"
)

type Role int

func (r Role) String() string {
	switch r {
	case RoleGuest:
		return "introductory"
	case RolePro:
		return "professional"
	case RoleEnt:
		return "enterprise"
	case RoleUltra:
		return "ultra"
	case RoleSuper:
		return "super"
	default:
		return fmt.Sprintf("unknown(%d)", r)
	}
}

const (
	RoleGuest Role = iota
	RolePro
	RoleEnt
	RoleUltra
	RoleSuper
)
const (
	PlanTypeFree       = "free"
	PlanTypePro        = "professional"
	PlanTypeEnterprise = "enterprise"
	PlanTypeUltra      = "ultra"
	PlanTypeDev        = "dev"
)

type Data struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Pat           string `json:"pat"`
	ValidPlanType string `json:"valid_plan_type"`
}

type Response struct {
	Code    int    `json:"code"`
	Data    []Data `json:"data"`
	Message string `json:"message"`
}

type UserInfo struct {
	AccountId string
	Passwd    string
	Role      Role
	Auth      string
}

var AuthMap = []string{
	"ZG727ghhtFnt7iuwp5JfjSDqMbJkR7Vk",
	"QsCIFf7oHO1rqHVW5I2VGEnsy16FFU1o",
	"4XAsJxygdrWp4JRrgwaceZBrVwxuq62V",
	"AqRNcpCtcO4ppZj3kliuFNL0zdxq59E5",
	"1cxD1esr3oECABpwPBHKZabM07fW0UIN",
	"CVlOktLWSUrFOrAUZaNmUHf1UVReOpsE",
	"XKOif1FSSwMM48BAxqUBR1NbeQEJyFlF",
	"j72JiwnhV5RzYgiBp6vhYBrE7flvjMVH",
	"B7Poni7ElSxz9dSzou0stXXWpSJXYcIi",
	"UG3q1UCigTPHGVGQHTNSRPyv5hkkHSW8",
}

func GetUserInfo() []UserInfo {
	users := []UserInfo{}
	for i := 0; i < len(AuthMap); i++ {
		users = append(users, UserInfo{
			Auth: AuthMap[i],
			Role: RolePro,
		})
	}
	return users
}
