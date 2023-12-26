package encoding

import (
	"encoding/json"
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
	"os"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// Чтение данных в JSON формате из файла
	jsonData, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}

	// Десериализация данных JSON в структуру
	var dockerCompose models.DockerCompose
	if err := json.Unmarshal(jsonData, &dockerCompose); err != nil {
		return err
	}

	// Сериализация структуры в YAML
	yamlData, err := yaml.Marshal(dockerCompose)
	if err != nil {
		return err
	}

	// Запись JSON в файл
	if err := os.WriteFile(j.FileOutput, yamlData, 0644); err != nil {
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Чтение данных из файла
	yamlData, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}

	// Десериализация данных YAML в структуру
	var dockerCompose models.DockerCompose
	if err := yaml.Unmarshal(yamlData, &dockerCompose); err != nil {
		return err
	}

	// Сериализация структуры в JSON
	jsonData, err := json.Marshal(dockerCompose)
	if err != nil {
		return err
	}

	// Запись JSON в файл
	if err := os.WriteFile(y.FileOutput, jsonData, 0644); err != nil {
		return err
	}

	return nil
}
