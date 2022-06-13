package enums

// STATUS
const (
	STATUS_CREATED string = "created"
	STATUS_PENDING string = "pendding"
	STATUS_PAID    string = "paid"
	STATUS_ABORT   string = "abort"
	STATUS_SUCCESS string = "success"
)

var STATUS = map[int]string{
	1: STATUS_CREATED,
	2: STATUS_PENDING,
	3: STATUS_PAID,
	4: STATUS_ABORT,
	5: STATUS_SUCCESS, 
}
