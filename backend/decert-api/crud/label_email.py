from typing import List, Optional
from unittest import result
from pymongo import MongoClient
from fastapi import HTTPException
import requests
import re
from ..core.utils import count_all_record
from fastapi.encoders import jsonable_encoder
from .utils import space_client

from ..core.config import (
    email_database_name,
    label_email_collection,
)

from ..model.label_email import (
    LabelEmailInDB,
    LabelEmailInUpdate,
    ListEmailInResponse,
)


def search_for_label_email(conn: MongoClient, word: str) -> List[LabelEmailInDB]:
    label_email = list(
        conn[email_database_name][label_email_collection]
        # .find({"$or": [{"$text": {"$search": word}}, {"content": {"$regex": word}}]})
        .find({"$text": {"$search": word}})
    )

    return ListEmailInResponse(
            label_emails=label_email,
            query_counts=len(label_email),
            document_counts=count_all_record(
                conn, email_database_name, label_email_collection
            ),
        )

def get_one_label_email(conn: MongoClient, email_id: str) -> LabelEmailInDB:
    row = conn[email_database_name][label_email_collection].find_one({"_id": email_id})
    if row:
        return LabelEmailInDB(**row)

    raise HTTPException(status_code=404, detail="Email not found")

def get_emails_with_none_human_label(conn: MongoClient) -> int:
    row = conn[email_database_name][label_email_collection].count_documents({"human_label": "none"})
    return row



def get_label_email(
    conn: MongoClient,
    offset: int,
    size: int,
    score: Optional[int] = None,
    human_label: Optional[str] = None,
    Display: Optional[str] = None,
    start_date: Optional[str] = None,
    end_date: Optional[str] = None,
) -> ListEmailInResponse:
    if start_date or end_date:
        initial_args = {
            "model_prediction.score": score,
            "human_label": human_label,
            "Display": Display,
            "date": {"$gte": start_date} or {"$lte": end_date},
        }
    elif start_date and end_date:
        initial_args = {
            "model_prediction.score": score,
            "human_label": human_label,
            "Display": Display,
            "date": {"$gte": start_date, "$lte": end_date},
        }
    else:
        initial_args = {
            "model_prediction.score": score,
            "human_label": human_label,
            "Display": Display,
            "date": None,
        }

    final_args = {k: v for k, v in initial_args.items() if v is not None}

    if final_args:
        label_email = list(
            conn[email_database_name][label_email_collection]
            .find(final_args)
            .skip(offset)
            .limit(size)
            .sort("created_at", -1)
        )
    else:
        label_email = list(
            conn[email_database_name][label_email_collection]
            .find()
            .skip(offset)
            .limit(size)
            .sort("created_at", -1)
        )

    return ListEmailInResponse(
        label_emails=label_email,
        query_counts=len(label_email),
        none_email_counts=get_emails_with_none_human_label(conn),
        document_counts=count_all_record(
            conn, email_database_name, label_email_collection
        ),
    )


def create_label_email(
    conn: MongoClient,
    mail_id: str,
    email: str,
    content: str,
    raw_content: str,
    trust_rules: dict,
    model_prediction: dict,
    mailbox: str,
    date: str,
    human_label: Optional[str] = "none",
    Display: Optional[str] = "false",
) -> LabelEmailInDB:
    space = space_client.upload_to_space(raw_content, mail_id)
    label_email = LabelEmailInDB(
        mail_id=mail_id,
        email=email,
        content=content,
        raw_content=space,
        trust_rules=trust_rules,
        model_prediction=model_prediction,
        human_label=human_label,
        Display=Display,
        mailbox=mailbox,
        date=date,
        created_at=LabelEmailInDB.default_datetime(),
    )

    label_email = jsonable_encoder(label_email)

    new_label_email = conn[email_database_name][label_email_collection].insert_one(
        label_email
    )

    return get_one_label_email(conn, email_id=new_label_email.inserted_id)


def update_label_email(conn: MongoClient, email_id: str, item: LabelEmailInUpdate):
    value = item.dict()
    db_label_email = get_one_label_email(conn, email_id)
    if not db_label_email:
        raise HTTPException(status_code=409, detail="Email Label Not Found")
    else:
        value["human_label"] = value["human_label"] or db_label_email.human_label
        value["Display"] = value["Display"] or db_label_email.Display

        conn[email_database_name][label_email_collection].update_one(
            {"_id": email_id}, 
            {"$set": {"human_label": value["human_label"], "Display": value["Display"]}},
        )

        return value

def delete_label_email(conn: MongoClient, mail_id: str, _id: str):
    conn[email_database_name][label_email_collection].delete_one({"_id": _id})
    space_client.delete_from_space(mail_id)


def get_raw_content(url: str) -> str:
    result = requests.get(url)
    return result.text

def get_mailbox_group_by_day(conn: MongoClient, date: str):
    date_tring = f"{date}"
    regx = re.compile(f"{date}", re.IGNORECASE)
    print(date)
    pipeline = [
        {
            '$addFields': {
                'matchObject': {
                    '$regexFind': {
                        'input': '$created_at', 
                        'regex': regx
                    }
                }
            }
        }, {
            '$match': {
                'matchObject': {
                    '$ne': None
                }
            }
        }, {
            '$group': {
                '_id': '$mailbox', 
                'count': {
                    '$sum': 1
                }
            }
        }
    ]

    result = conn[email_database_name][label_email_collection].aggregate(pipeline)
    return result