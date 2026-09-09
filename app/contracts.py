from dataclasses import dataclass
from datetime import datetime

FRAMEWORK_KIND = "FastAPI"
STORAGE_KIND = "SQLite"

@dataclass(frozen=True)
class EvidenceEvent:
    event_id: str
    version: int
    occurred_at: datetime
    source: str
