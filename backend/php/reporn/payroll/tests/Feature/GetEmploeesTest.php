<?php

use App\Enums\PaymentTypes;
use App\Models\Department;
use App\Models\Employee;
use App\Models\User;
use Laravel\Sanctum\Sanctum;

use function Pest\Laravel\getJson;

// 测试部门员工列表
it('should return all employees for a department', function () {
    Sanctum::actingAs(User::factory()->create(), ['*']);

    $development = Department::factory(['name' => 'Development'])->create();
    $marketing = Department::factory(['name' => 'Marketing'])->create();

    $developers = Employee::factory([
        'department_id' => $development->id,
        'payment_type' => PaymentTypes::from('salary')->value,
    ])->count(5)->create();

    Employee::factory([
        'department_id' => $marketing->id,
        'payment_type' => PaymentTypes::from('hourlyRate')->value,

    ])->count(2)->create();

    $employees = getJson(route('department.employees.index', ['department' => $development]))
        ->json('data');

    expect($employees)->toHaveCount(5);
    expect($employees)
        ->each(fn ($employee) => $employee->id->toBeIn($developers->pluck('uuid')));
});

// 测试过滤器
it('should filter employees', function () {
    Sanctum::actingAs(User::factory()->create(), ['*']);

    $development = Department::factory(['name' => 'Development'])->create();
    $marketing = Department::factory(['name' => 'Marketing'])->create();

    Employee::factory([
        'department_id' => $development->id,
        'payment_type' => PaymentTypes::from('salary')->value,
    ])->count(4)->create();

    $developer = Employee::factory([
        'full_name' => 'Test John Doe',
        'department_id' => $development->id,
        'payment_type' => PaymentTypes::from('salary')->value,
    ])->create();

    Employee::factory([
        'department_id' => $marketing->id,
        'payment_type' => PaymentTypes::from('hourlyRate')->value,
    ])->count(2)->create();

    $employees = getJson(
        route('department.employees.index', [
            'department' => $development,
            'filter' => [
                'full_name' => 'Test',
            ]
        ])
    )->json('data');

    expect($employees)->toHaveCount(1);
    expect($employees[0])->id->toBe($developer->uuid);
});
