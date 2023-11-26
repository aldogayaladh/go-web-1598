package products

var (
	QueryInsertProduct = `INSERT INTO storage.products(name,quantity,code_value,is_published,expiration,price)
	VALUES(?,?,?,?,?,?)`
	QueryGetAllProducts = `SELECT id, name, quantity, code_value,  is_published, expiration, price 
	FROM storage.products`
	QueryDeleteProduct  = `DELETE FROM storage.products WHERE id = ?`
	QueryGetProductById = `SELECT id, name, quantity, code_value, is_published, expiration, price
	FROM storage.products WHERE id = ?`
	QueryUpdateProduct = `UPDATE storage.products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ?
	WHERE id = ?`
)
