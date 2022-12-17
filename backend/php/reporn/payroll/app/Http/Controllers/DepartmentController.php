<?php

namespace App\Http\Controllers;

use App\Actions\CreateDepartmentAction;
use App\Actions\UpdateDepartmentAction;
use App\DTOs\DepartmentData;
use App\Http\Requests\StoreDepartmentRequest;
use App\Http\Requests\UpdateDepartmentRequest;
use App\Http\Resources\DepartmentResource;
use App\Models\Department;
use Illuminate\Http\Request;
use Illuminate\Http\Response;

class DepartmentController extends Controller
{
    public function __construct(
        private readonly CreateDepartmentAction $createDepartment,
        private readonly UpdateDepartmentAction $updateDepartment
    ) {
    }

    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        return DepartmentResource::collection(Department::all());
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
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(StoreDepartmentRequest $request)
    {
        $departmentData = new DepartmentData(...$request->validated());
        $department = $this->createDepartment->execute($departmentData);

        return DepartmentResource::make($department)
            ->response();
    }

    /**
     * Display the specified resource.
     *
     * @param  Department  $department
     * @return \Illuminate\Http\Response
     */
    public function show(Department $department)
    {
        return DepartmentResource::make($department);
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
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(UpdateDepartmentRequest $request, Department $department)
    {
        $departmentData = new DepartmentData(...$request->validated());
        $department = $this->updateDepartment->execute($department, $departmentData);

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
