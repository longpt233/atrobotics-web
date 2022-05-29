package response

type ListProductAllCategoryResponse struct {
	CategoryName string            `json:"categoryName"`
	ProductList  []ProductResponse `json:"productList"`
}
