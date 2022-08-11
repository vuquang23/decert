import datetime
from typing import Dict, List, Optional
from fastapi import HTTPException, Depends
from pydantic import BaseModel
from pymongo import MongoClient
from starlette.responses import JSONResponse
from fastapi.encoders import jsonable_encoder
import subprocess


def create_domain_regex(domain: str):
    split = domain.split(".")
    if len(split) == 1:
        raise HTTPException(status_code=404, detail="domain invalid")
    else:
        split[-2] = f"{split[-2]}\\"
        split[0] = f"(.*\.)?{split[0]}"
        result = ".".join(split)
        return result


def count_all_record(db: MongoClient, database_name: str, collection_name: str):
    return db[database_name][collection_name].count()


def create_team_email_regex(_id: str):
    _id = f".*{_id}.*"
    return _id


def convert_datetime_for_json_object(data: Dict, fields: List[str]) -> Dict:
    for field in fields:
        try:
            string_value = data[field]
            data[field] = datetime.fromisoformat(string_value)
        except Exception:
            pass
    return data

def create_aliased_response(model: BaseModel) -> JSONResponse:
    return JSONResponse(content=jsonable_encoder(model, by_alias=True))


def check_ebl(email: str):
    """
    Check if an email is in EBL
    Return true if exist in ebl db, else false
    """
    args = ["bash", "cms-api/dnslcheck.sh", email]
    ebl_script_output = subprocess.run(args, capture_output=True)
    stdout = ebl_script_output.stdout.decode("utf-8")
    stderr = ebl_script_output.stderr.decode("utf-8")

    return {"result": stdout != "", "output": stdout, "error": stderr}
