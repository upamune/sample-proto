package main

import (
	"github.com/upamune/sample-proto/proto"
	"golang.org/x/net/context"
)

type UserServer struct {
	store map[string]User
}

type User struct {
	Name string
	Age  int32
	Sex  proto.Sex
}

func (u User) ToProto() *proto.User {
	return &proto.User{
		Name: u.Name,
		Age:  u.Age,
		Sex:  u.Sex,
	}
}

func (s UserServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	u := User{
		Name: req.Name,
		Age:  req.Age,
		Sex:  req.Sex,
	}

	if _, ok := s.store[u.Name]; ok {
		return &proto.CreateResponse{
			Status:  proto.Status_FAIL,
			Message: "user already exists",
		}, nil
	}

	s.store[u.Name] = u

	return &proto.CreateResponse{
		Status:  proto.Status_OK,
		Message: "success",
	}, nil
}

func (s UserServer) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	userName := req.Name
	u, ok := s.store[userName]
	if !ok {
		return &proto.GetResponse{
			Status:  proto.Status_FAIL,
			Message: "no such user exists",
		}, nil
	}

	return &proto.GetResponse{
		Status:  proto.Status_OK,
		Message: "success",
		User:    u.ToProto(),
	}, nil
}

func (s UserServer) GetAll(ctx context.Context, req *proto.GetAllRequest) (*proto.GetAllResponse, error) {
	users := []*proto.User{}

	for _, user := range s.store {
		users = append(users, user.ToProto())
	}

	return &proto.GetAllResponse{
		Status:  proto.Status_OK,
		Message: "success",
		Users:   users,
	}, nil
}

func (s UserServer) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	userName := req.Name

	u, ok := s.store[userName]
	if !ok {
		return &proto.UpdateResponse{
			Status:  proto.Status_FAIL,
			Message: "no such user exists",
		}, nil
	}

	u.Age = req.Age
	u.Sex = req.Sex

	s.store[userName] = u

	return &proto.UpdateResponse{
		Status:  proto.Status_OK,
		Message: "success",
	}, nil
}

func (s UserServer) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	userName := req.Name

	_, ok := s.store[userName]
	if !ok {
		return &proto.DeleteResponse{
			Status:  proto.Status_FAIL,
			Message: "no such user exists",
		}, nil
	}

	delete(s.store, userName)

	return &proto.DeleteResponse{
		Status:  proto.Status_OK,
		Message: "success",
	}, nil
}
