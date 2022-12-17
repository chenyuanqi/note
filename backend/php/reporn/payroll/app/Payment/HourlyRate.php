<?php

namespace App\Payment;

use App\Enums\PaymentTypes;
use App\Models\Employee;
use App\Models\TimeLog;
use RuntimeException;

class HourlyRate extends PaymentType
{
    public function __construct(Employee $employee)
    {
        throw_if(
            $employee->hourly_rate === null,
            new RuntimeException('Hourly rate cannot be null')
        );
        parent::__construct($employee);
    }

    public function monthlyAmount(): int
    {
        // 工作时长 = 每月工作总分钟数 / 60
        $hoursWorked = TimeLog::query()
            ->whereBetween('stopped_at', [
                now()->startOfMonth(),
                now()->endOfMonth()
            ])
            ->sum('minutes') / 60;
        
        // 月薪 = 四舍五入(时长) * 时薪
        return round($hoursWorked) * $this->employee->hourly_rate;
    }

    public function type(): string
    {
        return PaymentTypes::HOURLY_RATE->value;
    }

    public function amount(): int
    {
        return $this->employee->hourly_rate;
    }
}
