from typing import Optional
from pydantic import BaseModel

class AnonymousInCreate(BaseModel):
    name: str
    domain: str

class AnonymousInGetAll(BaseModel):
    offset: int
    size: int

class AnonymousInGetOne(BaseModel):
    domain: str

class AnonymousInUpdate(BaseModel):
    name: str
    pattern: Optional[str]

class AnonymousInDelete(BaseModel):
    domain: str
