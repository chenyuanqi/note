<?php

use App\Http\Controllers\DepartmentController;
use App\Http\Controllers\DepartmentEmployeeController;
use App\Http\Controllers\EmployeeController;
use App\Http\Controllers\EmployeePaycheckController;
use App\Http\Controllers\PaycheckController;
use Illuminate\Support\Facades\Route;

Route::apiResource('departments', DepartmentController::class);
Route::apiResource('employees', EmployeeController::class);

Route::get(
    'departments/{department}/employees',
    [DepartmentEmployeeController::class, 'index']
)->name('department.employees.index');

Route::post(
    'paycheck',
    [PaycheckController::class, 'store']
)->name('payday.store');

Route::get(
    'employees/{employee}/paychecks',
    [EmployeePaycheckController::class, 'index']
)->name('employee.paychecks.index');
