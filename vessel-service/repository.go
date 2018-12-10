package main

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
	pb "shippy/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
)


const (
	DB_NAME        = "shippy"
	CON_COLLECTION = "vessel"
)

type Repository interface {
	FindAvailable(specification *pb.Specification) (*pb.Vessel, error)
	Create(v *pb.Vessel) error
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	// Find() 一般用来执行查询，如果想执行 select * 则直接传入 nil 即可
	// 通过 .All() 将查询结果绑定到 cons 变量上
	// 对应的 .One() 则只取第一行记录
	err := repo.collection().Find(bson.M{
			"capacity":  bson.M{"$gte": spec.Capacity},
			"maxweight": bson.M{"$bte": spec.MaxWeight},
		}).One(&vessel)

	if err != nil {
		return nil, errors.New("No vessel can't be use")
	}

	return vessel, nil
}

func (repo *VesselRepository) Create(v *pb.Vessel) error {
	return repo.collection().Insert(v)
}


func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(DB_NAME).C(CON_COLLECTION)
}
