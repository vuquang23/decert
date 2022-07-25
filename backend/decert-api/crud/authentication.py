from pymongo import MongoClient

from ..core.config import users_database_name, users_collection
from ..core.security import get_password_hash

def get_user_by_username(conn: MongoClient, username: str):
    row = conn[users_database_name][users_collection].find({"username": username}, {"_id": 0})
    return list(row)

def create_user(conn: MongoClient, user: dict):
    data = user.dict()
    data["password"] = get_password_hash(data["password"])
    data["role"] = "normal"
    insert = data
    conn[users_database_name][users_collection].insert_one(insert)
    return data
