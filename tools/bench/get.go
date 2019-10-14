package bench

import (
	"fmt"

	"github.com/flipkart-incubator/dkv/pkg/serverpb"
)

type GetHotKeysBenchmark struct {
	numHotKeys uint
}

func DefaultGetHotKeysBenchmark() *GetHotKeysBenchmark {
	return CreateGetHotKeysBenchmark(numHotKeys)
}

func CreateGetHotKeysBenchmark(numHotKeys uint) *GetHotKeysBenchmark {
	return &GetHotKeysBenchmark{numHotKeys}
}

func (this *GetHotKeysBenchmark) ApiName() string {
	return "dkv.serverpb.DKV.Get"
}

func (this *GetHotKeysBenchmark) CreateRequests(numRequests uint) interface{} {
	var getReqs []*serverpb.GetRequest
	for i, j := 0, 0; i < int(numRequests); i, j = i+1, (j+1)%int(this.numHotKeys) {
		key := []byte(fmt.Sprintf("ExistingKey%d", j))
		getReqs = append(getReqs, &serverpb.GetRequest{key})
	}
	return getReqs
}