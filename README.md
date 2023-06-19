# Шаламкова Алиса БПИ-218
# Коллоквиум, 4 вариант
Реализован API для управления студентами и оценками

## Запуск программы
Склонировать к себе репозиторий и в терминале с папкой, в которую склонирована репа, ввести команду
``
docker-compose up --build
``

## Работа с программой
Коллекция Postman с примерами запросов экспортирована в kollok.postman_collection.json.

По умолчанию в таблицах нет записей, поэтому сначала стоит отправлять пост запросы для наполнения базы данными. Формат запросов описан в файлах grade.go и student.go
