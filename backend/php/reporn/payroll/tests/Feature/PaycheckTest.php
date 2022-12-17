<?php

use App\Enums\PaymentTypes;
use App\Models\Employee;
use App\Models\TimeLog;
use Carbon\Carbon;

use function Pest\Laravel\assertDatabaseCount;
use function Pest\Laravel\assertDatabaseHas;
use function Pest\Laravel\postJson;
use function Pest\Laravel\travelTo;

// 为年薪员工创建薪资支票
it('should create paychecks for salary employees', function () {
    $employees = Employee::factory()
        ->count(2)
        ->sequence(
            [
                'salary' => 50000 * 100,
                'payment_type' => PaymentTypes::SALARY->value
            ],
            [
                'salary' => 70000 * 100,
                'payment_type' => PaymentTypes::SALARY->value
            ],
        )->create();

    postJson(route('payday.store'))
        ->assertNoContent();

    assertDatabaseHas('paychecks', [
        'employee_id' => $employees[0]->id,
        'net_amount' => 416666,
    ]);
    assertDatabaseHas('paychecks', [
        'employee_id' => $employees[1]->id,
        'net_amount' => 583333,
    ]);
});

// 为时薪员工创建薪资支票
it('should create paychecks for hourly rate employees', function () {
    // 从指定时间开始（时光旅行）
    travelTo(Carbon::parse('2022-02-10'), function () {
        // 创建雇员
        $employee = Employee::factory([
            'hourly_rate' => 10 * 100,
            'payment_type' => PaymentTypes::HOURLY_RATE->value,
        ])->create();  // 创建于2022-02-10 

        $dayBeforeYesterday = now()->subDays(2); // 2022-02-08 00:00:00
        $yesterday = now()->subDay(); // 2022-02-09 00:00:00
        $today = now();  // 2022-02-10 00:00:00

        // 创建工作时长记录
        TimeLog::factory()
            ->count(3)
            ->sequence(
                [
                    'employee_id' => $employee,
                    'minutes' => 90,
                    'started_at' => $dayBeforeYesterday,  // 2022-02-08 00:00:00 
                    'stopped_at' => $dayBeforeYesterday->copy()->addMinutes(90) // 2022-02-08 01:30:00
                ],
                [
                    'employee_id' => $employee,
                    'minutes' => 15,
                    'started_at' => $yesterday,  // 2022-02-09 00:00:00 
                    'stopped_at' => $yesterday->copy()->addMinutes(15) // 2022-02-08 00:15:00
                ],
                [
                    'employee_id' => $employee,
                    'minutes' => 51,
                    'started_at' => $today,  // 2022-02-10 00:00:00
                    'stopped_at' => $today->copy()->addMinutes(51)  // 2022-02-10 00:51:00
                ],
            )
            ->create();

        // 调用 API 创建薪资支票
        // 本月总工作时长=90+15+51=156/60=3(四舍五入)，应付薪资=3*1,000=3,000cents=$30.00
        postJson(route('payday.store'))
            ->assertNoContent();

        // 断言数据库是否存在对应的薪资支票记录
        assertDatabaseHas('paychecks', [
            'employee_id' => $employee->id,
            'net_amount' => 30 * 100,
        ]);
    });
});

// 只为时薪雇员创建当月的薪资支票
it('should create paychecks for hourly rate employees only for current month', function () {
    travelTo(Carbon::parse('2022-02-10'), function () {
        // 当前时间 2022-02-10
        $employee = Employee::factory([
            'hourly_rate' => 10 * 100,
            'payment_type' => PaymentTypes::HOURLY_RATE->value,
        ])->create();

        Timelog::factory()
            ->count(2)
            ->sequence(
                [
                    'employee_id' => $employee,
                    'minutes' => 60,
                    'started_at' => now()->subMonth(), // 2022-01-10 00:00:00
                    'stopped_at' => now()->subMonth()->addMinutes(60) // 2022-01-10 01:00:00
                ],
                [
                    'employee_id' => $employee,
                    'minutes' => 60,
                    'started_at' => now(),  // 2022-02-10 00:00:00
                    'stopped_at' => now()->addMinutes(60)  // 2022-02-10 01:00:00
                ],
            )
            ->create();

        // 调用 API 创建薪资账单，只创建当月的（2月）
        postJson(route('payday.store'))
            ->assertNoContent();

        // 断言 paychecks 表中的薪资支票，总薪资应该是2月份的，不包含1月的
        assertDatabaseHas('paychecks', [
            'employee_id' => $employee->id,
            'net_amount' => 10 * 100,
        ]);
    });
});

// 不应该为没有工作时长记录的雇员创建薪资支票
it('should not create paychecks for hourly rate employees without time logs', function () {
    travelTo(Carbon::parse('2022-02-10'), function () {
        Employee::factory([
            'hourly_rate' => 10 * 100,
            'payment_type' => PaymentTypes::HOURLY_RATE->value,
        ])->create();
        
        postJson(route('payday.store'))
            ->assertNoContent();

        assertDatabaseCount('paychecks', 0);
    });
});
