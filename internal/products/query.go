package products

var (
	QueryInsertProduct = `INSERT INTO storage.product(name,quantity,code_value,is_published,expiration,price)
	VALUES(?,?,?,?,?,?)`
	QueryGetAllProducts = `SELECT id, name, quantity, code_value,  is_published, expiration, price 
	FROM storage.product`
	QueryDeleteProduct  = `DELETE FROM storage.product WHERE id = ?`
	QueryGetProductById = `SELECT id, name, quantity, code_value, is_published, expiration, price
	FROM storage.product WHERE id = ?`
	QueryUpdateProduct = `UPDATE storage.product SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ?
	WHERE id = ?`
)
