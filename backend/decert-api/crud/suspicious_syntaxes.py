from pymongo import MongoClient 
from fastapi import HTTPException
from bson import ObjectId

from ..core.config import (
    database_name, 
    suspicius_syntaxes_collection
)
from ..model.suspicious_syntaxes import (
    SuspiciousSyntaxInCreate
)
from ..core.utils import count_all_record

def get_all_suspicious_syntaxes(conn: MongoClient, offset: str, size: str):
    row = list(conn[database_name][suspicius_syntaxes_collection].find().skip(offset).limit(size))
    for i in row:
        i["_id"] = str(i["_id"])
    data = {
        "results": row,
        "total": count_all_record(conn, database_name, suspicius_syntaxes_collection)
    }
    return data

def get_suspicious_syntax(conn: MongoClient, _id: str):
    row = list(conn[database_name][suspicius_syntaxes_collection].find({"_id": ObjectId(_id)}))
    return row

def create_suspicious_syntax(conn: MongoClient, item: dict):
    sus_syntax = item.dict()
    if len(list(conn[database_name][suspicius_syntaxes_collection].find({"pattern": sus_syntax["pattern"]}))) > 0:
        raise HTTPException(status_code=404, detail="Suspicious syntax exists")
    else:
        data = {**item.dict()}
        conn[database_name][suspicius_syntaxes_collection].insert_one(data)
        return data

def update_suspicious_syntax(conn: MongoClient, _id: str, item: dict):
    checker = get_suspicious_syntax(conn, _id)
    if len(checker) == 0:
        raise HTTPException(status_code=404, detail="Suspicious syntax not found")
    else:
        value = item.dict()
        name = value["name"] or list(checker)[0]["name"]
        pattern = value["pattern"] or list(checker)[0]["pattern"]

        update = {"name": name, "pattern": pattern}
        conn[database_name][suspicius_syntaxes_collection].update_one({"_id": ObjectId(_id)}, {"$set": update})
        return update

def delete_suspicious_syntax(conn: MongoClient, _id: str):
    checker = get_suspicious_syntax(conn, _id)
    if len(checker[0]) == 0:
        raise HTTPException(status_code=404, detail="suspicious syntax not found")
    else:
        delete = {"_id": ObjectId(_id)}
        conn[database_name][suspicius_syntaxes_collection].delete_one(delete)
        return delete

    