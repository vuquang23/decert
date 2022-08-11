import json
from pymongo import MongoClient 
from fastapi import HTTPException
import boto3

from ..core.config import (
    database_name, 
    ACCESS_ID,
    SPACE_ACCESS_KEY,
    ENDPOINT_URL,
    REGION_NAME
)

def check_domain_duplicate(conn: MongoClient, domain: str, collection1: str, collection2: str, collection3: str):
    row1 = list(conn[database_name][collection1].find({"_id": domain}))
    row2 = list(conn[database_name][collection2].find({"_id": domain}))
    row3 = list(conn[database_name][collection3].find({"_id": domain}))

    if len(row1) != 0:
        raise HTTPException(status_code=409, detail=f"The domain is already exist in {collection1} list")
    if len(row2) != 0:
        raise HTTPException(status_code=409, detail=f"The domain is already exist in {collection2} list")
    if len(row3) != 0:
        raise HTTPException(status_code=409, detail=f"The domain is already exist in {collection3} list")

class Space:
    def __init__(self, access_id: str, access_key: str, endpoint_url: str, region_name: str):
        self.client = boto3.session.Session().client(
            's3',
            region_name=region_name,
            endpoint_url=endpoint_url,
            aws_access_key_id=access_id,
            aws_secret_access_key=access_key
        )

    def upload_to_space(self, raw_content: str, _id: str):
        # Upload a file to your Space
        self.client.put_object(Bucket='scamadviser-static', # The path to the directory you want to upload the object to, starting with your Space name.
                        Key=f'{_id}.txt', # Object key, referenced whenever you want to access this file later.
                        Body=raw_content, # The object's contents.
                        ACL='public-read', # Defines Access-control List (ACL) permissions, such as private or public.
                        ContentType="text/plain"
                    )
        
        return f"https://scamadviser-static.{REGION_NAME}.digitaloceanspaces.com/{_id}.txt"
    
    def update_attachment_to_space(self, attachment_content: bytes, filename: str, _id: str):
        self.client.put_object(Bucket='scamadviser-static', # The path to the directory you want to upload the object to, starting with your Space name.
                        Key=f'{_id}.{filename}', # Object key, referenced whenever you want to access this file later.
                        Body=attachment_content, # The object's contents.
                        ACL='public-read', # Defines Access-control List (ACL) permissions, such as private or public.
                        ContentType="image/jpeg"
                    )
        
        return f"https://scamadviser-static.{REGION_NAME}.digitaloceanspaces.com/{_id}.{filename}"
    
    def delete_from_space(self, _id: str):
        self.client.delete_object(
            Bucket='scamadviser-static',
            Key=f'{_id}.txt',
        )

        return f'{_id}.txt'

space_client = Space(ACCESS_ID, SPACE_ACCESS_KEY, ENDPOINT_URL, REGION_NAME)
