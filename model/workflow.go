package model

type Workflow struct {
	Path     string  `json:"path"`
	Size     int     `json:"size"`
	Modified float64 `json:"modified"`
}

type WorkflowDetail struct {
	Config interface{} `json:"config"`
	Extra  struct {
		Ds struct {
			Scale  float64   `json:"scale"`
			Offset []float64 `json:"offset"`
		} `json:"ds"`
		GroupNodes map[string]struct {
			//----------------
		} `json:"groupNodes"`
	} `json:"extra"`
	Groups     []interface{}   `json:"groups"`
	LastLinkID int             `json:"last_link_id"`
	LastNodeID int             `json:"last_node_id"`
	Links      [][]interface{} `json:"links"`
	Nodes      []struct {
		ID      int        `json:"id"`
		Type    string     `json:"type"`
		Flags   struct{}   `json:"flags"`
		Inputs  []struct{} `json:"inputs"`
		Mode    int        `json:"mode"`
		Order   int        `json:"order"`
		Outputs []struct {
			Name      string `json:"name"`
			Label     string `json:"label"`
			Links     []int  `json:"links"`
			SlotIndex int    `json:"slot_index"`
			Type      string `json:"type"`
		} `json:"outputs"`
		Pos           []float64              `json:"pos"`
		Properties    map[string]interface{} `json:"properties"`
		Size          []float64              `json:"size"`
		WidgetsValues []interface{}          `json:"widgets_values"`
	} `json:"nodes"`
	Version float64 `json:"version"`
}
