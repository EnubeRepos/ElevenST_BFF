package pix

type PixOpts struct {
	Key         string
	Name        string
	Amount      float64
	Description string
	City        string
}

type QRCodeOpts struct {
	Content string
	Size    int
}

type APIRequestQRCode struct {
	Name            string  `json:"name"`
	Amount          float64 `json:"amount"`
	City            string  `json:"city"`
	Description     string  `json:"description"`
	Transactionid   string  `json:"transactionId"`
	Pixkey          string  `json:"pixKey"`
	Foregroundcolor string  `json:"foregroundColor"`
	Backgroundcolor string  `json:"backgroundColor"`
}

type PathResp struct {
	Path string `json:"path"`
}
