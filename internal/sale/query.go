package sale

var (
	QuerySelectSaleById         = `SELECT id, description, id_seller, id_product FROM storage.sale WHERE id = ?`
	QuerySelectSaleByCodSeller  = `SELECT id, description, id_seller, id_product FROM storage.sale WHERE id_seller = ?`
	QuerySelectSaleByCodProduct = `SELECT id, description, id_seller, id_product FROM storage.sale WHERE id_product = ?`
	QueryInsertSale             = `INSERT INTO storage.sale(description,id_seller,id_product) VALUES(?,?,?)`

	// Pueden armar querys personalizadas para cada caso de uso.
	// Query to get all sales and info about seller.
	QueryGetAllSales = `SELECT s.id, s.description, s.id_seller, s.id_product, sl.cod_seller, sl.description, sl.is_authorization 
	FROM storage.sale INNER JOIN storage.seller sl ON s.id_seller = sl.id`
	// Query to get all sales and info about product.
	QueryGetAllSalesProduct = `SELECT s.id, s.description, s.id_seller, s.id_product, p.name, p.quantity, p.code_value, p.is_published, p.expiration, p.price
	FROM storage.sale INNER JOIN storage.product p ON s.id_product = p.id`

	// Add all querys.
)
