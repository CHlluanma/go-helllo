namespace go example.shop.stock

include "base.thrift"

struct GetItemStockRequest {
    1: required i64 item_id
}

struct GetItemStockResponse {
    1:  i64 stock

    255: base.BaseResponse BaseResponse
}

service StockService {
    GetItemStockResponse GetItemStock(1: GetItemStockRequest req)
}
