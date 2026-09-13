def greet(name: str) -> str:
    """
    Возвращает приветствие для указанного имени.

    Args:
        name: Имя пользователя

    Returns:
        Приветственная строка
    """
    return f"Hello, {name}!"


def main() -> None:
    """Основная функция программы."""
    user_name = input("Enter your name: ")
    message = greet(user_name)
    print(message)


if __name__ == "__main__":
    main()