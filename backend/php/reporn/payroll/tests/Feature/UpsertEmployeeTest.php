<?php

use App\Models\Department;
use App\Models\Employee;

use function Pest\Laravel\postJson;

it('should return 422 if email is invalid', function (?string $email) {
    Employee::factory([
        'email' => 'taken@example.com',
    ])->create();

    postJson(route('employees.store'), [
        'fullName' => 'Test Employee',
        'email' => $email,
        'departmentId' => Department::factory()->create()->uuid,
        'jobTitle' => 'BE Developer',
        'paymentType' => 'salary',
        'salary' => 75000 * 100,
    ])->assertInvalid(['email']);
})->with([
    'taken@example.com',
    'invalid',
    null,
    '',
]);

it('should return 422 if payment type is invalid', function () {
    postJson(route('employees.store'), [
        'fullName' => 'Test Employee',
        'email' => 'test@example.com',
        'departmentId' => Department::factory()->create()->uuid,
        'jobTitle' => 'BE Developer',
        'paymentType' => 'invalid',
        'salary' => 75000 * 100,
    ])->assertInvalid(['paymentType']);
});

it('should return 422 if salary missing when payment type is salary', function (string $paymentType, ?int $salary) {
    postJson(route('employees.store'), [
        'fullName' => 'Test Employee',
        'email' => 'test@example.com',
        'departmentId' => Department::factory()->create()->uuid,
        'jobTitle' => 'BE Developer',
        'paymentType' => $paymentType,
        'salary' => $salary,
    ])->assertInvalid(['salary']);
})->with([
    ['paymentType' => 'salary', 'salary' => null],
    ['paymentType' => 'salary', 'salary' => 0],
]);

it('should return 422 if hourly rate missing when payment type is hourly rate', function (string $paymentType, ?int $hourlyRate) {
    postJson(route('employees.store'), [
        'fullName' => 'Test Employee',
        'email' => 'test@example.com',
        'departmentId' => Department::factory()->create()->uuid,
        'jobTitle' => 'BE Developer',
        'paymentType' => $paymentType,
        'hourlyRate' => $hourlyRate,
    ])->assertInvalid(['hourlyRate']);
})->with([
    ['paymentType' => 'hourlyRate', 'hourlyRate' => null],
    ['paymentType' => 'hourlyRate', 'hourlyRate' => 0],
]);

it('should store an employee with payment type salary', function () {
    $employee = postJson(route('employees.store'), [
        'fullName' => 'John Doe',
        'email' => 'test@example.com',
        'departmentId' => Department::factory()->create()->uuid,
        'jobTitle' => 'BE Developer',
        'paymentType' => 'salary',
        'salary' => 75000 * 100,
    ])->json('data');

    expect($employee)
        ->attributes->name->toBe('John Doe')
        ->attributes->email->toBe('test@example.com')
        ->attributes->jobTitle->toBe('BE Developer')
        ->attributes->payment->type->toBe('salary')
        ->attributes->payment->amount->cents->toBe(75000 * 100)
        ->attributes->payment->amount->dollars->toBe('$75,000.00');
});

it('should store an employee with payment type hourly rate', function () {
    $employee = postJson(route('employees.store'), [
        'fullName' => 'John Doe',
        'email' => 'test@example.com',
        'departmentId' => Department::factory()->create()->uuid,
        'jobTitle' => 'BE Developer',
        'paymentType' => 'hourlyRate',
        'hourlyRate' => 30 * 100,
    ])->json('data');

    expect($employee)
        ->attributes->name->toBe('John Doe')
        ->attributes->email->toBe('test@example.com')
        ->attributes->jobTitle->toBe('BE Developer')
        ->attributes->payment->type->toBe('hourlyRate')
        ->attributes->payment->amount->cents->toBe(30 * 100)
        ->attributes->payment->amount->dollars->toBe('$30.00');
});
