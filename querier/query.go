package querier

//import "github.com/we-and/weand_backend_common/app"

type Querier struct {
	//H interface{}
}

func BuildQuerier(h interface{}) Querier {
	q := Querier{}
	//	q.SetHandler(h)
	return q
}

//func (q *Querier) getHandler() app.AppHandler {
//	return q.H.(app.AppHandler)
//}

//func (q *Querier) SetHandler(h interface{}) {
//	q.H = h
//}
