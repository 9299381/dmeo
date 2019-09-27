package constants

const Err string = "9999::错误"
const ErrStop string = "9990::暂停服务"
const ErrConvert string = "9000::类型转换错误"
const ErrJson string = "9001::JSON报文解析错误"
const ErrRoute string = "9010::路由解析错误"

const ErrNoToken string = "9020::缺少authToken"
const ErrTokenFmt string = "9021::Token格式错误"
const ErrTokenExp string = "9022::Token过期,请重新登录"
const ErrTokenSign string = "9023::签名错误,请重新登录"