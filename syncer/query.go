package syncer

var _ QueryServer = &Query{}

type Query struct {
	UnimplementedQueryServer
}
