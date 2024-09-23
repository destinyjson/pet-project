package service

type MessageService struct {
	repo MessageRepository
}

func NewService(repo MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(message RequestBody) (RequestBody, error) {
	return s.repo.CreateMessage(message)
}

func (s *MessageService) GetAllMessages() ([]RequestBody, error) {
	return s.repo.GetAllMessages()
}

func (s *MessageService) DeleteMessageByID(id int) (RequestBody, error) {
	return s.repo.DeleteMessageByID(id)
}

func (s *MessageService) UpdateMessageByID(id int, message RequestBody) (RequestBody, error) {
	return s.repo.UpdateMessageByID(id, message)
}
