import stock_pb2_grpc
import stock_pb2
import grpc
from investing import get_data


def get_stock_price():
    try:
        data = get_data()
    except:
        print("Error")
        #data = []

    if len(data) < 1:
        stock_request = stock_pb2.StockPrice(
            symbol="ERROR", last="ERROR")
        yield stock_request

    for _, value in data.iterrows():
        stock_request = stock_pb2.StockPrice(
            symbol="AAPL", last=str(value.Close))
        yield stock_request


def run():
    with grpc.insecure_channel('server:50051') as channel:
        stub = stock_pb2_grpc.SaveSharePriceStub(channel)

        server_response = stub.StoreStockDatabase(get_stock_price())
        print("Server Response Received:", server_response)


if __name__ == "__main__":
    run()
