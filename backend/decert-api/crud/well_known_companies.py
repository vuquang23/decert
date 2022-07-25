import json
from pymongo import MongoClient 
from fastapi import HTTPException

from ..core.config import (
    database_name, 
    well_known_companies_collection,
    disposable_collection,
    free_collection
)
from ..core.utils import create_domain_regex, count_all_record
from .utils import check_domain_duplicate

def get_all_companies(conn: MongoClient, offset: int, size: int):
    row = list(conn[database_name][well_known_companies_collection].find().skip(offset).limit(size))
    return {
        "results": row,
        "total": count_all_record(conn, database_name, well_known_companies_collection)
    }

def get_one_company(conn: MongoClient, domain: str):
    row = list(conn[database_name][well_known_companies_collection].find({"_id": domain}))
    return row

def create_well_known_company(conn: MongoClient, domain: str, name: str):

    check_domain_duplicate(conn, domain, well_known_companies_collection, disposable_collection, free_collection)

    data = {
        "_id": domain,
        "pattern": create_domain_regex(domain),
        "name": name
    }
    conn[database_name][well_known_companies_collection].insert_one(data)
    return data

def update_one_company(conn: MongoClient, domain: str, company: dict):
    checker = get_one_company(conn, domain)
    if len(checker) == 0:
        raise HTTPException(status_code=404, detail="The company's domain not found")
    else:
        value = company.dict()
        name = value["name"] or list(checker)[0]["name"]
        pattern = value["pattern"] or list(checker)[0]["pattern"]

        update = {"name": name, "pattern": pattern}
        conn[database_name][well_known_companies_collection].update_one({"_id": domain}, {"$set": update})
        update["_id"] = domain
        return update

def delete_one_company(conn: MongoClient, domain: str):
    checker = get_one_company(conn, domain)
    if len(checker) == 0:
        raise HTTPException(status_code=404, detail="The company's domain not found")
    else:
        conn[database_name][well_known_companies_collection].delete_one({"_id": domain})
        return {"domain": domain}

def replace_company(conn: MongoClient, data_pandas_object):
    payload = json.loads(data_pandas_object.to_json(orient='records'))
    for data in payload:
        domain = list(data.values())[0]
        key = list(data.keys())
        spliter = domain.split(".")

        data["_id"] = data.pop(key[0])
        data["pattern"] = create_domain_regex(domain)
        data["name"] = spliter[-2]

    conn[database_name][well_known_companies_collection].drop()
    conn[database_name][well_known_companies_collection].insert(payload)
    return {"data": conn[database_name][well_known_companies_collection].count_documents({})}
