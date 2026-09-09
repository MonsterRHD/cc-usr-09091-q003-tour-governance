import unittest
from datetime import UTC, datetime
from app.contracts import FRAMEWORK_KIND, STORAGE_KIND, EvidenceEvent

class ContractTests(unittest.TestCase):
    def test_event_has_stable_version(self):
        event = EvidenceEvent("sample", 1, datetime.now(UTC), "fixture")
        self.assertEqual(event.version, 1)
        self.assertEqual(FRAMEWORK_KIND, "FastAPI")
        self.assertEqual(STORAGE_KIND, "SQLite")
