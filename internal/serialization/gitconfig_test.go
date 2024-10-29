package serialization_test

import (
	"reflect"
	"testing"
	"wrynnhall/gitgudr/internal/serialization"
)

func TestDeserialize_GitgudrConfig(t *testing.T) {
	data := `
[gitgudr]
ignore_repo = repo1
ignore_repo = repo2
directory = dir1
directory = dir2
`

	expected := serialization.GitgudrConfig{
		IgnoreRepos: []string{"repo1", "repo2"},
		Directories: []string{"dir1", "dir2"},
	}

	var result serialization.GitgudrConfig
	err := serialization.Deserialize(data, &result)
	if err != nil {
		t.Fatalf("Deserialize failed: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, result)
	}
}

func TestDeserialize_WiprConfig(t *testing.T) {
	data := `
[gitgudr]
ignore_repo = repo1
ignore_repo = repo2
directory = dir1
directory = dir2

[gitgudr_wipr]
ignore_repo = repo3
directory = dir3
`

	expected := serialization.WiprConfig{
		IgnoreRepos: []string{"repo3"},
		Directories: []string{"dir3"},
	}

	var result serialization.WiprConfig
	err := serialization.Deserialize(data, &result)
	if err != nil {
		t.Fatalf("Deserialize failed: %v", err)
	}

	if !reflect.DeepEqual(result.IgnoreRepos, expected.IgnoreRepos) {
		t.Errorf("Expected IgnoreRepos %+v, but got %+v", expected.IgnoreRepos, result.IgnoreRepos)
	}

	if !reflect.DeepEqual(result.Directories, expected.Directories) {
		t.Errorf("Expected Directories %+v, but got %+v", expected.Directories, result.Directories)
	}
}

func TestDeserialize_EmptyConfig(t *testing.T) {
	data := ``

	expected := serialization.GitgudrConfig{
		IgnoreRepos: nil,
		Directories: nil,
	}

	var result serialization.GitgudrConfig
	err := serialization.Deserialize(data, &result)
	if err != nil {
		t.Fatalf("Deserialize failed: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, result)
	}
}

func TestDeserialize_InvalidFieldType(t *testing.T) {
	data := `
[gitgudr]
ignore_repo = repo1
`

	type InvalidConfig struct {
		IgnoreRepos int `gitgudr:"ignore_repo"`
	}

	var result InvalidConfig
	err := serialization.Deserialize(data, &result)
	if err == nil {
		t.Fatal("Expected an error due to unsupported field type, but got none")
	}
}

func TestDeserialize_UnknownSection(t *testing.T) {
	data := `
[unknown_section]
ignore_repo = repo1
`

	var result serialization.GitgudrConfig
	err := serialization.Deserialize(data, &result)
	if err != nil {
		t.Fatalf("Deserialize failed: %v", err)
	}

	if len(result.IgnoreRepos) != 0 || len(result.Directories) != 0 {
		t.Errorf("Expected empty fields, but got %+v", result)
	}
}
