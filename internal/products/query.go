package products

var (
	QueryInsertProduct = `INSERT INTO storage.products(name,quantity,code_value,is_published,expiration,price)
	VALUES(?,?,?,?,?,?)`
)
