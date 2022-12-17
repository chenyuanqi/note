<?php
namespace App\DTOs;

use App\Http\Requests\UpsertEmployeeRequest;
use App\Models\Department;

class EmployeeData
{
    public function __construct(
        public readonly string $fullName,
        public readonly string $email,
        public readonly Department $department,
        public readonly string $jobTitle,
        public readonly string $paymentType,
        public readonly ?int $salary,
        public readonly ?int $hourlyRate,
    ) {}

    public static function fromRequest(UpsertEmployeeRequest $request): self
    {
        return new static(
            $request->fullName,
            $request->email,
            $request->getDepartment(),
            $request->jobTitle,
            $request->paymentType,
            $request->salary,
            $request->hourlyRate,
        );
    }
}
