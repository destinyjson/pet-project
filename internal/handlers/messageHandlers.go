package handlers

import (
	"context"
	"fmt"
	"pet-project/internal/messageService"
	"pet-project/internal/web/messages"
)

type MessageHandler struct {
	MsgService *messageService.MessageService
}

func NewMessageHandler(msgService *messageService.MessageService) *MessageHandler {
	return &MessageHandler{
		MsgService: msgService,
	}
}

func (m *MessageHandler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	allMessages, err := m.MsgService.GetAllMessages()
	if err != nil {
		return nil, err
	}

	response := messages.GetMessages200JSONResponse{}

	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.Id,
			Message: &msg.Message,
		}
		response = append(response, message)
	}

	return response, nil
}

func (m *MessageHandler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	messageRequest := request.Body

	messageToCreate := messageService.RequestBody{Message: *messageRequest.Message}
	createdMessage, err := m.MsgService.CreateMessage(messageToCreate)
	if err != nil {
		return nil, err
	}

	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.Id,
		Message: &createdMessage.Message,
	}

	return response, nil
}

func (m *MessageHandler) DeleteMessagesId(_ context.Context, request messages.DeleteMessagesIdRequestObject) (messages.DeleteMessagesIdResponseObject, error) {
	id := request.Id

	deletedMessage, err := m.MsgService.DeleteMessageByID(id)
	if err != nil {
		return nil, err
	}

	response := messages.DeleteMessagesId200JSONResponse{
		Id:      &deletedMessage.Id,      // ID удалённого сообщения
		Message: &deletedMessage.Message, // Сообщение удалённого объекта
	}

	return response, nil
}

func (m *MessageHandler) PatchMessagesId(_ context.Context, request messages.PatchMessagesIdRequestObject) (messages.PatchMessagesIdResponseObject, error) {
	id := request.Id

	newMessage := request.Body
	if newMessage == nil || newMessage.Message == nil {
		return nil, fmt.Errorf("message cannot be empty")
	}

	messageToUpdate := messageService.RequestBody{Message: *newMessage.Message}

	deletedMessage, err := m.MsgService.UpdateMessageByID(id, messageToUpdate)
	if err != nil {
		return nil, err
	}
	response := messages.PatchMessagesId200JSONResponse{
		Id:      &deletedMessage.Id,
		Message: &deletedMessage.Message,
	}
	return response, nil
}
