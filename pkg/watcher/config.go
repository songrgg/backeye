package watcher

type Config struct {
	Points []PointConfig `json:"points"`
}

type PointConfig struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Type string `json:"type"`
	Conf string `json:"conf"`
}
