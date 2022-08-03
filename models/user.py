from datetime import datetime
from typing import Final, NewType

from tortoise import fields, models

from . import basic_fields


Permission: Final = NewType("Permission", str)
Role: Final = NewType("Role", str)


class Permissions:
    __storage: dict[Permission, bool]

    def add(self, permission: Permission):
        self.__storage[permission] = True

    def __iadd__(self, permission: Permission):
        self.add(permission)
        return self

    def has(self, permission: Permission) -> bool:
        return self.__storage.get(permission, False)

    def __getitem__(self, permission: Permission) -> bool:
        return self.has(permission)

    def remove(self, permission: Permission):
        if self.has(permission):
            self.__storage.pop(permission)

    def __isub__(self, permission: Permission):
        self.remove(permission)
        return self


class AuthorizationRules:
    roles: dict[Role, Permissions] = {}

    def has_role(self, role: Role) -> bool:
        return role in self.roles

    def get_permissions(self, role: Role) -> Permissions | None:
        return self.roles.get(role, None)


class Authorization(models.Model):
    RULES: Final = AuthorizationRules()

    admin = fields.BooleanField(default=False)
    role = fields.CharField(max_length=12)

    class Meta:  # type: ignore
        abstract = True

    def is_admin(self) -> bool:
        return self.admin

    def has_role(self, role: Role) -> bool:
        return self.role == role

    def has_permission(self, permission: Permission) -> bool:
        permissions = self.RULES.get_permissions(Role(self.role))

        if permissions:
            return permissions[permission]

        return False


class User(Authorization):
    id = basic_fields.ID_FIELD
    enabled = fields.BooleanField(default=True)
    username = fields.CharField(max_length=64, null=False, unique=True, index=True)
    email = fields.CharField(max_length=320, null=False, unique=True, index=True)
    password = fields.CharField(max_length=128, null=False, index=True)

    created_at = basic_fields.CREATED_AT_FIELD
    updated_at = basic_fields.UPDATED_AT_FIELD

    class Meta: # type: ignore
        table = "vel_users"

    def __init__(self, username: str, email: str, password: str):
        self.username = username
        self.email = email
        self.password = password

        self.enabled = True
        self.created_at = datetime.now()

        super().__init__()

    def __eq__(self, other: object) -> bool:
        if not isinstance(other, User):
            return False

        return self.id == other.id


CREATED_BY_FIELD: fields.ForeignKeyRelation[User] = fields.ForeignKeyField(
    "vel.User", related_name=False
)
UPDATED_BY_FIELD: fields.ForeignKeyRelation[User] = fields.ForeignKeyField(
    "vel.User", related_name=False, null=True
)
