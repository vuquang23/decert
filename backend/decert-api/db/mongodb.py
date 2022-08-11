from pymongo import MongoClient
from ..core.config import MONGODB_URI

class DataBase:
    def __init__(self):
        self.connection_string = MONGODB_URI

    def connect(self):
        client = MongoClient(self.connection_string)
        return client

client = DataBase()
db = client.connect()

def get_database():
    return db