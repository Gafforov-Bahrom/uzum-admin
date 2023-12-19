package admin

type Product struct {
	Id          uint64 `json:"-"`
	Name        string `json:"name"`
	Description string `json:"-"`
	Count       uint64 `json:"-"`
	Price       uint64 `json:"price"`
}

type GetProductsResponse struct {
	Count    uint64     `json:"count"`
	Products []*Product `json:"products"`
}

type AddProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Count       uint64 `json:"count"`
	Price       uint64 `json:"price"`
}

type UpdateProductRequest struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Count       uint64 `json:"count"`
	Price       uint64 `json:"price"`
}

type CancelOrderRequest struct {
	OrderId uint64 `json:"user_id"`
}

type GetUserRoleOut struct {
	Id   uint64
	Role string
}
