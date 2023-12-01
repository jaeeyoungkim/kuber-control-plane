package main

type CommonResponse[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
}

func main() {
	//metrics := lib.GetApiserverRequestTotalMetrics()
	//c := CommonResponse[[]lib.ApiServerRequestTotal]{
	//	Data:    metrics,
	//	Message: "this is ok",
	//}
	//marshal, err := json.Marshal(c)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(marshal))
}
