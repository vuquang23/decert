import email
from typing import List, Optional
from pymongo import MongoClient
from fastapi import HTTPException
from fastapi.encoders import jsonable_encoder
from bson.objectid import ObjectId
import requests

from ..core.config import TRUST_RULE_API_BASE_PATH, email_database_name, email_display_collection
from ..model.email_display import EmailDisplayInDB
from ..core.utils import count_all_record


def get_all_email_display(
    conn: MongoClient, offset: str, size: str
) -> List[EmailDisplayInDB]:
    list_email_display: List[EmailDisplayInDB] = []
    email_display = list(
        conn[email_database_name][email_display_collection]
        .find()
        .skip(offset)
        .limit(size)
    )
    for email in email_display:
        list_email_display.append(EmailDisplayInDB(**email))

    return list_email_display


def get_email_display(conn: MongoClient, email_display: str) -> EmailDisplayInDB:
    email_display = conn[email_database_name][email_display_collection].find_one(
        {"email": email_display}
    )
    if email_display:
        return EmailDisplayInDB(**email_display)

    raise HTTPException(status_code=404, detail="Email Display not found")


def create_email_display(
    email: str 
) -> EmailDisplayInDB:
    trust_rule_result = requests.get(f"{TRUST_RULE_API_BASE_PATH}/api/v1/email/check-scam", params={"email": email}).json()
    return trust_rule_result


def update_email_display(
    conn: MongoClient, email_display: str, display: str
) -> EmailDisplayInDB:
    display_email_db = get_email_display(conn, email_display)

    display_email_db.display = display or display_email_db.display

    conn[email_database_name][email_display_collection].update_one(
        {"email": email_display}, {"$set": {"display": display_email_db.display}}
    )
    return display_email_db


def delete_email_display(conn: MongoClient, email_display: str) -> None:
    conn[email_database_name][email_display_collection].delete_one(
        {"email": email_display}
    )