package request


type BannerRequest struct {
	BannerProductId string `json:"banner_product_id"`
	BannerImage     string `json:"banner_image"`
}