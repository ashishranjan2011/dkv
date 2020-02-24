package sync

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/flipkart-incubator/dkv/internal/server/storage"
	"github.com/flipkart-incubator/dkv/internal/server/sync/raftpb"
	"github.com/flipkart-incubator/dkv/pkg/serverpb"
	"github.com/gogo/protobuf/proto"
)

type dkvReplStore struct {
	kvs storage.KVStore
}

// NewDKVReplStore creates a wrapper out of the given KVStore
// that performs synchronous replication of all operations
// over Nexus onto multiple replicas.
func NewDKVReplStore(kvs storage.KVStore) *dkvReplStore {
	return &dkvReplStore{kvs}
}

func (dr *dkvReplStore) Save(req []byte) ([]byte, error) {
	intReq := new(raftpb.InternalRaftRequest)
	if err := proto.Unmarshal(req, intReq); err != nil {
		return nil, err
	}
	switch {
	case intReq.Put != nil:
		return dr.put(intReq.Put)
	case intReq.Get != nil:
		return dr.get(intReq.Get)
	case intReq.MultiGet != nil:
		return dr.multiGet(intReq.MultiGet)
	default:
		return nil, errors.New("Unknown request to Save in dkv")
	}
}

func (dr *dkvReplStore) put(putReq *serverpb.PutRequest) ([]byte, error) {
	err := dr.kvs.Put(putReq.Key, putReq.Value)
	return nil, err
}

func (dr *dkvReplStore) get(getReq *serverpb.GetRequest) ([]byte, error) {
	if vals, err := dr.kvs.Get(getReq.Key); err != nil {
		return nil, err
	} else {
		return vals[0], nil
	}
}

func (dr *dkvReplStore) multiGet(multiGetReq *serverpb.MultiGetRequest) ([]byte, error) {
	if vals, err := dr.kvs.Get(multiGetReq.Keys...); err != nil {
		return nil, err
	} else {
		return gobEncode(vals)
	}
}

func gobEncode(byteArrays [][]byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(byteArrays); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (dr *dkvReplStore) Close() error {
	return dr.kvs.Close()
}

func (dr *dkvReplStore) Backup() ([]byte, error) {
	// TODO: Implement this
	return nil, nil
}

func (dr *dkvReplStore) Restore(data []byte) error {
	// TODO: Implement this
	return nil
}
