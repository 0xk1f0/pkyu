package internal

import (
	"io"
	"os"

	"github.com/goccy/go-yaml"
)

type Pod struct {
	Spec Spec `yaml:"spec"`
}

type Spec struct {
	Containers []Container `yaml:"containers"`
}

type Container struct {
	Image string `yaml:"image"`
}

func ReadKubefile(path string) ([]byte, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return []byte{}, err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return []byte{}, err
	}

	return content, nil
}

func GetKubefileImages(content []byte) ([]string, error) {
	var pod Pod
	var images []string
	if err := yaml.Unmarshal(content, &pod); err != nil {
		return []string{}, err
	}

	for _, container := range pod.Spec.Containers {
		images = append(images, container.Image)
	}

	return images, nil
}
