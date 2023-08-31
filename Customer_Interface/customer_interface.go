package customerinterface

import customermodel "github.com/GURUAKASHSM/netxd_dal/Customer_Model"

type ICustomer interface{
	CreateCustomer(customer *customermodel.Customer)(*customermodel.CustomerResponse, error)
}
