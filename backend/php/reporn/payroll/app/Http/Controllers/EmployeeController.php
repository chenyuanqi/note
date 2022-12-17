<?php

namespace App\Http\Controllers;

use App\Actions\UpsertEmployeeAction;
use App\DTOs\EmployeeData;
use App\Http\Requests\UpsertEmployeeRequest;
use App\Http\Resources\EmployeeResource;
use App\Models\Employee;
use Spatie\QueryBuilder\QueryBuilder;

class EmployeeController extends Controller
{
    public function __construct(
        private readonly UpsertEmployeeAction $upsertEmployee
    ) {
    }


    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
       $employees = QueryBuilder::for(Employee::class)
        ->allowedFilters(
            ['full_name', 'job_title', 'email', 'department.name']
        )
        ->allowedIncludes(['department'])
        ->get();
        
        return EmployeeResource::collection($employees);
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  UpsertEmployeeRequest  $request
     * @return \Illuminate\Http\Response
     */
    public function store(UpsertEmployeeRequest $request)
    {
        $employee = $this->upsertEmployee->execute(
            new Employee(),
            EmployeeData::fromRequest($request)
        );

        return EmployeeResource::make($employee)->response();
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  UpsertEmployeeRequest $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(UpsertEmployeeRequest $request, Employee $employee)
    {
        $employee = $this->upsertEmployee->execute(
            $employee,
            EmployeeData::fromRequest($request)
        );
        return response()->noContent();
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        //
    }
}
