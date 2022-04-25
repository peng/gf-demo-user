package dao_test

import (
	"testing"

	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao"
)

func TestCreate(t *testing.T) {
	err := dao.Trans()

	if err != nil {
		t.Error(err.Error())
	}
}