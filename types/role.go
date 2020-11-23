package types

type Role struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Summary string `json:"summary"`
	Skills  []struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"skills"`
}

type RoleVM struct {
	ID      int    `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}
