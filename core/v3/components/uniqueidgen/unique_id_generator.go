package uniqueidgen

import (
	"simple_text_editor/core/v3/types"
	"time"
)

type GeneratorStruct struct {
	previousId int64
}

func (r *GeneratorStruct) GenerateId() int64 {
	currentId := time.Now().UnixNano()
	if r.previousId == currentId { // If new ID was requested immediately, sleep a little and generate again
		time.Sleep(500 * time.Millisecond)
		currentId = time.Now().UnixNano()
	}
	r.previousId = currentId
	return currentId
}

func CreateIUniqueIdGenerator() types.IUniqueIdGenerator {
	return &GeneratorStruct{}
}
