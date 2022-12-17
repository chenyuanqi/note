<?php

namespace App\Http\Controllers;

use App\Http\Resources\PaycheckResource;
use App\Models\Employee;
use App\Models\Paycheck;
use Illuminate\Http\Request;
use Spatie\QueryBuilder\QueryBuilder;

class EmployeePaycheckController extends Controller
{
    public function index(Request $request, Employee $employee)
    {
        $paychecks = QueryBuilder::for(Paycheck::class)
            ->allowedIncludes(['employee'])
            ->whereBelongsTo($employee)
            ->get(); 

        return PaycheckResource::collection($paychecks);
    }
}
