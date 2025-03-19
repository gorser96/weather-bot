package services

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

// Структура для хранения данных о чате и топике
type ChatTopic struct {
	ChatID  int64 `json:"chat_id"`  // ID чата (группы)
	TopicID int   `json:"topic_id"` // ID топика
}

// Структура для хранения списка чатов и топиков
type ChatTopicList struct {
	Chats []ChatTopic `json:"chats"`
}

// Путь к файлу JSON
const filePath = "chat_topics.json"

// Функция для чтения данных из файла JSON
func ReadChatTopics() (ChatTopicList, error) {
	var chatTopics ChatTopicList

	// Проверяем, существует ли файл
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Если файл не существует, возвращаем пустой список
		return chatTopics, nil
	}

	// Читаем содержимое файла
	data, err := os.ReadFile(filePath)
	if err != nil {
		return chatTopics, fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	// Десериализуем JSON в структуру
	err = json.Unmarshal(data, &chatTopics)
	if err != nil {
		return chatTopics, fmt.Errorf("ошибка при десериализации JSON: %v", err)
	}

	return chatTopics, nil
}

// Функция для записи данных в файл JSON
func WriteChatTopics(chatTopics ChatTopicList) error {
	// Сериализуем структуру в JSON
	data, err := json.MarshalIndent(chatTopics, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка при сериализации JSON: %v", err)
	}

	// Записываем данные в файл
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %v", err)
	}

	return nil
}

// Функция для добавления нового чата и топика
func AddChatTopic(chatID int64, topicID int) error {
	// Читаем текущие данные
	chatTopics, err := ReadChatTopics()
	if err != nil {
		return err
	}

	hasInCache := slices.ContainsFunc(chatTopics.Chats, func(chat ChatTopic) bool { return chat.ChatID == chatID && chat.TopicID == topicID })
	if hasInCache {
		return nil
	}

	// Добавляем новый элемент
	newChatTopic := ChatTopic{
		ChatID:  chatID,
		TopicID: topicID,
	}
	chatTopics.Chats = append(chatTopics.Chats, newChatTopic)

	// Записываем обновленные данные в файл
	err = WriteChatTopics(chatTopics)
	if err != nil {
		return err
	}

	return nil
}

func RemoveChatTopic(chatID int64, topicID int) error {
	// Читаем текущие данные
	chatTopics, err := ReadChatTopics()
	if err != nil {
		return err
	}

	chats := slices.DeleteFunc(chatTopics.Chats, func(chat ChatTopic) bool { return chat.ChatID == chatID && chat.TopicID == topicID })

	chatTopics.Chats = chats
	// Записываем обновленные данные в файл
	err = WriteChatTopics(chatTopics)
	if err != nil {
		return err
	}

	return nil
}

// Функция для вывода всех чатов и топиков
func ListChatTopics() (*ChatTopicList, error) {
	// Читаем текущие данные
	chatTopics, err := ReadChatTopics()
	if err != nil {
		return nil, err
	}

	return &chatTopics, nil
}
