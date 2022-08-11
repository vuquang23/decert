from typing import List, Optional
from pydantic import BaseModel
from .common import IDModelMixin, DateTimeModelMixin
from .scammodel import ScamModel

class UnspamEmail(ScamModel):
    mail_id: Optional[str]
    email: str
    content: str
    raw_content: Optional[str] 
    trust_rules: dict
    model_prediction: dict
    human_label: str = None
    Display: str = "false"
    mailbox: str = "unspam"
    attachment: Optional[str] = ""
    date: str


class UnspamEmailInCreate(ScamModel):
    mail_id: str
    email: str
    content: str
    raw_content: str
    trust_rules: dict
    model_prediction: dict
    mailbox: str
    date: str
    attachment: Optional[str] = ""
    human_label: Optional[str] = "none"
    Display: Optional[str] = "false"


class UnspamEmailInDB(UnspamEmail, DateTimeModelMixin, IDModelMixin):
    pass


class UnspamEmailInUpdate(BaseModel):
    human_label: Optional[str] 
    Display: Optional[str] 
    attachment: Optional[str] = ""



class UnspamEmailInResponse(ScamModel):
    unspam_email: UnspamEmailInDB


class ListUnspamEmailInResponse(ScamModel):
    unspam_emails: List[UnspamEmailInDB]
    query_counts: int
    document_counts: int




# class UnspamEmailInUpdate(BaseModel):
#     display: str

# class UnspamEmailInResponse(BaseModel):
#     unspam_email: dict

# class ListUnspamEmailInResponse(BaseModel):
#     unspam_emails: list
#     query_counts: int
#     document_counts: int