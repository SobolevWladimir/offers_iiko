#Offers Iiko 
## Сервер для синзронизация лояльности iikoCard->сайт 





## Пример  

 Запрос на http://adress:port/marketing/offers/check
 ```
{
    "address": {
        "building": "",
        "entrance": 0,
        "floor": 0,
        "room": 0,
        "street": ""
    },
    "cityId": 1,
    "mac": "B0:6E:BF:33:96:EF",
    "order": {
        "products": [
            {
                "offerValue": 0,
                "product": {
                    "added": false,
                    "alias": null,
                    "caloric": 203,
                    "composition": "рис, нори, сыр, лосось, авокадо, угорь, унаги соус, кунжут",
                    "hit": false,
                    "hot": false,
                    "id": 332,
                    "images": [
                        "423d922f-db87-4107-8454-8e8beed956a5.jpg",
                        null
                    ],
                    "name": "Дракон",
                    "new": false,
                    "notDeliverySeparately": false,
                    "number": "8",
                    "pfc": {
                        "carbohydrates": "14.3",
                        "fat": "12.8",
                        "protein": "7.7"
                    },
                    "sity_info": {
                        "id": 2567,
                        "price": 349,
                        "profit": 0
                    },
                    "type": {
                        "id": 1,
                        "name": "Роллы"
                    },
                    "vendor1": 50019,
                    "vendor2": 50019,
                    "volume": null,
                    "weight": "285"
                },
                "quantity": 1
            },
            {
                "offerValue": 0,
                "product": {
                    "added": false,
                    "alias": null,
                    "caloric": 153.4,
                    "composition": "рис, нори, сыр, креветка, огурец, лосось",
                    "hit": false,
                    "hot": false,
                    "id": 333,
                    "images": [
                        "2440c0f1-7168-4bd0-bbd9-5aa28c7f5ddd.jpg",
                        null
                    ],
                    "name": "Банзай",
                    "new": false,
                    "notDeliverySeparately": false,
                    "number": "8",
                    "pfc": {
                        "carbohydrates": "14.1",
                        "fat": "7.5",
                        "protein": "7.6"
                    },
                    "sity_info": {
                        "id": 2565,
                        "price": 299,
                        "profit": 0
                    },
                    "type": {
                        "id": 1,
                        "name": "Роллы"
                    },
                    "vendor1": 50219,
                    "vendor2": 50219,
                    "volume": null,
                    "weight": "275"
                },
                "quantity": 1
            },
            {
                "offerValue": 0,
                "product": {
                    "added": false,
                    "alias": null,
                    "caloric": 160.1,
                    "composition": "рис, нори, сыр, угорь, лосось, огурец, икра масаго",
                    "hit": false,
                    "hot": false,
                    "id": 334,
                    "images": [
                        "c4f44716-e3f0-45dd-b242-97abb4c011a3.jpg",
                        null
                    ],
                    "name": "Шеф",
                    "new": false,
                    "notDeliverySeparately": false,
                    "number": "8",
                    "pfc": {
                        "carbohydrates": "14.8",
                        "fat": "7.8",
                        "protein": "7.7"
                    },
                    "sity_info": {
                        "id": 2581,
                        "price": 309,
                        "profit": 0
                    },
                    "type": {
                        "id": 1,
                        "name": "Роллы"
                    },
                    "vendor1": 50319,
                    "vendor2": 50319,
                    "volume": null,
                    "weight": "260"
                },
                "quantity": 1
            }
        ],
        "totalPrice": 957
    },
    "orderInfo": {
        "bonusPay": 0,
        "cash": 0,
        "comment": "",
        "email": "admin@kapibaras.ru",
        "name": "Тестовый",
        "noChange": false,
        "orderTime": "on-ready",
        "orderType": "delivery",
        "paidDelivery": {},
        "payMethod": 6,
        "persons": 1,
        "phone": "+7(123)-123-1234",
        "promocode": "test1",
        "takeAwayPoint": 0,
        "time": ""
    },
    "platform": "linux",
    "token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJjbGllbnQiOjY2OTZ9.U0J6Q_XZqXoOIey_t7-5jQjLycVmq8hHRoX3zU0RnGY",
    "transaction_mail": "admin@kapibaras.ru"
} 
 ```
 Ответ: 
