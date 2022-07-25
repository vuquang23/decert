from pymongo import MongoClient 
from fastapi import HTTPException
from typing import Optional

from ..core.config import (
    database_name,
    team_email_collections
)
from ..core.utils import count_all_record, create_team_email_regex

def get_all_team_email(conn: MongoClient, offset: str, size: str):
    row = list(conn[database_name][team_email_collections].find().skip(offset).limit(size))
    return {
        "results": row,
        "total": count_all_record(conn, database_name, team_email_collections)
    }

def get_one_team_email(conn: MongoClient, _id: str):
    row = list(conn[database_name][team_email_collections].find({"_id": _id}))
    return row

def create_one_team_email(conn: MongoClient, _id: str):
    checker = get_one_team_email(conn, _id)
    if len(checker) > 0:
        raise HTTPException(status_code=404, detail="Team email pattern exists")
    else:
        create = {"_id": _id, "pattern": create_team_email_regex(_id)}
        conn[database_name][team_email_collections].insert_one(create)
        return create

def update_one_team_email(conn: MongoClient, item: dict):
    value = item.dict()
    checker = get_one_team_email(conn, value["id"])
    if len(checker) == 0:
        raise HTTPException(status_code=404, detail="Team email pattern don't exists")
    else:
        pattern = value["pattern"] or list(checker)[0]["pattern"]
        update = {"pattern": pattern}
        conn[database_name][team_email_collections].update_one({"_id": value["id"]}, {"$set": update})
        update["_id"] = value["id"]
        return update

def delete_one_team_email(conn: MongoClient, _id: str):
    checker = get_one_team_email(conn, _id)
    if len(checker) == 0:
        raise HTTPException(status_code=404, detail="Team email pattern not found")
    else:
        delete = {"_id": _id}
        conn[database_name][team_email_collections].delete_one(delete)
        return delete




