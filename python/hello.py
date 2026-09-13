import sys
import time
from datetime import datetime


def main():
    start_time = time.perf_counter()

    # Если имя передали через командную строку — берём первый аргумент
    if len(sys.argv) > 1:
        name = sys.argv[1].strip()
    else:
        name = input("Введите ваше имя: ").strip()

    if not name:
        print("Имя не может быть пустым.")
        return

    current_time = datetime.now().strftime("%d.%m.%Y %H:%M:%S")

    # ANSI-коды цветов
    red = "\033[31m"
    green = "\033[32m"
    reset = "\033[0m"

    # Если имя начинается с «а» или «А» — используем красный цвет
    if name.lower().startswith("а"):
        color_code = red
    else:
        color_code = green

    print(f"{color_code}" f"Привет, {name}! Текущее время: {current_time}" f"{reset}")

    end_time = time.perf_counter()
    duration_sec = end_time - start_time
    duration_ms = duration_sec * 1000  # переводим в миллисекунды

    print(f"Время выполнения: {duration_ms:.3f} мс.")


if __name__ == "__main__":
    main()
