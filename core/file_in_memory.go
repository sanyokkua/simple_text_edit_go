package core

type FileInMemory struct {
	Descriptor      FileInformation // Contains base file information
	OriginalContent string          // Content that was read after open
	ActualContent   string          // Content that can be changed during the time
}

func CreateFileInMemoryStruct(descriptor FileInformation, content string) FileInMemory {
	return FileInMemory{
		Descriptor:      descriptor,
		OriginalContent: content,
		ActualContent:   content,
	}
}

func (receiver *FileInMemory) GetDescriptor() *FileInformation {
	return &receiver.Descriptor
}

func (receiver *FileInMemory) SetActualContent(content string) {
	receiver.ActualContent = content
	if receiver.HasChanges() {
		receiver.Descriptor.IsChanged = true
	} else {
		receiver.Descriptor.IsChanged = false
	}
}

func (receiver *FileInMemory) SetOriginalContent(content string) {
	receiver.OriginalContent = content
	if receiver.HasChanges() {
		receiver.Descriptor.IsChanged = true
	} else {
		receiver.Descriptor.IsChanged = false
	}
}

func (receiver *FileInMemory) SetOpened(isOpened bool) {
	receiver.Descriptor.IsOpenedNow = isOpened
}

func (receiver *FileInMemory) HasChanges() bool {
	return receiver.OriginalContent != receiver.ActualContent
}

func CreateEmptyFileInMemory() FileInMemory {
	fileInfo := CreateFileInformationStruct("")
	file := CreateFileInMemoryStruct(fileInfo, "")
	return file
}

func CreateExistingFileInMemory(filePath string, fileContent string) FileInMemory {
	fileInfo := CreateFileInformationStruct(filePath)
	file := CreateFileInMemoryStruct(fileInfo, fileContent)
	return file
}
