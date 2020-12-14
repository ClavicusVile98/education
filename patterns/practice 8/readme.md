## 1. Найти в проекте Front Controller и расписать классы, которые с ним взаимодействуют.

С **Front Controller** взаимодействуют:
класс *Request*, функция *createFromGlobals()*;
класс *ContainerBuilder*, создание экземпляра класса;
класс *Framework\Registry*, функция *addContainer()*;
класс *Kernel*, функция *handle()*;
класс *Response*, функция *send()*.

## 2. Найти в проекте Registry и объяснить, почему он был применён.

В классе **src/Order/Basket.php** используется глобальная переменная **private const BASKET_DATA_KEY = 'basket'**,
которая хранит сессионный ключ всех продуктов в корзине.
В процедуре **addProduct()** наглядно представлена работа реестра.
В строчке **$basket = $this->session->get(static::BASKET_DATA_KEY, [])** мы получаем продукты из корзины. 

Далее после проверки на наличие определенного продукта в корзине, мы сохраняем данные **$this->session->set(static::BASKET_DATA_KEY, $basket**.

Также реестр применен в классе **Security** в методах **getuser(), authentication() и logout()**.
Обозначена переменная как **private const SESSION_USER_IDENTITY = 'userId'**.

## 3. Рассказать, какой тип модели используется в проекте.

В проекте используется тип модели Table Module.
