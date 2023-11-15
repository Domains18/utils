package dbs 

type FindOption interface {
	apply(*option)
}

type option struct {
	query []Query
	order any
	offset int
	limit int
	preloads []string
}

type 
