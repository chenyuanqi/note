
### 整洁架构（Clean Architecture）
1. 整洁架构简介
整洁架构（Clean Architecture）是一种软件设计原则，旨在降低模块之间的耦合度，提高代码的可读性、可维护性和可扩展性。整洁架构的核心思想是将软件分层，不同的层次具有不同的职责，层与层之间的依赖关系从内向外，即内层不依赖外层，外层依赖内层。
![ibGTGz.png](https://i.328888.xyz/2023/04/04/ibGTGz.png)

2. 整洁架构的四大层次
整洁架构主要分为四个层次，分别是：

  1) 实体层（Entities）：实体层包含了业务逻辑的核心概念和规则，是业务领域的核心模型。实体层不依赖于任何其他层次。

  2) 用例层（Use Cases）：用例层包含了应用程序的所有用例，定义了软件如何实现业务规则。用例层依赖于实体层，但不依赖于外部的细节（如数据库、UI等）。

  3) 接口适配器层（Interface Adapters）：接口适配器层负责将用例层与外部细节进行适配。它可以包括数据访问对象（DAO）、API控制器等。接口适配器层依赖于用例层和实体层。

  4) 框架和驱动层（Frameworks and Drivers）：框架和驱动层是软件与外部世界的接口，包括数据库、操作系统、第三方库等。这一层依赖于内部的所有层次。

3. 相关示例
假设我们要开发一个电子商务系统，我们可以按照以下方式应用整洁架构：

  1) 实体层：创建具有业务逻辑的实体类，例如产品（Product）、订单（Order）等。

  2) 用例层：编写用例类，例如创建订单（CreateOrder）、计算订单总价（CalculateOrderTotal）等。这些用例类依赖于实体类，但不依赖于具体的数据库或UI实现。

  3) 接口适配器层：为不同的外部细节实现适配器，例如数据库访问对象（DAO）和API控制器。例如，我们可以创建一个用于MySQL的订单DAO（OrderDaoMySQL）和一个用于MongoDB的订单DAO（OrderDaoMongoDB），还可以创建一个订单API控制器（OrderApiController）用于处理HTTP请求。

  4) 框架和驱动层：这一层包括我们使用的数据库、操作系统、Web框架等。在此示例中，可能包括MySQL、MongoDB、Spring Boot等。我们需要根据具体的技术选择，实现适当的适配器，并将其与用例层和实体层进行连接。

通过遵循整洁架构的原则，我们可以确保代码的可读性、可维护性和可扩展性。不同的层次之间的依赖关系清晰明了，这有助于我们更轻松地进行代码重构、技术升级和功能扩展。

4. 优缺点
整洁架构的优点包括：
  - 模块化：通过将系统分层，我们可以更好地组织代码，降低模块间的耦合度。
  - 可维护性：由于各层之间的依赖关系清晰，代码结构易于理解，更容易维护。
  - 可扩展性：系统的各个部分可以独立地进行扩展，而不会影响到其他部分。
  - 易于测试：因为依赖关系清晰，可以方便地针对各个层进行单元测试和集成测试。

整洁架构的缺点包括：
  - 学习成本：对于初学者来说，理解和应用整洁架构可能需要一定的时间和精力。
  - 代码量可能较多：由于分层和解耦，可能导致某些代码量变多，但这可以通过优化和重构来改善。

### 具体实现 - 实体层
以下是使用 Go 语言实现的产品（Product）和订单（Order）实体类示例。

1. 产品（Product）实体类：

```go
package entities

type Product struct {
    ID          string
    Name        string
    Description string
    Price       float64
}

func NewProduct(id, name, description string, price float64) *Product {
    return &Product{
        ID:          id,
        Name:        name,
        Description: description,
        Price:       price,
    }
}
```

2. 订单（Order）实体类：

```go
package entities

import "time"

type Order struct {
    ID         string
    CustomerID string
    Products   []*Product
    TotalPrice float64
    CreatedAt  time.Time
}

func NewOrder(id, customerID string, products []*Product, totalPrice float64) *Order {
    return &Order{
        ID:         id,
        CustomerID: customerID,
        Products:   products,
        TotalPrice: totalPrice,
        CreatedAt:  time.Now(),
    }
}

// 计算订单总价的方法
func (o *Order) CalculateTotalPrice() {
    totalPrice := 0.0
    for _, product := range o.Products {
        totalPrice += product.Price
    }
    o.TotalPrice = totalPrice
}
```

