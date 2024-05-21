package outputs

type WebResponse struct {
	BaseOutput
	Data interface{} `json:"data"`
}
