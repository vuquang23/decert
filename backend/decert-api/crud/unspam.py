import datetime
from typing import List, Optional
from pymongo import MongoClient
from fastapi import HTTPException
from bson.objectid import ObjectId
from fastapi.encoders import jsonable_encoder
import requests
from ..core.utils import count_all_record
from .utils import space_client

from ..core.config import (
    email_database_name,
    label_email_collection,
)

from ..model.unspam import (
    UnspamEmailInDB,
    UnspamEmailInUpdate,
)

def check_mail_exist(conn: MongoClient, mail_id: str) -> bool:
    row = conn[email_database_name][label_email_collection].find_one({"mail_id": mail_id})
    if row:
        return True
    else:
        return False


def count_all_records_unspam(conn: MongoClient) -> int:
    return conn[email_database_name][label_email_collection].count(
        {"mailbox": "unspam"}
    )


def get_raw_content(url: str) -> str:
    result = requests.get(url)
    return result.text


def get_one_unspam_email(conn: MongoClient, email_id: str) -> UnspamEmailInDB:
    row = conn[email_database_name][label_email_collection].find_one({"_id": email_id})
    if row:
        row["raw_content"] = get_raw_content(row["raw_content"])
        return UnspamEmailInDB(**row)

    raise HTTPException(status_code=404, detail="Email not found")


def get_one_unspam_email_by_mail_id(conn: MongoClient, email_id: str) -> UnspamEmailInDB:
    row = conn[email_database_name][label_email_collection].find_one({"mail_id": email_id})
    if row:
        row["raw_content"] = get_raw_content(row["raw_content"])
        return UnspamEmailInDB(**row)

    raise HTTPException(status_code=404, detail="Email not found")

def get_all_unspam_email(
    conn: MongoClient, offset: int, size: int
) -> List[UnspamEmailInDB]:
    list_unspams: List[UnspamEmailInDB] = []
    rows = (
        conn[email_database_name][label_email_collection]
        .find({"mailbox": "unspam"})
        .skip(offset)
        .limit(size)
    )
    for row in rows:
        list_unspams.append(UnspamEmailInDB(**row))

    return list_unspams, len(list_unspams), count_all_records_unspam(conn)


def update_unspam_email(
    conn: MongoClient, email_id: str, payload: UnspamEmailInUpdate
) -> dict:
    value = payload.dict()
    unspam_email_in_db = get_one_unspam_email(conn, email_id=email_id)

    value["Display"] = value["Display"] or unspam_email_in_db.Display
    value["human_label"] = value["human_label"] or unspam_email_in_db.Display
    value["attachment"] = value["attachment"] or unspam_email_in_db.attachment

    conn[email_database_name][label_email_collection].update_one(
        {"_id": email_id},
        {"$set": {"Display": value["Display"], "human_label": value["human_label"], "attachment": value["attachment"]}},
    )

    return value


def create_unspam_email(
    conn: MongoClient,
    mail_id: str,
    email: str,
    content: str,
    raw_content: str,
    trust_rules: dict,
    model_prediction: dict,
    date: str,
    mailbox: str,
    attachment: Optional[str] = "",
    human_label: Optional[str] = "none",
    Display: Optional[str] = "false",
) -> UnspamEmailInDB:
    space_content = space_client.upload_to_space(raw_content, mail_id)
    unspam_email = UnspamEmailInDB(
        mail_id=mail_id,
        email=email,
        content=content,
        raw_content=space_content,
        trust_rules=trust_rules,
        model_prediction=model_prediction,
        human_label=human_label,
        Display=Display,
        mailbox=mailbox,
        attachment=attachment,
        date=date,
        created_at=UnspamEmailInDB.default_datetime(),
    )

    unspam_email = jsonable_encoder(unspam_email)

    new_label_email = conn[email_database_name][label_email_collection].insert_one(
        unspam_email
    )

    return get_one_unspam_email(conn, email_id=new_label_email.inserted_id)


def delete_unspam_email(conn: MongoClient, _id: str, mail_id: str) -> None:
    conn[email_database_name][label_email_collection].delete_one({"_id": _id})
    space_client.delete_from_space(mail_id)
