package test

import (
	"fmt"
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/repository"
	"math"
	"testing"
)

func TestSliceOfMTS(t *testing.T) {
	slice := repository.GetSliceOfMTS(0, math.MaxInt64)
	if slice == nil {
		t.Error("nil")
	}
	for i, e := range slice {
		fmt.Print(i, " = ")
		fmt.Println(e.Data.Id)
	}
}
