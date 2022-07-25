from typing import List, Optional
from pymongo import MongoClient
from fastapi import HTTPException

from ..core.config import (
    email_database_name,
    label_email_collection,
)

from ..model.statistics import (
    EmailGroupByModel
)

def email_group_by_prediction_label(conn: MongoClient) -> EmailGroupByModel:
    pipeline = [
        {
            '$group': {
                '_id': '$model_prediction.score', 
                'total': {
                    '$sum': 1
                }
            }
        }
    ]
    row = list(conn[email_database_name][label_email_collection].aggregate(pipeline))
    if row:
        data = {}
        for item in row:
            if item.get("_id") == 0:
                data["0"] = item.get("total")
            elif item.get("_id") == 1:
                data["1"] = item.get("total")
        return EmailGroupByModel(legit = data.get("1"), phishing = data.get("0"))


def email_group_by_human_label(conn: MongoClient):
    pipeline = [
        {
            '$group': {
                '_id': '$human_label', 
                'total': {
                    '$sum': 1
                }
            }
        }
    ]
    row = list(conn[email_database_name][label_email_collection].aggregate(pipeline))
    if row:
        return row
