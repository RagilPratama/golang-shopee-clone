package models

type Product struct {
	ID                 string  `json:"id" gorm:"primaryKey"`
	Title              string  `json:"title"`
	Jenis              string  `json:"jenis"`
	Image              string  `json:"image"`
	Jumlah             int     `json:"jumlah"`
	HargaSatuan        float64 `json:"harga_satuan"`
	HargaPromo         float64 `json:"harga_promo"`
	TotalJumlah        int     `json:"total_jumlah"`
	TotalHarga         float64 `json:"total_harga"`
	EstimasiPengiriman int     `json:"estimasi_pengiriman"`
	ShopName           string  `json:"shop_name"`
	Badge              string  `json:"badge"`
	Live               bool    `json:"live"`
	Status             string  `json:"status"`
	CoinRewardText     string  `json:"coin_reward_text"`
	VariantLabel       string  `json:"variant_label"`
}
