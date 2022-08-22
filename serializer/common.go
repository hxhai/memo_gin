package serializer

//基础序列化
type Response struct {
	Status int         `json:"status"` //状态
	Data   interface{} `json:"data"`   //数据
	Msg    string      `json:"msg"`    //返回信息
	Error  string      `json:"error"`  //错误信息
}

//TokenData带有token的data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

//带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

//带总数的返回
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}
