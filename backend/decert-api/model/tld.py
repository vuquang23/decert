from pydantic import Field
from typing import List, Optional
from .common import IDModelMixin
from .scammodel import ScamModel

class TLDBase(IDModelMixin, ScamModel):
    name: str
    tld: str
    value: int

class TLDInDB(TLDBase):
    pass

class TLDInCreate(ScamModel):
    name: str
    tld: str
    value: int

class TLDInUpdate(ScamModel):
    name: Optional[str]=None
    tld: Optional[str]=None
    value: Optional[int]=None

class TLDInResponse(ScamModel):
    tlds: TLDInDB

class ManyTLDInResponse(ScamModel):
    tlds: List[TLDInDB]
    document_counts: int
