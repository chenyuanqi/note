<?php
namespace App\Payment;

use App\Enums\PaymentTypes;
use App\Models\Employee;
use RuntimeException;

class Salary extends PaymentType
{
    public function __construct(Employee $employee)
    {
        throw_if(
            $employee->salary === null,
            new RuntimeException('salary cannot be null')
        );
        parent::__construct($employee);
    }

    public function monthlyAmount(): int
    {
        // 月薪 = 年薪 / 12
        return $this->employee->salary / 12;   
    }

    public function type(): string
    {
        return PaymentTypes::SALARY->value;
    }

    public function amount(): int
    {
        return $this->employee->salary;
    }
}