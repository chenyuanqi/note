<?php
namespace App\Enums;

use App\Models\Employee;
use App\Payment\HourlyRate;
use App\Payment\PaymentType;
use App\Payment\Salary;

enum PaymentTypes: string
{
    case SALARY = 'salary';
    case HOURLY_RATE = 'hourlyRate';

    public function makePaymentType(Employee $employee): PaymentType
    {
        return match($this) {
            self::SALARY => new Salary($employee),
            self::HOURLY_RATE => new HourlyRate($employee),
        };
    }
}