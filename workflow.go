package comfy4go

import (
	"encoding/json"
	"github.com/iazkaban/comfy4go/model"
	"os"
)

func ReadWorkflowFromFile(filename string) (*model.WorkflowDetail, error) {
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ReadWorkflow(body)
}

func ReadWorkflow(body []byte) (*model.WorkflowDetail, error) {
	rs := &model.WorkflowDetail{}
	err := json.Unmarshal(body, rs)
	if err != nil {
		return nil, err
	}

	return rs, nil
}
