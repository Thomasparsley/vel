from typing import Final, NewType, TypeVar, Any
from datetime import datetime, timedelta

from tortoise import fields, models

from fastapi import Cookie
from passlib.context import CryptContext
from jose import JWTError, jwt

from vel.exceptions import InvalidAuthenticationError

from . import basic_fields
from .hashids import HashidsMixin, HashidsSingleton
from .config_factory import ConfigFactory


JWT_ALGORITHM: Final = "HS256"
SECRET_KEY: str = ConfigFactory().get().SECRET_KEY

UserType = TypeVar("UserType", bound="UserModel")
RoleType: Final = NewType("RoleType", str)
PermissionType: Final = NewType("PermissionType", str)


pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")


def get_password_hash(password: str):
    return pwd_context.hash(password)


def create_access_token(
    user: "UserModel", data: dict[str, Any], expires_delta: timedelta | None = None
):
    data["id"] = user.hashed_id
    payload = data.copy()

    if expires_delta:
        expire = datetime.utcnow() + expires_delta
    else:
        expire = datetime.utcnow() + timedelta(minutes=15)

    payload.update({"exp": expire})

    encoded_jtw = jwt.encode(payload, SECRET_KEY, algorithm=JWT_ALGORITHM)
    return encoded_jtw


async def try_get_current_user(token: str | None = Cookie(default=None)):
    if not token:
        return None

    try:
        payload = jwt.decode(token, SECRET_KEY, algorithms=[JWT_ALGORITHM])
    except JWTError:
        return None

    hashed_id = payload.get("id")
    if not hashed_id:
        return None

    user = await UserModel.get_with_hashed_id(hashed_id)
    return user


async def get_current_user(token: str | None = Cookie(default=None)):
    user = try_get_current_user(token)
    if not user:
        raise InvalidAuthenticationError
    return user


class UserModel(models.Model, HashidsMixin):
    id = basic_fields.ID_FIELD
    enabled = fields.BooleanField(default=True)
    username = fields.CharField(max_length=64, null=False, unique=True, index=True)
    email = fields.CharField(max_length=320, null=False, unique=True, index=True)
    password = fields.CharField(max_length=128, null=False, index=True)

    created_at = basic_fields.CREATED_AT_FIELD
    updated_at = basic_fields.UPDATED_AT_FIELD

    class Meta:  # type: ignore
        abstract = True

    def __init__(self, username: str, email: str, password: str):
        super().__init__()

        self.username = username
        self.email = email
        self.password = get_password_hash(password)

        self.enabled = True
        self.created_at = datetime.now()

    def __eq__(self, other: object) -> bool:
        if not isinstance(other, UserModel):
            return False

        return self.id == other.id

    @classmethod
    async def get_with_hashed_id(cls, hashed_id: str):
        id = HashidsSingleton().decode_single(hashed_id)
        if not id:
            return None

        return await cls.get_or_none(id=id)

    def verify_password(self, to_verify: str):
        return pwd_context.verify(to_verify, self.password)


class Permissions:
    __storage: dict[PermissionType, bool]

    def add(self, permission: PermissionType):
        self.__storage[permission] = True

    def __iadd__(self, permission: PermissionType):
        self.add(permission)
        return self

    def has(self, permission: PermissionType) -> bool:
        return self.__storage.get(permission, False)

    def __getitem__(self, permission: PermissionType) -> bool:
        return self.has(permission)

    def remove(self, permission: PermissionType):
        if self.has(permission):
            self.__storage.pop(permission)

    def __isub__(self, permission: PermissionType):
        self.remove(permission)
        return self


class AuthorizationRules:
    roles: dict[RoleType, Permissions] = {}

    def has_role(self, role: RoleType) -> bool:
        return role in self.roles

    def get_permissions(self, role: RoleType) -> Permissions | None:
        return self.roles.get(role, None)


class Authorization(models.Model):
    RULES: Final = AuthorizationRules()

    admin = fields.BooleanField(default=False)
    role = fields.CharField(max_length=12)

    class Meta:  # type: ignore
        abstract = True

    def is_admin(self) -> bool:
        return self.admin

    def has_role(self, role: RoleType) -> bool:
        return self.role == role

    def has_permission(self, permission: PermissionType) -> bool:
        permissions = self.RULES.get_permissions(RoleType(self.role))

        if permissions:
            return permissions[permission]

        return False


def authenticate_user(user: UserModel | None, password: str | None) -> bool:
    if not user:
        return False
    elif not password:
        return False

    return user.verify_password(password)