在这个示例中，我们使用Go语言定义了`Product`和`Order`实体类，并为它们提供了构造函数。`Order`类还包含一个`CalculateTotalPrice`方法，用于计算订单的总价。实体层主要关注业务逻辑和实体间的关系，不涉及具体的数据存储和UI实现。

### 具体实现 - 用例层
用例层主要包括处理业务逻辑的用例，例如创建订单（CreateOrder）和计算订单总价（CalculateOrderTotal）。以下是使用Go语言实现用例层的一个简单示例。

首先，我们需要定义一个订单存储库接口（OrderRepository），用于实现数据存储层的抽象。这样做的目的是为了使用例层与具体的数据存储实现解耦。

```go
package usecases

import (
    "entities"
)

type OrderRepository interface {
    Store(order *entities.Order) error
    FindByID(id string) (*entities.Order, error)
}
```

接下来，我们实现用例层的功能，例如创建订单（CreateOrder）：

```go
package usecases

import (
    "entities"
)

type OrderInteractor struct {
    OrderRepository OrderRepository
}

func (interactor *OrderInteractor) CreateOrder(customerID string, products []*entities.Product) (*entities.Order, error) {
    // 创建订单实例
    order := entities.NewOrder("", customerID, products, 0)

    // 计算订单总价
    order.CalculateTotalPrice()

    // 将订单存储到存储库中
    err := interactor.OrderRepository.Store(order)
    if err != nil {
        return nil, err
    }

    return order, nil
}
```

在这个示例中，我们定义了一个名为`OrderInteractor`的结构体，它包含一个`OrderRepository`接口。`OrderInteractor`结构体实现了`CreateOrder`方法，用于创建订单。此方法首先创建一个新的`Order`实例，然后计算订单总价，并将订单存储到存储库中。通过这种方式，我们使用例层与具体的数据存储实现解耦，使代码更具可维护性和可扩展性。

接下来，我们可以实现更多的用例，例如根据订单ID查找订单（FindOrderById）：

```go
package usecases

import (
    "entities"
)

func (interactor *OrderInteractor) FindOrderById(id string) (*entities.Order, error) {
    order, err := interactor.OrderRepository.FindByID(id)
    if err != nil {
        return nil, err
    }
    return order, nil
}
```

在这个示例中，我们在`OrderInteractor`结构体中添加了`FindOrderById`方法。这个方法会调用`OrderRepository`接口的`FindByID`方法来查询订单。这样一来，我们的用例层与具体的数据存储实现保持解耦，有助于提高代码的可维护性和可扩展性。

在电子商务系统示例中，我们可以实现一个计算订单总价的用例。这个用例的目的是根据订单中的所有产品的价格计算出订单的总价。下面是一个简单的实现示例：

首先，在`OrderInteractor`结构体中添加一个名为`CalculateOrderTotal`的方法：

```go
package usecases

import (
    "entities"
)

func (interactor *OrderInteractor) CalculateOrderTotal(order *entities.Order) float64 {
    order.CalculateTotalPrice()
    return order.TotalPrice
}
```

在这个示例中，`CalculateOrderTotal`方法接收一个`Order`实例作为参数。然后，它调用`Order`实体类中的`CalculateTotalPrice`方法来计算订单的总价。这个用例关注于业务逻辑，即计算订单总价，并与实体层（`Order`实体类）进行交互。

在实际使用中，您可能需要先根据订单ID查找订单，然后再计算订单总价。例如，您可以将`FindOrderById`和`CalculateOrderTotal`方法组合使用，如下所示：

```go
func main() {
    // 假设 orderID 是已知的订单ID
    orderID := "example_order_id"

    // 初始化 OrderInteractor 和 OrderRepository
    interactor := &usecases.OrderInteractor{
        OrderRepository: &yourConcreteOrderRepositoryImplementation{}, // 使用您的具体实现替换
    }

    // 根据订单ID查找订单
    order, err := interactor.FindOrderById(orderID)
    if err != nil {
        log.Fatalf("Error finding order: %v", err)
    }

    // 计算订单总价
    totalPrice := interactor.CalculateOrderTotal(order)
    fmt.Printf("The total price of order %s is: %f", order.ID, totalPrice)
}
```

