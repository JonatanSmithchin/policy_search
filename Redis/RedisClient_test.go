package Redis

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestRecord(t *testing.T) {
	for i := 10; i < 31; i++ {
		err := Record(strconv.Itoa(i), 10)
		require.NoError(t, err)
	}

}

func TestUpDateRec(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 10; i < 31; i++ {
		for j := 0; j < 12; j++ {
			err := UpDateRec(strconv.Itoa(i), strconv.Itoa(rand.Intn(10)))
			require.NoError(t, err)
		}
	}

}

func TestGetRec(t *testing.T) {
	result, err := GetRec("14", "8")
	res := reflect.ValueOf(result)
	log.Print(res.Kind())
	v := res.Interface()
	rec := v.([]interface{})
	s := fmt.Sprintf("%v", rec)
	ss := strings.TrimPrefix(s, "[")
	sss := strings.TrimSuffix(ss, "]")
	//var rec interface{}
	//if res.Kind() == reflect.Slice {
	//	rec = res.Index(0)
	//	i := rec.([]interface{})
	//	for _, i2 := range i {
	//		log.Print(i2)
	//	}
	//}
	require.NoError(t, err)
	log.Print(s)
	log.Print(sss)
}

func TestUpdateRecm(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		err := UpdateRecm("4", strconv.Itoa(100))
		require.NoError(t, err)
	}
}

func TestGetAllRecommend(t *testing.T) {
	res, err := GetAllRecommend("4")
	for id, v := range res {
		log.Println(id + ": " + v)
	}
	require.NoError(t, err)
}
