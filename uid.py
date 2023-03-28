"""Mostly a port of https://github.com/Mediagone/small-uid"""

import random
import time
from typing import Optional

MAX_20BIT_INT = 524288


class UID:
    _uid: int

    def __init__(
        self, int: Optional[int] = None, existing: Optional[set["UID"]] = None
    ):
        if existing is None:
            existing = set()
        self._initialize(int)
        while self in existing:
            self._initialize(int)

    def __repr__(self) -> str:
        return str(self._uid)

    def __str__(self) -> str:
        return str(self._uid)

    def __eq__(self, __o: "UID") -> bool:
        return self._uid == __o._uid

    def __hash__(self) -> int:
        return self._uid

    @property
    def int(self):
        return self._uid

    @property
    def hex(self):
        return hex(self._uid)

    def _initialize(self, i: Optional["int"] = None):
        if i is None:
            time_ms = int(time.time() * 1000)
            time_padded = time_ms * 1000000
            rand = random.randrange(MAX_20BIT_INT)
            self._uid = time_padded + rand
        else:
            self._uid = i
