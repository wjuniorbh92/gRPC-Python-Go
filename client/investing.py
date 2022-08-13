import yfinance as yf


def get_data():
    """
    This function gets the data from the API and returns it.
    """
    data = yf.download("AAPL", period="7d", interval="1m")

    return data