```
[
    {
        "type": 2,
        "data": {
            "max": 1,
            "sale_type": 1,
            "target": 332,
            "value": 349,
            "orderItemId": "11e64765-68e0-4e17-9fd3-2ea4bfe485f2"
        }
    },
    {
        "type": 6,
        "data": {
            "type": 1,
            "name": "Goga",
            "description": ", придёт и накажет тебя :)",
            "iiko_action_id": "",
            "target": 0
        }
    },
    {
        "type": 6,
        "data": {
            "type": 0,
            "name": "sssss",
            "description": ", kkkkkkkkkk :)",
            "iiko_action_id": "",
            "target": 1
        }
    },
    {
        "type": 1,
        "data": {
            "id": 450,
            "name": "Креветки и кальмар в сливочном соусе",
            "weight": "360",
            "sity_info": {
                "id": 0,
                "price": 0,
                "profit": 0
            },
            "vendor1": "122119",
            "vendor2": "122119",
            "type": {
                "id": 8,
                "name": "Вок"
            },
            "comment": "2019",
            "alias": "",
            "number": "",
            "caloric": 283.8,
            "pfs": "",
            "composition": "",
            "image": "54a55530-b45f-4191-b897-62f3b9c26aa7.jpg",
            "full_image": "",
            "images": [
                "54a55530-b45f-4191-b897-62f3b9c26aa7.jpg",
                ""
            ],
            "new": false,
            "hot": false,
            "hit": false,
            "not_delivery_separately": false,
            "volume": 0,
            "article": "",
            "added": {
                "meat": {
                    "id": 4292,
                    "product": {
                        "id": 452,
                        "name": "Креветки и кальмар",
                        "weight": "40",
                        "sity_info": {
                            "id": 0,
                            "price": 0,
                            "profit": 0
                        },
                        "vendor1": "301919",
                        "vendor2": "301919",
                        "type": {
                            "id": 7,
                            "name": "Для Вока"
                        },
                        "comment": "",
                        "alias": "",
                        "number": "",
                        "caloric": 0,
                        "pfs": "",
                        "composition": "",
                        "image": "",
                        "full_image": "",
                        "images": [
                            "",
                            ""
                        ],
                        "new": false,
                        "hot": false,
                        "hit": false,
                        "not_delivery_separately": false,
                        "volume": 0,
                        "article": ""
                    },
                    "parent": 450,
                    "quantity": 1
                },
                "vegetable": {
                    "id": 4293,
                    "product": {
                        "id": 210,
                        "name": "Овощи",
                        "weight": "80",
                        "sity_info": {
                            "id": 0,
                            "price": 0,
                            "profit": 0
                        },
                        "vendor1": "301819",
                        "vendor2": "301819",
                        "type": {
                            "id": 7,
                            "name": "Для Вока"
                        },
                        "comment": "",
                        "alias": "",
                        "number": "",
                        "caloric": 0,
                        "pfs": "",
                        "composition": "",
                        "image": "",
                        "full_image": "",
                        "images": [
                            "",
                            ""
                        ],
                        "new": false,
                        "hot": false,
                        "hit": false,
                        "not_delivery_separately": false,
                        "volume": 0,
                        "article": ""
                    },
                    "parent": 450,
                    "quantity": 1
                },
                "noodle": [
                    {
                        "id": 4301,
                        "product": {
                            "id": 198,
                            "name": "Удон",
                            "weight": "200",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "300319",
                            "vendor2": "300319",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "177595fe-31e7-4fe2-adc5-c832ba334933.jpg",
                            "full_image": "",
                            "images": [
                                "177595fe-31e7-4fe2-adc5-c832ba334933.jpg",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    },
                    {
                        "id": 4302,
                        "product": {
                            "id": 199,
                            "name": "Соба",
                            "weight": "200",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "300219",
                            "vendor2": "300219",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "badece39-0bce-497a-ac52-57f39b82d4a9.jpg",
                            "full_image": "",
                            "images": [
                                "badece39-0bce-497a-ac52-57f39b82d4a9.jpg",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    },
                    {
                        "id": 4303,
                        "product": {
                            "id": 200,
                            "name": "Рамен",
                            "weight": "200",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "300119",
                            "vendor2": "300119",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "9a299153-b6b0-4422-8a08-44fa07c5bbac.jpg",
                            "full_image": "",
                            "images": [
                                "9a299153-b6b0-4422-8a08-44fa07c5bbac.jpg",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    },
                    {
                        "id": 4304,
                        "product": {
                            "id": 201,
                            "name": "Фунчоза",
                            "weight": "200",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "300419",
                            "vendor2": "300419",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "9f35319e-4c4a-40eb-a808-4decff5d1b29.jpg",
                            "full_image": "",
                            "images": [
                                "9f35319e-4c4a-40eb-a808-4decff5d1b29.jpg",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    },
                    {
                        "id": 4305,
                        "product": {
                            "id": 571,
                            "name": "Рис",
                            "weight": "200",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "300519",
                            "vendor2": "300519",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "413936f2-2543-4eeb-b07e-20572aabbd89.jpg",
                            "full_image": "",
                            "images": [
                                "413936f2-2543-4eeb-b07e-20572aabbd89.jpg",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    }
                ],
                "toping": [
                    {
                        "id": 4294,
                        "product": {
                            "id": 202,
                            "name": "Сыр тертый",
                            "weight": "20",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "302019",
                            "vendor2": "302019",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "",
                            "full_image": "",
                            "images": [
                                "",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    },
                    {
                        "id": 4295,
                        "product": {
                            "id": 203,
                            "name": "Грибы шиитаке",
                            "weight": "20",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "302119",
                            "vendor2": "302119",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "",
                            "full_image": "",
                            "images": [
                                "",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    },
                    {
                        "id": 4296,
                        "product": {
                            "id": 204,
                            "name": "Ананас",
                            "weight": "20",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "302219",
                            "vendor2": "302219",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "",
                            "full_image": "",
                            "images": [
                                "",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    },
                    {
                        "id": 4297,
                        "product": {
                            "id": 205,
                            "name": "Брокколи",
                            "weight": "20",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "302319",
                            "vendor2": "302319",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "",
                            "full_image": "",
                            "images": [
                                "",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    },
                    {
                        "id": 4298,
                        "product": {
                            "id": 206,
                            "name": "Кедровые орешки",
                            "weight": "15",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "302419",
                            "vendor2": "302419",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "",
                            "full_image": "",
                            "images": [
                                "",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    },
                    {
                        "id": 4299,
                        "product": {
                            "id": 207,
                            "name": "Кунжут",
                            "weight": "2",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "302519",
                            "vendor2": "302519",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "",
                            "full_image": "",
                            "images": [
                                "",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    },
                    {
                        "id": 4300,
                        "product": {
                            "id": 208,
                            "name": "Острый перец",
                            "weight": "2",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "302619",
                            "vendor2": "302619",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "",
                            "full_image": "",
                            "images": [
                                "",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    },
                    {
                        "id": 4460,
                        "product": {
                            "id": 586,
                            "name": "Арахис",
                            "weight": "20",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "302719",
                            "vendor2": "302719",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "",
                            "full_image": "",
                            "images": [
                                "",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    },
                    {
                        "id": 4461,
                        "product": {
                            "id": 587,
                            "name": "Кимчи",
                            "weight": "40",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "302819",
                            "vendor2": "302819",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "",
                            "full_image": "",
                            "images": [
                                "",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    },
                    {
                        "id": 4462,
                        "product": {
                            "id": 588,
                            "name": "Сливки",
                            "weight": "40",
                            "sity_info": {
                                "id": 0,
                                "price": 0,
                                "profit": 0
                            },
                            "vendor1": "302919",
                            "vendor2": "302919",
                            "type": {
                                "id": 7,
                                "name": "Для Вока"
                            },
                            "comment": "",
                            "alias": "",
                            "number": "",
                            "caloric": 0,
                            "pfs": "",
                            "composition": "",
                            "image": "",
                            "full_image": "",
                            "images": [
                                "",
                                ""
                            ],
                            "new": false,
                            "hot": false,
                            "hit": false,
                            "not_delivery_separately": false,
                            "volume": 0,
                            "article": ""
                        },
                        "parent": 450,
                        "quantity": 1
                    }
                ],
                "souse": [],
                "composite": []
            }
        }
    }
]
```

####  Как использовать     
 - Поменяйте настройки для подключения к БД в config/default.go
 - В /service/products  и  /service/setting реализуйте свои методы в зависимости от своей сруктуры в БД  
  
