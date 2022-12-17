<?php
namespace App\Actions;

use App\DTOs\EmployeeData;
use App\Models\Employee;
use Illuminate\Support\Facades\Validator;
use Illuminate\Validation\Rule;
use Illuminate\Validation\ValidationException;

class UpsertEmployeeAction
{    
    /**
     * @throws ValidationException
     */
    public function execute(Employee $employee, EmployeeData $employeeData): Employee
    {
        $this->validate($employeeData);

        $employee->full_name = $employeeData->fullName;
        $employee->email = $employeeData->email;
        $employee->department_id = $employeeData->department->id;
        $employee->job_title = $employeeData->jobTitle;
        $employee->payment_type = $employeeData->paymentType;
        $employee->salary = $employeeData->salary;
        $employee->hourly_rate = $employeeData->hourlyRate;
        $employee->save();
        return $employee;
    }

    /**
     * @throws ValidationException
     */
    private function validate(EmployeeData $employeeData): void
    {
        $rules = [
            $employeeData->paymentType => [
                'required',
                'numeric',
                Rule::notIn([0]),
            ]
        ];
        $validator = Validator::make([
            'paymentType' => $employeeData->paymentType,
            'salary' => $employeeData->salary,
            'hourlyRate' => $employeeData->hourlyRate,
        ], $rules);
        $validator->validate();
    }
}
