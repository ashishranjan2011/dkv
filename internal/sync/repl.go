package sync

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/golang/protobuf/proto"
	"io"

	"github.com/flipkart-incubator/dkv/internal/storage"
	"github.com/flipkart-incubator/dkv/internal/sync/raftpb"
	"github.com/flipkart-incubator/dkv/pkg/serverpb"
	"github.com/flipkart-incubator/nexus/pkg/db"
)

type dkvReplStore struct {
	kvs storage.KVStore
}

// NewDKVReplStore creates a wrapper out of the given KVStore
// that performs synchronous replication of all operations
// over Nexus onto multiple replicas.
func NewDKVReplStore(kvs storage.KVStore) db.Store {
	return &dkvReplStore{kvs}
}

func (dr *dkvReplStore) Save(_ db.RaftEntry, req []byte) (res []byte, err error) {
	intReq := new(raftpb.InternalRaftRequest)
	if err = proto.Unmarshal(req, intReq); err != nil {
		return nil, err
	}

	switch {
	case intReq.Put != nil:
		res, err = dr.put(intReq.Put)
	case intReq.MultiPut != nil:
		res, err = dr.multiPut(intReq.MultiPut)
	case intReq.Delete != nil:
		res, err = dr.delete(intReq.Delete)
	case intReq.Cas != nil:
		res, err = dr.cas(intReq.Cas)
	default:
		err = errors.New("Unknown Save request in dkv")
	}
	return
}

func (dr *dkvReplStore) Load(req []byte) ([]byte, error) {
	intReq := new(raftpb.InternalRaftRequest)
	if err := proto.Unmarshal(req, intReq); err != nil {
		return nil, err
	}
	switch {
	case intReq.Get != nil:
		return dr.get(intReq.Get)
	case intReq.MultiGet != nil:
		return dr.multiGet(intReq.MultiGet)
	default:
		return nil, errors.New("Unknown Load request in dkv")
	}
}

func (dr *dkvReplStore) put(putReq *serverpb.PutRequest) ([]byte, error) {
	err := dr.kvs.Put(&serverpb.KVPair{Key: putReq.Key, Value: putReq.Value, ExpireTS: putReq.ExpireTS})
	return nil, err
}

func (dr *dkvReplStore) multiPut(multiPutReq *serverpb.MultiPutRequest) ([]byte, error) {
	puts := make([]*serverpb.KVPair, len(multiPutReq.PutRequest))
	for i, request := range multiPutReq.PutRequest {
		puts[i] = &serverpb.KVPair{Key: request.Key, Value: request.Value, ExpireTS: request.ExpireTS}
	}
	err := dr.kvs.Put(puts...)
	return nil, err
}

func (dr *dkvReplStore) cas(casReq *serverpb.CompareAndSetRequest) ([]byte, error) {
	res, err := dr.kvs.CompareAndSet(casReq.Key, casReq.OldValue, casReq.NewValue)
	succ, fail := []byte{0}, []byte{1}
	if res && err == nil {
		return succ, nil
	}
	return fail, err
}

func (dr *dkvReplStore) delete(delReq *serverpb.DeleteRequest) ([]byte, error) {
	err := dr.kvs.Delete(delReq.Key)
	return nil, err
}

func (dr *dkvReplStore) get(getReq *serverpb.GetRequest) ([]byte, error) {
	vals, err := dr.kvs.Get(getReq.Key)
	if err != nil {
		return nil, err
	}
	return gobEncode(vals)
}

func (dr *dkvReplStore) multiGet(multiGetReq *serverpb.MultiGetRequest) ([]byte, error) {
	vals, err := dr.kvs.Get(multiGetReq.Keys...)
	if err != nil {
		return nil, err
	}
	return gobEncode(vals)
}

func gobEncode(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (dr *dkvReplStore) Close() error {
	return dr.kvs.Close()
}

// TODO: implement this correctly
func (dr *dkvReplStore) GetLastAppliedEntry() (db.RaftEntry, error) {
	return db.RaftEntry{}, errors.New("not implemented")
}

func (dr *dkvReplStore) Backup(_ db.SnapshotState) (io.ReadCloser, error) {
	return dr.kvs.GetSnapshot()
}

func (dr *dkvReplStore) Restore(data io.ReadCloser) error {
	return dr.kvs.PutSnapshot(data)
}
