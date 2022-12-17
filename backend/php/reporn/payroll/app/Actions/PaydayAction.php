<?php
namespace App\Actions;

use App\Models\Employee;

class PaydayAction
{
    public function execute(): void
    {
        foreach (Employee::all() as $employee) {
            $amount = $employee->payment_type->monthlyAmount();
            if ($amount == 0) {
                continue;
            }

            $employee->paychecks()->create([
                'net_amount' => $amount,
                'payed_at' => now(),
            ]);
        }
    }
}