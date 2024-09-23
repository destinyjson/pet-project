package handlers

import (
	"context"
	"fmt"
	"pet-project/internal/userService"
	"pet-project/internal/web/users"
)

type UserHandler struct {
	UsrService *userService.UserService
}

func NewUserHandler(usrService *userService.UserService) *UserHandler {
	return &UserHandler{
		UsrService: usrService,
	}
}

func (u *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.UsrService.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Name:     &usr.Name,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (u *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.User{
		Name:     *userRequest.Name,
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := u.UsrService.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Name:     &createdUser.Name,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

func (u *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id

	deletedUser, err := u.UsrService.DeleteUserByID(id)
	if err != nil {
		return nil, err
	}
	response := users.DeleteUsersId200JSONResponse{
		Id:       &deletedUser.ID,
		Name:     &deletedUser.Name,
		Email:    &deletedUser.Email,
		Password: &deletedUser.Password,
	}

	return response, nil
}

func (u *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id

	newUser := request.Body
	if newUser == nil {
		return nil, fmt.Errorf("body cannot be empty")
	}

	userToUpdate := userService.User{
		Name:     *newUser.Name,
		Email:    *newUser.Email,
		Password: *newUser.Password,
	}

	deletedUser, err := u.UsrService.UpdateUserByID(id, userToUpdate)
	if err != nil {
		return nil, err
	}
	response := users.PatchUsersId200JSONResponse{
		Id:       &deletedUser.ID,
		Name:     &deletedUser.Name,
		Email:    &deletedUser.Email,
		Password: &deletedUser.Password,
	}
	return response, nil
}
