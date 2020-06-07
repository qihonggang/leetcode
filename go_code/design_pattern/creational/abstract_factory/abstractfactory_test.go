package abstract_factory

import "testing"

func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func TestRdbFactory(t *testing.T) {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetail(factory)
}

func TestXMLFactory(t *testing.T) {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	getMainAndDetail(factory)
}
