from typing import List, Optional
from pydantic import BaseModel
from .common import IDModelMixin, DateTimeModelMixin
from .scammodel import ScamModel


class LabelEmail(ScamModel):
    mail_id: Optional[str]
    email: str
    content: str
    raw_content: Optional[str] 
    trust_rules: dict
    model_prediction: dict
    human_label: str = None
    Display: str = "false"
    mailbox: str
    date: str


class LabelEmailInCreate(LabelEmail):
    pass


class LabelEmailInDB(LabelEmail, DateTimeModelMixin, IDModelMixin):
    pass


class LabelEmailFilter(ScamModel):
    score: Optional[int] = 1
    human_label: Optional[str] = "none"


class LabelEmailInUpdate(BaseModel):
    human_label: Optional[str]
    Display: Optional[str]


class LabelEmailInResponse(ScamModel):
    label_email: LabelEmailInDB


class ListEmailInResponse(ScamModel):
    label_emails: List
    query_counts: int
    none_email_counts: int
    document_counts: int

class LabelEmailConverterInResponse(ScamModel):
    raw_content: str

class LabelEmailConverterDTO(BaseModel):
    url: str