package permission

type Permission struct {
	Id        int64
	Name      string // resource name
	Actions   string // 1:add 2:del 3:select 4:update
	Creator   string
	CreatedAt string
}
