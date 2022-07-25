from pymongo import MongoClient
from ..model.user import UserInUpdate

from ..core.config import users_database_name, users_collection
from ..core.security import get_password_hash

def get_all_user(conn: MongoClient):
    row = conn[users_database_name][users_collection].find({}, {"_id": 0})
    return list(row)

def get_user_by_username(conn: MongoClient, username: str):
    row = conn[users_database_name][users_collection].find({"username": username}, {"_id": 0})
    return list(row)

def create_user(conn: MongoClient, user: dict):
    data = user.dict()
    data["password"] = get_password_hash(data["password"])
    insert = data
    conn[users_database_name][users_collection].insert_one(insert)
    return data

def update_user(conn: MongoClient, info: UserInUpdate):
    dbuser = get_user_by_username(conn, info.username)

    dbuser[0]["password"] = info.password or dbuser[0]["password"]
    dbuser[0]["role"] = info.role or dbuser[0]["role"]

    if info.password:
        dbuser[0]["password"] = get_password_hash(dbuser[0]["password"])

    update = dbuser[0]
    conn[users_database_name][users_collection].update_one({"username": info.username}, {"$set": update})
    return dbuser

def delete_user(conn: MongoClient, username: str):
    delete = {"username": username}
    data = delete
    conn[users_database_name][users_collection].delete_one(delete)
    return data