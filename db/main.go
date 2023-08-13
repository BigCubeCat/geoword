package db

func LoadDataBase(path string) *LevelDbDriver {
	ld := &LevelDbDriver{}
	ld.Init(path)
	return ld
}
