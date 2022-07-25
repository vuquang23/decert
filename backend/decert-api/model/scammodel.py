import datetime
from bson import ObjectId
from pydantic import BaseConfig, BaseModel
import pytz

def convert_datetime_to_realworld(dt: datetime.datetime) -> str:
    return dt.replace(tzinfo=pytz.timezone('Asia/Ho_Chi_Minh')).isoformat()

class ScamModel(BaseModel):
    class Config(BaseConfig):
        arbitrary_types_allowed = True
        allow_population_by_field_name = True
        json_encoders = {
            datetime.datetime: convert_datetime_to_realworld,
            ObjectId: str
        }
