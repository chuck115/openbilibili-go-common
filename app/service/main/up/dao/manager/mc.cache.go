// Code generated by $GOPATH/src/go-common/app/tool/cache/mc. DO NOT EDIT.

/*
  Package manager is a generated mc cache package.
  It is generated from:
  type _mc interface {
		// mc: -key=upSpecialCacheKey -expire=d.upSpecialExpire -encode=pb
		AddCacheUpSpecial(c context.Context, mid int64, us *upgrpc.UpSpecial) (err error)
		// mc: -key=upSpecialCacheKey
		CacheUpSpecial(c context.Context, mid int64) (res *upgrpc.UpSpecial, err error)
		// mc: -key=upSpecialCacheKey
		DelCacheUpSpecial(c context.Context, mid int64) (err error)
		// mc: -key=upSpecialCacheKey -expire=d.upSpecialExpire -encode=pb
		AddCacheUpsSpecial(c context.Context, mu map[int64]*upgrpc.UpSpecial) (err error)
		// mc: -key=upSpecialCacheKey
		CacheUpsSpecial(c context.Context, mid []int64) (res map[int64]*upgrpc.UpSpecial, err error)
		// mc: -key=upSpecialCacheKey
		DelCacheUpsSpecial(c context.Context, mids []int64) (err error)
	}
*/

package manager

import (
	"context"
	"fmt"

	upgrpc "go-common/app/service/main/up/api/v1"
	"go-common/library/cache/memcache"
	"go-common/library/log"
	"go-common/library/stat/prom"
)

var _ _mc

// AddCacheUpSpecial Set data to mc
func (d *Dao) AddCacheUpSpecial(c context.Context, id int64, val *upgrpc.UpSpecial) (err error) {
	if val == nil {
		return
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	key := upSpecialCacheKey(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.upSpecialExpire, Flags: memcache.FlagProtobuf}
	if err = conn.Set(item); err != nil {
		prom.BusinessErrCount.Incr("mc:AddCacheUpSpecial")
		log.Errorv(c, log.KV("AddCacheUpSpecial", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// CacheUpSpecial get data from mc
func (d *Dao) CacheUpSpecial(c context.Context, id int64) (res *upgrpc.UpSpecial, err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := upSpecialCacheKey(id)
	reply, err := conn.Get(key)
	if err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:CacheUpSpecial")
		log.Errorv(c, log.KV("CacheUpSpecial", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	res = &upgrpc.UpSpecial{}
	err = conn.Scan(reply, res)
	if err != nil {
		prom.BusinessErrCount.Incr("mc:CacheUpSpecial")
		log.Errorv(c, log.KV("CacheUpSpecial", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DelCacheUpSpecial delete data from mc
func (d *Dao) DelCacheUpSpecial(c context.Context, id int64) (err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := upSpecialCacheKey(id)
	if err = conn.Delete(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:DelCacheUpSpecial")
		log.Errorv(c, log.KV("DelCacheUpSpecial", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheUpsSpecial Set data to mc
func (d *Dao) AddCacheUpsSpecial(c context.Context, values map[int64]*upgrpc.UpSpecial) (err error) {
	if len(values) == 0 {
		return
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	for id, val := range values {
		key := upSpecialCacheKey(id)
		item := &memcache.Item{Key: key, Object: val, Expiration: d.upSpecialExpire, Flags: memcache.FlagProtobuf}
		if err = conn.Set(item); err != nil {
			prom.BusinessErrCount.Incr("mc:AddCacheUpsSpecial")
			log.Errorv(c, log.KV("AddCacheUpsSpecial", fmt.Sprintf("%+v", err)), log.KV("key", key))
			return
		}
	}
	return
}

// CacheUpsSpecial get data from mc
func (d *Dao) CacheUpsSpecial(c context.Context, ids []int64) (res map[int64]*upgrpc.UpSpecial, err error) {
	l := len(ids)
	if l == 0 {
		return
	}
	keysMap := make(map[string]int64, l)
	keys := make([]string, 0, l)
	for _, id := range ids {
		key := upSpecialCacheKey(id)
		keysMap[key] = id
		keys = append(keys, key)
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	replies, err := conn.GetMulti(keys)
	if err != nil {
		prom.BusinessErrCount.Incr("mc:CacheUpsSpecial")
		log.Errorv(c, log.KV("CacheUpsSpecial", fmt.Sprintf("%+v", err)), log.KV("keys", keys))
		return
	}
	for key, reply := range replies {
		var v *upgrpc.UpSpecial
		v = &upgrpc.UpSpecial{}
		err = conn.Scan(reply, v)
		if err != nil {
			prom.BusinessErrCount.Incr("mc:CacheUpsSpecial")
			log.Errorv(c, log.KV("CacheUpsSpecial", fmt.Sprintf("%+v", err)), log.KV("key", key))
			return
		}
		if res == nil {
			res = make(map[int64]*upgrpc.UpSpecial, len(keys))
		}
		res[keysMap[key]] = v
	}
	return
}

// DelCacheUpsSpecial delete data from mc
func (d *Dao) DelCacheUpsSpecial(c context.Context, ids []int64) (err error) {
	if len(ids) == 0 {
		return
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	for _, id := range ids {
		key := upSpecialCacheKey(id)
		if err = conn.Delete(key); err != nil {
			if err == memcache.ErrNotFound {
				err = nil
				continue
			}
			prom.BusinessErrCount.Incr("mc:DelCacheUpsSpecial")
			log.Errorv(c, log.KV("DelCacheUpsSpecial", fmt.Sprintf("%+v", err)), log.KV("key", key))
			return
		}
	}
	return
}