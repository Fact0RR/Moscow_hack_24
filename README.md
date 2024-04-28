<div align="center">
  
# ЦИФРОВОЙ ПРОРЫВ: СЕЗОН ИИ <br> Компетентный подбор образовательных курсов

<img height="90" alt="logo" src="assets/logo1.png">

</div> 


## Оглавление
- ### [Задание](#1)
- ### [Решение](#2)
- ### [Запуск кода](#3)
- ### [Уникальность нашего решения](#4)
- ### [Стек](#5)
- ### [Команда](#6)
- ### [Ссылки](#7)

## <a name="1"> Задание </a>

Сегодня пользователям сложно ориентироваться в большом числе обучающих курсов на образовательных платформах, а также прорабатывать свою личную траекторию развития для получения необходимого уровня знаний, соответствующего желаемой профессии.

Необходимо разработать MVP на основе искусственного интеллекта, который сможет определить, какую программу обучения стоит изучить, чтобы приобрести необходимые навыки и знания для успешного трудоустройства. Решение должно анализировать требования, предъявляемые кандидатам в вакансиях, и предлагать персонализированные рекомендации образовательных курсов, продуктов или других обучающих материалов, соответствующих потребностям рынка труда.

## <a name="2">Решение </a>

### Архетиктура решения

<div align="center">
<img height="300" alt="logo" src="assets/arh.jpg">

## <a name="3">Запуск кода </a>

### Последовательные шаги для запуска кода:
1. Склонируйте гит репозиторий;
```Bash
git clone https://github.com/BuldakovN/RRR-Video-Action-Classification.git
```
2. Скачайте веса для модели детекции https://drive.google.com/drive/folders/1iQRcneyDnb3q7LZ9wcGfkDt3W5zAKjwD?usp=sharing и разместите их в папках ```model/weights``` и ```modelStream/weights```

3. Запуск контейнеров и сервера стриминга:
```Bash
cd RRR-Video-Action-Classification
docker-compose build
docker-compose up

cd modelStream

pip3 install -r requirements.txt
python3 serverStream.py
```

## <a name="4">Уникальность нашего решения </a>

1. Возможность распознавания действий нескольких людей на одном кадре.

2. Работает в режиме реального времени.

3. Модель была предобучена на большом датасете.

4. Удобный веб-интерфейс для работы с моделью.

## <a name="5">Стек </a>
<div align="center">
  <img src="https://github.com/devicons/devicon/blob/master/icons/python/python-original-wordmark.svg" title="Python" alt="Python" height="40"/>&nbsp;
  <img src="https://github.com/devicons/devicon/blob/master/icons/postgresql/postgresql-original-wordmark.svg" title="postgresql" alt="postgresql" height="40"/>&nbsp;
  <img src="https://github.com/gin-gonic/logo/blob/master/color.svg" title="gin-gonic" alt="gin-gonic" height="40"/>&nbsp;
  
  <img src="https://github.com/devicons/devicon/blob/master/icons/docker/docker-original-wordmark.svg" title="docker" alt="docker" height="40"/>&nbsp;
  <img src="https://github.com/devicons/devicon/blob/master/icons/javascript/javascript-original.svg" title="JS" alt="JS"  height="40"/>&nbsp;
  <img src="https://fronty.com/static/uploads/1.11-30.11/languages%20in%202022/go.png" title="GO" alt="GO" height="40"/>&nbsp;
  <img src="https://upload.wikimedia.org/wikipedia/commons/3/3c/Flask_logo.svg"  title="Flask" alt="Flask" height="40"/>
</div>

## <a name="6">Команда </a>

*Состав команды "Герои ML и Тильта"*   

- <h4><img align="center" height="25" src="https://user-images.githubusercontent.com/51875349/198863127-837491f2-b57f-4c75-9840-6a4b01236c7a.png">: @Ubludor, Маслов Денис - Fullstack-developer</h3>
- <h4><img align="center" height="25" src="https://user-images.githubusercontent.com/51875349/198863127-837491f2-b57f-4c75-9840-6a4b01236c7a.png">: @Skadar7, Кузнецов Денис - ML-engineer</h3>
- <h4><img align="center" height="25" src="https://user-images.githubusercontent.com/51875349/198863127-837491f2-b57f-4c75-9840-6a4b01236c7a.png">: @Llaceyne, Гулария Лана - Designer, Frontend-developer</h3>
- <h4><img align="center" height="25" src="https://user-images.githubusercontent.com/51875349/198863127-837491f2-b57f-4c75-9840-6a4b01236c7a.png">: @vseesheoleg, Сивец Олег - ML-engineer</h3>

## <a name="7">Ссылки </a>

- [ссылка на веса модели детекции](https://drive.google.com/drive/folders/1iQRcneyDnb3q7LZ9wcGfkDt3W5zAKjwD?usp=sharing)&nbsp;
- [ссылка на скринкаст](https://drive.google.com/file/d/1Md4uNQFHtO_w9xxaigyc6ftbF0Tru6qk/view?usp=sharing)&nbsp;
