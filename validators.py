def str_max_length(name: str, value: str, max: int) -> str:
    if len(value) > max:
        raise ValueError(f"Max length of {name} is {max} chars")
    return value
