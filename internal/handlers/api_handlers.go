package handlers

import (
	"context"
	"fmt"
	"pet-project/internal/service"
	"pet-project/internal/web/messages"
)

type Handler struct {
	Service *service.MessageService
}

func NewHandler(service *service.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	// Получение всех сообщений из сервиса
	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := messages.GetMessages200JSONResponse{}

	// Заполняем слайс response всеми сообщениями из БД
	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.Id,
			Message: &msg.Message,
		}
		response = append(response, message)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	messageRequest := request.Body
	// Обращаемся к сервису и создаем сообщение
	messageToCreate := service.RequestBody{Message: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)
	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.Id,
		Message: &createdMessage.Message,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *Handler) DeleteMessagesId(_ context.Context, request messages.DeleteMessagesIdRequestObject) (messages.DeleteMessagesIdResponseObject, error) {
	// Получаем ID из запроса
	id := request.Id

	// Вызываем сервис для удаления сообщения по ID
	deletedMessage, err := h.Service.DeleteMessageByID(id)
	if err != nil {
		return nil, err
	}

	// Возвращаем успешный ответ. Если требуется, можно включить удалённое сообщение в ответ.
	response := messages.DeleteMessagesId200JSONResponse{
		Id:      &deletedMessage.Id,      // ID удалённого сообщения
		Message: &deletedMessage.Message, // Сообщение удалённого объекта
	}

	return response, nil
}

func (h *Handler) PatchMessagesId(_ context.Context, request messages.PatchMessagesIdRequestObject) (messages.PatchMessagesIdResponseObject, error) {
	id := request.Id

	newMessage := request.Body
	if newMessage == nil || newMessage.Message == nil {
		return nil, fmt.Errorf("message cannot be empty")
	}

	messageToUpdate := service.RequestBody{Message: *newMessage.Message}

	deletedMessage, err := h.Service.UpdateMessageByID(id, messageToUpdate)
	if err != nil {
		return nil, err
	}
	response := messages.PatchMessagesId200JSONResponse{
		Id:      &deletedMessage.Id,
		Message: &deletedMessage.Message,
	}
	return response, nil
}
