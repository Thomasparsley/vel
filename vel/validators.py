def str_max_length(name: str, value: str, max: int) -> str:
    if len(value) > max:
        raise ValueError(f"Max length of {name} is {max} chars")
    return value


def str_min_length(name: str, value: str, min: int) -> str:
    if len(value) < min:
        raise ValueError(f"Min length of {name} is {min} chars")
    return value