在这个示例中，我们首先根据订单ID查找订单，然后使用`CalculateOrderTotal`方法计算订单总价。这样，我们可以实现用例层与实体层的解耦，使代码更具可维护性和可扩展性。

请注意，用例层应该只关注业务逻辑，并通过存储库接口与数据存储层进行交互。

### 具体实现 - 接口适配器层
接口适配器层主要负责适配不同的外部细节实现，例如数据库访问对象（DAO）和 API 控制器。以下是使用 Go 语言实现接口适配器层的一个简单示例。

首先，我们需要实现一个基于内存的订单存储库（MemoryOrderRepository），以满足`OrderRepository`接口的要求：

```go
package adapters

import (
	"errors"
	"entities"
	"sync"
	"usecases"
)

type MemoryOrderRepository struct {
	orders map[string]*entities.Order
	mu     sync.Mutex
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	return &MemoryOrderRepository{
		orders: make(map[string]*entities.Order),
	}
}

func (r *MemoryOrderRepository) Store(order *entities.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.ID] = order
	return nil
}

func (r *MemoryOrderRepository) FindByID(id string) (*entities.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	order, found := r.orders[id]
	if !found {
		return nil, errors.New("Order not found")
	}
	return order, nil
}
```

接下来，我们实现一个简单的HTTP API控制器，用于处理创建订单和查询订单的请求：

```go
package adapters

import (
	"encoding/json"
	"net/http"
	"usecases"
)

type OrderController struct {
	OrderInteractor *usecases.OrderInteractor
}

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// 解析请求体以获取订单信息（例如客户ID和产品列表）
	var input struct {
		CustomerID string
		ProductIDs []string
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// 从请求中获取产品列表
	products := []*entities.Product{} // 在实际项目中，需要通过产品ID从数据库或其他数据源中获取产品信息

	// 创建订单
	order, err := c.OrderInteractor.CreateOrder(input.CustomerID, products)
	if err != nil {
		http.Error(w, "Error creating order", http.StatusInternalServerError)
		return
	}

	// 将订单返回给客户端
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (c *OrderController) GetOrder(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取订单ID
	orderID := r.URL.Query().Get("id")
	if orderID == "" {
		http.Error(w, "Missing order ID", http.StatusBadRequest)
		return
	}

	// 查找订单
	order, err := c.OrderInteractor.FindOrderById(orderID)
	if err != nil {
		http.Error(w, "Error finding order", http.StatusInternalServerError)
		return
	}

	// 将订单返回给客户端
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
```

在这个示例中，我们实现了一个基于内存的订单存储库（MemoryOrderRepository）和一个简单的HTTP API控制器（OrderController）。这些组件可以与用例层进行交互。

现在，我们已经实现了基于内存的订单存储库（MemoryOrderRepository）和HTTP API控制器（OrderController），接下来我们需要将它们集成到Web服务器中。下面是一个简单的Web服务器实现示例，它使用了`net/http`包来处理HTTP请求。

```go
package main

import (
	"adapters"
	"net/http"
	"usecases"
)

func main() {
	// 创建内存订单存储库实例
	orderRepository := adapters.NewMemoryOrderRepository()

	// 创建订单用例层实例
	orderInteractor := &usecases.OrderInteractor{
		OrderRepository: orderRepository,
	}

	// 创建订单控制器实例
	orderController := &adapters.OrderController{
		OrderInteractor: orderInteractor,
	}

	// 注册HTTP路由和处理函数
	http.HandleFunc("/create_order", orderController.CreateOrder)
	http.HandleFunc("/get_order", orderController.GetOrder)

	// 启动Web服务器
	http.ListenAndServe(":8080", nil)
}
```

在这个示例中，我们首先创建了内存订单存储库（MemoryOrderRepository）实例、订单用例层实例（OrderInteractor）和订单控制器实例（OrderController）。然后，我们使用`http.HandleFunc`函数注册了HTTP路由和处理函数，最后使用`http.ListenAndServe`启动了Web服务器。

现在，您可以运行此示例程序，并通过发送HTTP请求来创建和查询订单。例如，您可以使用`curl`工具或其他HTTP客户端（如Postman）向`/create_order`和`/get_order`端点发送请求。

