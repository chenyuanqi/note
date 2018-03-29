
### 什么是 BDD?
有两种不同的行为驱动开发 (BDD): SpecBDD 和 StoryBDD。 SpecBDD 专注于代码的技术行为，而 StoryBDD 专注于业务逻辑或功能的行为和互动。  

采用 StoryBDD 时, 你编写可读的故事来描述应用程序的行为。接着这些故事可以作为应用程序的实际测试案例执行。Behat 是使用在 PHP 应用程序中的 StoryBDD框架，它受到 Ruby 的 Cucumber 项目的启发并且实现了 Gherkin DSL 來描述功能的行为。  

采用 SpecBDD 时, 你编写规格来描述实际的代码应该有什么行为。你应该描述函数或者方法应该有什么行为，而不是测试函数或者方法。PHP 提供了 PHPSpec 框架来达到这个目的，这个框架受到了 Ruby 的 RSpec project 项目的启发。  

### BDD 在 PHP 的世界
[Behat](http://behat.org/), PHP 的 StoryBDD 框架， 受到了 Ruby’s Cucumber 项目的启发。  
[PHPSpec](http://www.phpspec.net/), PHP 的 SpecBDD 框架， 受到了 Ruby’s RSpec 项目的启发。  
[Codeception](http://codeception.com/) 是一个使用 BDD 准则的全栈测试框架。  
