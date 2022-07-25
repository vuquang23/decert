from cgitb import lookup
from pydantic import BaseModel, EmailStr, Field
from typing import List, Optional
from bson.objectid import ObjectId
from .common import DateTimeModelMixin, IDModelMixin, PyObjectId
from .scammodel import ScamModel

class EmailDisplayBase(ScamModel):
    email: str
    email_score: str
    look_ups: int
    esp: dict
    phishing: dict
    scam_adviser: dict
    trust_rules: dict
    voted_as_spam: int
    display: str

class EmailDisplayInDB(DateTimeModelMixin, EmailDisplayBase):
    pass

class EmailDisplayInCreate(ScamModel):
    email: EmailStr

class EmailDisplayInUpdate(ScamModel):
    display: Optional[str]

class EmailDisplayInResponse(ScamModel):
    email_display: EmailDisplayInDB

class ManyEmailDisplayInResponse(ScamModel):
    email_display: List[EmailDisplayInDB]
    document_counts: int
