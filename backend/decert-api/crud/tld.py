from typing import List, Optional
from pymongo import MongoClient 
from fastapi import HTTPException
from fastapi.encoders import jsonable_encoder
from bson import ObjectId

from ..core.config import (
    database_name, 
    tld_collection
)
from ..model.tld import TLDInDB, TLDInCreate, TLDInUpdate
from ..core.utils import count_all_record

def get_all_tlds(conn: MongoClient, offset: str, size: str) -> List[TLDInDB]:
    list_tlds: List[TLDInDB] = []
    tlds = list(conn[database_name][tld_collection].find().skip(offset).limit(size))
    for tld in tlds:
        list_tlds.append(TLDInDB(**tld))
    
    return list_tlds

def get_tld(conn: MongoClient, tld_id: str) -> TLDInDB:
    tld = conn[database_name][tld_collection].find_one({"_id": tld_id})
    if tld:
        return TLDInDB(**tld)

    raise HTTPException(status_code=404, detail="TLD not found")

def create_tld(
    conn: MongoClient, 
    name: str,
    tld: str,
    value: int,
) -> TLDInDB:
    tld = TLDInDB(name=name, tld=tld, value=value)
    tld = jsonable_encoder(tld)
    new_tld = conn[database_name][tld_collection].insert_one(tld)
    return get_tld(
        conn,
        tld_id = new_tld.inserted_id
    )

def update_tld(
    conn: MongoClient,
    tld_id: ObjectId,
    name: Optional[str]=None,
    tld: Optional[str]=None,
    value: Optional[int]=None,
) -> TLDInDB:
    tld_in_db = get_tld(conn, tld_id)

    tld_in_db.name = name or tld_in_db.name
    tld_in_db.tld = tld or tld_in_db.tld
    tld_in_db.value = value or tld_in_db.value

    conn[database_name][tld_collection].update_one(
        {"_id": tld_id},
        {"$set": {
            "name": tld_in_db.name,
            "tld": tld_in_db.tld,
            "value": tld_in_db.value
        }}
    )
    return tld_in_db

def delete_tld(conn: MongoClient, tld_id: str) -> None:
    conn[database_name][tld_collection].delete_one({"_id": tld_id})

# checker
def check_tld_exist(conn: MongoClient, tld_id: ObjectId) -> bool:
    count = get_tld(conn, tld_id)
    if count:
        return True
    return False