整洁架构通过将关注点分离到不同的层次，使得代码更容易理解和维护。在这个示例中，我们展示了如何使用Go语言实现整洁架构中的各个层次。实际上，您可能需要根据您的项目需求和使用的技术栈来调整代码实现。关键在于保持各个层次之间的解耦，并关注每个层次的主要职责。


### 具体实现 - 框架和驱动层
框架和驱动层主要负责支持各种外部系统和服务，例如数据库、消息队列、Web框架等。以下是使用Go语言实现框架和驱动层的一个简单示例。

假设我们需要实现一个基于MySQL的订单存储库（MysqlOrderRepository），以满足`OrderRepository`接口的要求。我们可以使用`database/sql`包和`github.com/go-sql-driver/mysql`驱动来实现这个存储库。

首先，安装MySQL驱动：

```sh
go get -u github.com/go-sql-driver/mysql
```

然后，实现`MysqlOrderRepository`：

```go
package drivers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"entities"
	"usecases"
)

type MysqlOrderRepository struct {
	db *sql.DB
}

func NewMysqlOrderRepository(dataSourceName string) (*MysqlOrderRepository, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &MysqlOrderRepository{db: db}, nil
}

func (r *MysqlOrderRepository) Store(order *entities.Order) error {
	// 在此处实现存储订单到MySQL数据库的逻辑
	// 例如，执行INSERT INTO语句
	return nil
}

func (r *MysqlOrderRepository) FindByID(id string) (*entities.Order, error) {
	// 在此处实现从MySQL数据库中查找订单的逻辑
	// 例如，执行SELECT语句
	return nil, nil
}
```

此外，您可能需要一个Web框架来简化HTTP路由和处理程序的实现。在这个示例中，我们将使用`github.com/gin-gonic/gin`框架。

首先，安装Gin框架：

```sh
go get -u github.com/gin-gonic/gin
```

然后，使用Gin框架重写Web服务器的实现：

```go
package main

import (
	"adapters"
	"drivers"
	"github.com/gin-gonic/gin"
	"usecases"
)

func main() {
	// 创建MySQL订单存储库实例
	orderRepository, err := drivers.NewMysqlOrderRepository("user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		panic(err)
	}

	// 创建订单用例层实例
	orderInteractor := &usecases.OrderInteractor{
		OrderRepository: orderRepository,
	}

	// 创建订单控制器实例
	orderController := &adapters.OrderController{
		OrderInteractor: orderInteractor,
	}

	// 创建Gin引擎实例
	router := gin.Default()

	// 注册HTTP路由和处理函数
	router.POST("/create_order", orderController.CreateOrder)
	router.GET("/get_order", orderController.GetOrder)

	// 启动Web服务器
	router.Run(":8080")
}
```

在这个示例中，我们使用了`github.com/go-sql-driver/mysql`驱动实现了一个基于MySQL的订单存储库（MysqlOrderRepository），并使用`github.com/gin-gonic/gin`框架简化了Web服务器的实现。这些组件属于整洁架构中的框架和驱动层，它们与其他层次进行解耦，可以根据项目需求和使用的技术栈进行替换。

如果您需要使用消息队列，例如RabbitMQ，作为系统间通信的方式，您可以创建一个发送消息的驱动层组件。这里以`github.com/streadway/amqp`库为例：

首先，安装RabbitMQ驱动：

```sh
go get -u github.com/streadway/amqp
```

然后，实现一个简单的RabbitMQ消息发送器：

```go
package drivers

import (
	"github.com/streadway/amqp"
)

type RabbitMQSender struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQSender(url string) (*RabbitMQSender, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQSender{conn: conn, channel: channel}, nil
}

func (s *RabbitMQSender) Send(queueName string, message []byte) error {
	_, err := s.channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return s.channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
}
```

在这个示例中，我们实现了一个简单的RabbitMQ消息发送器（RabbitMQSender），可以发送消息到指定队列。在实际项目中，您可以将这个发送器与其他层次组件结合使用，例如在订单创建成功后，通过RabbitMQ通知其他系统。

整洁架构使得您可以根据需要自由替换框架和驱动层的组件，而不会影响其他层次。例如，您可以在不修改用例层或接口适配器层的情况下，将MySQL订单存储库替换为PostgreSQL存储库。这有助于提高代码的可维护性和可扩展性。
