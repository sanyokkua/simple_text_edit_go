package types

import (
	"reflect"
	"testing"
)

func TestButton_EqualTo(t *testing.T) {
	type args struct {
		another Button
	}
	tests := []struct {
		name string
		r    Button
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.EqualTo(tt.args.another); got != tt.want {
				t.Errorf("EqualTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtensionsMap_GetByExtension(t *testing.T) {
	type args struct {
		ext FileTypeExtension
	}
	tests := []struct {
		name    string
		r       ExtensionsMap
		args    args
		want    *FileTypesJsonStruct
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetByExtension(tt.args.ext)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByExtension() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByExtension() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilesMap_Add(t *testing.T) {
	type args struct {
		fileStruct *FileStruct
	}
	tests := []struct {
		name    string
		r       FilesMap
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Add(tt.args.fileStruct); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFilesMap_GetById(t *testing.T) {
	type args struct {
		fileId int64
	}
	tests := []struct {
		name    string
		r       FilesMap
		args    args
		want    *FileStruct
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetById(tt.args.fileId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilesMap_IsPathPresentInMap(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		r    FilesMap
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.IsPathPresentInMap(tt.args.path); got != tt.want {
				t.Errorf("IsPathPresentInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilesMap_Remove(t *testing.T) {
	type args struct {
		fileId int64
	}
	tests := []struct {
		name    string
		r       FilesMap
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Remove(tt.args.fileId); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTypesMap_GetByTypeKey(t *testing.T) {
	type args struct {
		ext FileTypeKey
	}
	tests := []struct {
		name    string
		r       TypesMap
		args    args
		want    *FileTypesJsonStruct
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetByTypeKey(tt.args.ext)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByTypeKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByTypeKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
