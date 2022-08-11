from pymongo import MongoClient 
from fastapi import HTTPException
import json

from ..core.config import (
    database_name, 
    countries_collection
)
from ..core.utils import count_all_record

def search_one_spam_country(conn: MongoClient, word: str):
    row = list(conn[database_name][countries_collection].find({"$text": {"$search": word}}))
    return {
        "results": row,
        "total": count_all_record(conn, database_name, countries_collection)
    }

def get_all_spam_countries(conn: MongoClient, offset: str, size: str):
    row = list(conn[database_name][countries_collection].find({}).skip(offset).limit(size))
    return {
        "results": row,
        "total": count_all_record(conn, database_name, countries_collection)
    }

def get_spam_country(conn: MongoClient, country: str):
    row = list(conn[database_name][countries_collection].find({"_id": country}))
    return row

def create_spam_country(conn: MongoClient, country: str, value: int, code: str):
    if len(get_spam_country(conn, country)) > 0:
        raise HTTPException(status_code=404, detail="country exists")
    else:
        data = {
            "_id": country,
            "value": value,
            "code": code
        }
        conn[database_name][countries_collection].insert_one(data)
        return data

def update_spam_country(conn: MongoClient, country: str, value: str, code: str):
    check_spam_country = get_spam_country(conn, country) 
    if len(check_spam_country) == 0:
        raise HTTPException(status_code=404, detail="Spam country not found")
    else:
        value = value or list(check_spam_country)[0]["value"]
        code = code or list(check_spam_country)[0]["code"]
        update = {"value": value, "code": code}
        conn[database_name][countries_collection].update_one({"_id": country}, {"$set": update})
        return update

def detele_spam_country(conn: MongoClient, country: str):
    check_spam_country = get_spam_country(conn, country) 
    if len(check_spam_country) == 0:
        raise HTTPException(status_code=404, detail="Spam country not found")
    else:
        delete = {"_id": country}
        conn[database_name][countries_collection].delete_one(delete)
        return delete
            
def replace_spam_country(conn: MongoClient, data_pandas_object):
    payload = json.loads(data_pandas_object.to_json(orient='records'))
    conn[database_name][countries_collection].drop()
    conn[database_name][countries_collection].insert(payload)
    return {"data": conn[database_name][countries_collection].count_documents({})}