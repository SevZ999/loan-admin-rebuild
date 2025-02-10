package pools

import "github.com/small-ek/antgo/utils/pool"

func Register() {
	pool.New(10, 500)
}
