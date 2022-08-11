import datetime
from pydantic import BaseModel, Field, validator
from bson import ObjectId

class PyObjectId(ObjectId):
    @classmethod
    def __get_validators__(cls):
        yield cls.validate

    @classmethod
    def validate(cls, v):
        if not ObjectId.is_valid(v):
            raise ValueError("Invalid objectid")
        return ObjectId(v)

    @classmethod
    def __modify_schema__(cls, field_schema):
        field_schema.update(type="string")

class IDModelMixin(BaseModel):
    id: PyObjectId = Field(default_factory=PyObjectId, alias="_id")

class DateTimeModelMixin(BaseModel):
    created_at: datetime.datetime = None  

    @validator("created_at", pre=True)
    def default_datetime(
            cls,
            value: datetime.datetime = None,
    ) -> datetime.datetime:
        return value or datetime.datetime.now()