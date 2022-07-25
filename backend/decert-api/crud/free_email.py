from pymongo import MongoClient 
from typing import Optional
from fastapi import HTTPException

from ..core.config import (
    database_name, 
    free_collection,
    well_known_companies_collection,
    disposable_collection
)
from ..core.utils import create_domain_regex, count_all_record
from .utils import check_domain_duplicate

def create_free_email(conn: MongoClient, name: str, domain: str):

    check_domain_duplicate(conn, domain, well_known_companies_collection, disposable_collection, free_collection)

    pattern = create_domain_regex(domain)
    data = {
        "_id": domain,
        "pattern": pattern,
        "name": name
    }
    conn[database_name][free_collection].insert_one(data)
    return data

def get_free_email(conn: MongoClient, offset: str, size: str):
    row = list(conn[database_name][free_collection].find().skip(offset).limit(size))
    return {
        "results": row,
        "total": count_all_record(conn, database_name, free_collection)
    }

def get_one_free_email(conn: MongoClient, domain: str):
    row = list(conn[database_name][free_collection].find({"_id": domain}))
    return row

def update_free_email(conn: MongoClient, domain: str, name: str, pattern: Optional[str]=None):
    check_free_email = get_one_free_email(conn, domain)
    if len(check_free_email) == 0:
        raise HTTPException(status_code=404, detail="domain not found")
    else:
        name = name or list(check_free_email)[0]["name"]
        pattern = pattern or list(check_free_email)[0]["pattern"]
        update = {"name": name, "pattern": pattern}
        conn[database_name][free_collection].update_one({"_id": domain}, {"$set": update})
        return update

def delete_free_email(conn: MongoClient, domain: str):
    delete = {"_id": domain}
    conn[database_name][free_collection].delete_one(delete)
    return delete
